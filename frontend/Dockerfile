# Start with the base Node.js image
FROM node:20.2.0-bullseye-slim

# Set the working directory
WORKDIR /app

# Copy the package.json and yarn.lock files
COPY package.json ./

# Install the dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Check svelte is working
RUN npm run check

# Build the Svelte application
RUN npm run build

# Expose the port on which the application will run
EXPOSE 8080

# Set the entrypoint for the Docker container
CMD ["npm", "start"]