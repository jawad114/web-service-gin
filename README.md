# 🔥 GoTasker - Task Management REST API

A mini Trello/To-Do app backend built with Go, Gin, and SQLite.

## 🚀 Features

- **User Authentication**: Register & Login with JWT tokens
- **Password Security**: Bcrypt hashing for secure password storage
- **Task Management**: Full CRUD operations for tasks
- **Authorization**: Users can only access their own tasks
- **Database**: SQLite with GORM ORM
- **RESTful API**: Clean, well-structured endpoints

## 🛠️ Tech Stack

- **Go** - Main backend language
- **Gin** - HTTP web framework
- **GORM** - ORM for database operations
- **SQLite** - Database (easy local development)
- **JWT** - Authentication tokens
- **Bcrypt** - Password hashing

## 📦 Installation & Setup

1. **Clone and navigate to the project:**

   ```bash
   cd gotasker
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Run the application:**

   ```bash
   go run main.go
   ```

4. **The API will be available at:** `http://localhost:8080`

## 📚 API Documentation

### Authentication Endpoints

#### Register User

```http
POST /api/register
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "password123"
}
```

#### Login User

```http
POST /api/login
Content-Type: application/json

{
  "username": "john_doe",
  "password": "password123"
}
```

### Task Endpoints (Authentication Required)

Add the JWT token to the Authorization header:

```
Authorization: Bearer <your-jwt-token>
```

#### Get All Tasks

```http
GET /api/tasks
Authorization: Bearer <token>
```

#### Get Specific Task

```http
GET /api/tasks/:id
Authorization: Bearer <token>
```

#### Create New Task

```http
POST /api/tasks
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Complete project",
  "description": "Finish the GoTasker API",
  "status": "todo",
  "priority": 2
}
```

#### Update Task

```http
PUT /api/tasks/:id
Authorization: Bearer <token>
Content-Type: application/json

{
  "title": "Updated task title",
  "status": "in-progress",
  "priority": 3
}
```

#### Delete Task

```http
DELETE /api/tasks/:id
Authorization: Bearer <token>
```

## 📊 Task Status Values

- `todo` - Task is pending
- `in-progress` - Task is being worked on
- `done` - Task is completed

## 🔧 Project Structure

```
gotasker/
├── main.go              # Application entry point
├── models/              # Database models
│   ├── user.go         # User model
│   └── task.go         # Task model
├── handlers/            # HTTP request handlers
│   ├── user.go         # User authentication handlers
│   └── task.go         # Task CRUD handlers
├── middleware/          # Middleware functions
│   └── auth.go         # JWT authentication middleware
├── database/            # Database configuration
│   └── database.go     # Database connection and setup
└── README.md           # This file
```

## 🧪 Testing the API

### 1. Register a new user:

```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'
```

### 2. Login to get a token:

```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

### 3. Create a task (use the token from login):

```bash
curl -X POST http://localhost:8080/api/tasks \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <your-token>" \
  -d '{
    "title": "Learn Go",
    "description": "Study Go programming language",
    "status": "todo",
    "priority": 2
  }'
```

### 4. Get all tasks:

```bash
curl -X GET http://localhost:8080/api/tasks \
  -H "Authorization: Bearer <your-token>"
```

## 🔒 Security Features

- **Password Hashing**: All passwords are hashed using bcrypt
- **JWT Authentication**: Secure token-based authentication
- **User Isolation**: Users can only access their own tasks
- **Input Validation**: Request validation using Gin's binding

## 🚀 Future Enhancements

- [ ] Add categories/projects for tasks
- [ ] Real-time updates using WebSockets
- [ ] Docker containerization
- [ ] PostgreSQL support
- [ ] Task search and filtering
- [ ] Task due date reminders
- [ ] User profile management

## 📝 License

This project is open source and available under the MIT License.
