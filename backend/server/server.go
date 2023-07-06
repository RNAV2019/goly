package server

import (
	"fmt"
	"log"
	"rnav2022/goly/model"
	"rnav2022/goly/utils"
	"strconv"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var clients = make(map[*websocket.Conn]bool)

// WebsocketHandler handles websocket connections
func websocketHandler(c *websocket.Conn) {
	clients[c] = true

	// Handle WebSocket messages from the client if needed
	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("WebSocket read failed:", err)
			delete(clients, c)
			break
		}
		log.Println("Received message:", string(msg))
	}

	// Cleanup when client disconnects
	defer c.Close()
	delete(clients, c)
}

// Redirects the user to the goly url
func Redirect(c *fiber.Ctx) error {
	url := c.Params("redirect")
	goly, err := model.FindByGolyUrl(url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not find goly " + err.Error(),
		})
	}
	// Logging how many times it was clicked
	goly.Clicked += 1
	err = model.UpdateGoly(goly)
	if err != nil {
		fmt.Println("error updating goly " + err.Error())
	}

	go func() {
		// Notify connected WebSocket clients
		message := []byte("redirect")
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("WebSocket write failed:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}()

	return c.Redirect(goly.Redirect, fiber.StatusTemporaryRedirect)
}

// GetAllGolies gets all golies
func GetAllGolies(c *fiber.Ctx) error {
	golies, err := model.GetAllGolies()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error getting all goly links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(golies)
}

// GetGoly gets a goly
func GetGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}
	goly, err := model.GetGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not retrieve goly " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)
}

// CreateGoly creates a goly
func CreateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}
	goly.URL = utils.RandomURL(8)
	err = model.CreateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not create goly " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)
}

// UpdateGoly updates a goly
func UpdateGoly(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var goly model.Goly
	err := c.BodyParser(&goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error parsing JSON " + err.Error(),
		})
	}
	err = model.UpdateGoly(goly)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not update goly " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(goly)
}

// DeleteGoly deletes a goly
func DeleteGoly(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not parse id " + err.Error(),
		})
	}
	err = model.DeleteGoly(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error could not delete goly " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "goly deleted",
	})
}

func SetupAndListen() {
	router := fiber.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// WebSocket endpoint
	router.Get("/ws", websocket.New(websocketHandler))

	// Redirect endpoint
	router.Get("/r/:redirect", Redirect)

	// Goly endpoints
	router.Get("/goly", GetAllGolies)
	router.Get("/goly/:id", GetGoly)
	router.Post("/goly", CreateGoly)
	router.Patch("/goly", UpdateGoly)
	router.Delete("/goly/:id", DeleteGoly)
	log.Fatal(router.Listen(":3000"))
}
