# Snowflake ID Generator API

This project implements a Snowflake ID generator as a RESTful API using the Gin web framework in Go. Snowflake IDs are unique identifiers that are distributed and generated without the need for synchronization. This makes them ideal for use in distributed systems where you need to ensure uniqueness across different machines without central coordination.

## Features

- **Unique ID Generation**: Generates unique IDs based on the Snowflake algorithm.
- **High Performance**: Utilizes Go's concurrency features and the efficiency of Gin.
- **RESTful API**: Simple and easy-to-use HTTP GET endpoint for obtaining IDs.

## Prerequisites

Before you begin, ensure you have installed Go (version 1.16 or later) on your machine. You can download and install it from [golang.org](https://github.com/Benjamin-Pasternak/snowflake_go.git).

## Getting Started

### Installation

Clone the repository to your local machine:

```bash
git clone https://your-repository-url.git
cd snowflake-id-generator
```
Install the required dependencies:

```bash
go mod tidy
```
## Running the Server
To start the server, run:

```bash
go run ./cmd/snowflake_go/main.go
```
This will start the Gin server on `localhost:8080`, and it will start listening for requests.

## Usage
To generate a new Snowflake ID, make a GET request to the `/generate-id` endpoint:

```bash
curl http://localhost:8080/generate-id
```

This will return a JSON response containing the generated ID:

```bash
{
    "id": 1234567890123456789
}
```

## API Reference
GET /generate-id

Generates a new unique Snowflake ID.

URL: `/generate-id`
Method: `GET`
Auth required: No
Permissions required: None
Success Response:

Code: `200 OK`
Content example:
```bash
{
    "id": 1234567890123456789
}
```

## Contributing
Contributions to this project are welcome! Please fork the repository and submit a pull request with your enhancements or fixes. Follow the (Forking workflow))[https://www.atlassian.com/git/tutorials/comparing-workflows/forking-workflow].