📝 Task Tracker API

A multi-user task/job management REST API built with Go, Gin, and GORM, ready to deploy on Render.
```bash
🚀 Features
🔐 User signup and login with JWT authentication
🧩 Create, list, update, and delete tasks (per user)
💾 SQLite for local testing, PostgreSQL for production
⚙️ Modular architecture: handlers, services, repositories, models
⚡ Lightweight and fast with Gin
🌍 Fully deployable to Render, Railway, or Fly.io
🧪 Includes a demo script that runs automatically against your live API
```

```bash
📂 Project Structure
task-tracker/
├── cmd/
│   └── server/        # Main application entry
├── internal/
│   ├── handler/       # HTTP handlers
│   ├── service/       # Business logic
│   ├── repository/    # Database layer
│   └── model/         # Structs and types
├── go.mod
├── demo_render.go     # Live demo script for Render
└── README.md
```
```bash
⚙️ Environment Variables
Local .env file:
PORT=8080
GIN_MODE=release
JWT_SECRET=supersecretkey
DATABASE_URL=              # Leave empty for SQLite locally
Render (Production)
PORT=10000
GIN_MODE=release
JWT_SECRET=supersecretkey
DATABASE_URL=postgres://username:password@host:5432/dbname
```
```bash
🌐 API Endpoints
Method	URL	Body	Description
POST	/signup	{ "email": "...", "password": "..." }	Register new user
POST	/login	{ "email": "...", "password": "..." }	Login and get JWT token
GET	/tasks	—	List all tasks for logged-in user
POST	/tasks	{ "title": "Task Name" }	Create a new task
PUT	/tasks/:id	{ "completed": true/false }	Update a task
DELETE	/tasks/:id	—	Delete a task
All /tasks endpoints require:
Authorization: Bearer <token>
```
```bash
💻 Example Usage (Live on Render)
Base URL:
https://task-tracker-5cg1.onrender.com
1️⃣ Signup a user
curl -X POST https://task-tracker-5cg1.onrender.com/signup \
-H "Content-Type: application/json" \
-d '{"email":"test@example.com","password":"password123"}'
2️⃣ Login and get JWT
curl -X POST https://task-tracker-5cg1.onrender.com/login \
-H "Content-Type: application/json" \
-d '{"email":"test@example.com","password":"password123"}'
Response:
{ "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." }
3️⃣ Create a task
curl -X POST https://task-tracker-5cg1.onrender.com/tasks \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <token>" \
-d '{"title":"Learn Go"}'
4️⃣ List tasks
curl -X GET https://task-tracker-5cg1.onrender.com/tasks \
-H "Authorization: Bearer <token>"
5️⃣ Update a task
curl -X PUT https://task-tracker-5cg1.onrender.com/tasks/1 \
-H "Content-Type: application/json" \
-H "Authorization: Bearer <token)" \
-d '{"completed": true}'
6️⃣ Delete a task
curl -X DELETE https://task-tracker-5cg1.onrender.com/tasks/1 \
-H "Authorization: Bearer <token>"
```
```bash
🧪 Demo Script
You can run demo_render.go locally to automatically:
Login to your live Render API
Fetch a JWT token
Create, list, update, and delete tasks automatically
▶️ Run the demo:
go run demo_render.go
🧠 Example output:
🔑 Logging in and fetching token...
✅ Token fetched successfully!
```
```bash
🪄 Creating tasks...
✅ Created: {"id":1,"title":"Learn Go","completed":false}
✅ Created: {"id":2,"title":"Build API","completed":false}

📋 Listing tasks:
[{"id":1,"title":"Learn Go","completed":false},{"id":2,"title":"Build API","completed":false}]

🔁 Updating task #1
✅ Updated: {"id":1,"title":"Learn Go","completed":true}

🗑 Deleting task #2
✅ Deleted task 2
```
```bash
🏁 Final list:
[{"id":1,"title":"Learn Go","completed":true}]
📦 Deployment on Render
Push your repo to GitHub.
Create a New Web Service on Render:
Runtime: Go
Branch: main
Build Command:
go build -o app ./cmd/server
Start Command:
./app
Set environment variables in the Render dashboard:
PORT, JWT_SECRET, DATABASE_URL
Deploy 🚀 — your API is now live and shareable.

🧰 Tech Stack
Go (Gin + GORM)
PostgreSQL (on Render)
JWT Authentication
REST API Architecture
Deployed via Render

🧪 Future Improvements
Switch to PostgreSQL for persistent storage in production
Add user roles and permissions
Add Swagger/OpenAPI documentation
Add unit and integration tests
Extend tasks with due dates, priorities, or categories
```
🧑‍💻 Author
Giorgos Laliotis
Task Tracker API
🔗 https://task-tracker-5cg1.onrender.com