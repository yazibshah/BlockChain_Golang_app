# Blockchain Application

This is a blockchain application built using Go (Golang) and the Gin framework.

## Description

The blockchain application allows users to interact with a decentralized content sharing platform. Users can create posts, view existing posts, and participate in the blockchain network.

## Features

- Create new posts with content and author information
- View and display existing posts in reverse chronological order
- Request to mine new blocks in the blockchain network

## Requirements

- Go (Golang) installed on your system
- External packages and dependencies:
  - `github.com/gin-gonic/gin`: Web framework for building the application's routes and handling HTTP requests
  - `github.com/go-resty/resty/v2`: HTTP client library for making requests to the blockchain network

## Installation

1. Clone the repository:


2. Change into the project directory:


3. Install the necessary dependencies using Go modules:

go mod download


## Usage

1. Build and run the application:



go run main.go:
The application will start a web server running on `http://localhost:8080`.

2. Open a web browser and visit `http://localhost:8080` to access the blockchain application.

3. Interact with the application by creating new posts, viewing existing posts, and performing other available actions.

## Contributing

Contributions to the project are welcome! If you find any bugs, have suggestions, or would like to contribute new features or improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

