# 📝 Task Tracker API

A simple and clean **REST API** for managing tasks, built with **Go**, **Gin**, and **GORM**.  
This API allows you to create, read, update, and delete tasks, demonstrating a production-style Go project structure.

---

## 🚀 Features

- Create, list, update, and delete tasks
- Modular architecture: handlers, services, repositories, models
- Uses SQLite for local persistence
- Easy to extend with PostgreSQL, authentication, or deployment
- Lightweight and fast with Gin
- ✅ Includes a demo script to showcase the API automatically

---

## 📂 Project Structure
```bash
task-tracker/
├── cmd/
│ └── server/ # Main application entry
├── internal/
│ ├── handler/ # HTTP handlers
│ ├── service/ # Business logic
│ ├── repository/ # Database layer
│ └── model/ # Structs and types
├── go.mod
├── demo.go # Demo script for automated API demonstration
└── README.md
```
---

## ⚙️ Installation

1. Clone the repo:

```bash
git clone https://github.com/g-laliotis/task-tracker.git
cd task-tracker

Install dependencies:
go mod tidy

Run the server:
go run ./cmd/server
Server runs at http://localhost:8080.

🌐 API Endpoints
Method	URL	Body	Description
GET	/tasks	—	List all tasks
POST	/tasks	{ "title": "Task Name" }	Create a new task
PUT	/tasks/:id	{ "completed": true/false }	Update a task
DELETE	/tasks/:id	—	Delete a task

💻 Example Usage
Create a Task:
curl -X POST http://localhost:8080/tasks \
-H "Content-Type: application/json" \
-d '{"title":"Learn Go"}'
Get All Tasks:
curl http://localhost:8080/tasks
Update Task Completion:
curl -X PUT http://localhost:8080/tasks/1 \
-H "Content-Type: application/json" \
-d '{"completed": true}'
Delete a Task:
curl -X DELETE http://localhost:8080/tasks/1

🧪 Demo Script

Run the demo:
Make sure the server is running:
go run ./cmd/server
In a separate terminal, run:
go run demo.go
The demo will:
Create three tasks
List all tasks
Mark task 1 as completed
Delete task 2
Show the final list of tasks

🧪 Future Improvements
Add user authentication
Switch to PostgreSQL for production
Add Swagger/OpenAPI docs
Include unit and integration tests
Deploy to Render, Fly.io, or Railway

🧑‍💻 Author
Giorgos Laliotis