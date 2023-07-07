package main

import (
	"rnav2022/goly/model"
	"rnav2022/goly/server"
)

// func init() {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("Loaded .env")
// }

func main() {
	model.Setup()
	server.SetupAndListen()
}
