# AdvancedToDo

[![Go Badge](https://img.shields.io/badge/Go-1.17-blue?logo=go&style=flat)](https://golang.org/)
[![MongoDB Badge](https://img.shields.io/badge/MongoDB-4.4-green?logo=mongodb&style=flat)](https://www.mongodb.com/)
[![JWT Badge](https://img.shields.io/badge/JWT-JSON%20Web%20Tokens-orange?logo=jwt.io&style=flat)](https://jwt.io/)
[![Twilio Badge](https://img.shields.io/badge/Twilio-voice%20calling-red?logo=twilio&style=flat)](https://www.twilio.com/)

AdvancedToDo is a task management application built using Go, MongoDB, JWT authentication, and Twilio for voice calling reminders.

## Directory Structure

```
AdvancedToDo
|-- api
|   |-- handlers
|       |-- handlers.go
|-- cron
|   |-- cron.go
|-- db
|   |-- mongo.go
|-- middleware
|   |-- middleware.go
|-- models
|   |-- models.go
|-- go.mod
|-- go.sum
|-- main.go
```

## Technologies Used

- **Go**: Go is a statically typed, compiled programming language designed for simplicity and efficiency.
- **MongoDB**: MongoDB is a NoSQL document-oriented database that provides high performance, high availability, and easy scalability.
- **JWT (JSON Web Tokens)**: JWT is an open standard (RFC 7519) that defines a compact and self-contained way for securely transmitting information between parties as a JSON object.
- **Twilio**: Twilio is a cloud communications platform as a service company that allows software developers to programmatically make and receive phone calls, send and receive text messages, and perform other communication functions using its web service APIs.

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/AdvancedToDo.git
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
   go run main.go
   ```

## Usage

1. Create a task:

   ```bash
   POST /api/v1/tasks
   ```

2. Create a subtask:

   ```bash
   POST /api/v1/subtasks
   ```

3. Get all user tasks:

   ```bash
   GET /api/v1/tasks
   ```

4. Get all user subtasks:

   ```bash
   GET /api/v1/subtasks
   ```

5. Update a task:

   ```bash
   PUT /api/v1/tasks/:id
   ```

6. Update a subtask:

   ```bash
   PUT /api/v1/subtasks/:id
   ```

7. Delete a task:

   ```bash
   DELETE /api/v1/tasks/:id
   ```

8. Delete a subtask:

   ```bash
   DELETE /api/v1/subtasks/:id
   ```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## Authors

- Samriddha Basu

## Acknowledgements

Special thanks to the creators of Go, MongoDB, JWT, and Twilio for their amazing technologies!

## Support

For support, contact [samriddhabasu1234@gmail.com](mailto:samriddhabasu1234@gmail.com).
