REST API
A RESTful backend API built with Go and SQLite, designed to demonstrate practical backend development, API design, database persistence, and HTTP-based resource management.
🚀 Features
RESTful API architecture
Go HTTP server
SQLite database persistence
Full CRUD operations
Partial resource updates with PATCH
JSON request and response handling
Resource-based routing
HTTP status code handling
Environment-based configuration
Automatic database/table initialization
Error handling for invalid resources
Lightweight dependency footprint
🛠️ Tech Stack
Go
SQLite
REST API
HTTP
JSON
Git & GitHub
📁 API Resources
The API currently provides the following resources:
Portfolio
GET /portfolio
Returns the professional portfolio profile.
Skills
GET    /skills
POST   /skills
GET    /skills/{id}
PUT    /skills/{id}
PATCH  /skills/{id}
DELETE /skills/{id}
The /skills resource demonstrates complete CRUD functionality.
🔌 API Endpoints
Method
Endpoint
Description
GET
/portfolio
Retrieve portfolio information
GET
/skills
Retrieve all skills
POST
/skills
Create a new skill
GET
/skills/{id}
Retrieve a specific skill
PUT
/skills/{id}
Replace/update a skill
PATCH
/skills/{id}
Partially update a skill
DELETE
/skills/{id}
Delete a skill
🧪 Example Requests
Get all skills
curl http://localhost:8090/skills
Example response:
[
  {
    "id": 1,
    "name": "SNMP Monitoring"
  }
]
Create a skill
curl -X POST http://localhost:8090/skills \
  -H "Content-Type: application/json" \
  -d '{"name":"Network Automation"}'
Get a specific skill
curl http://localhost:8090/skills/1
Update a skill with PUT
curl -X PUT http://localhost:8090/skills/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"Network Engineering"}'
Partially update a skill
curl -X PATCH http://localhost:8090/skills/1 \
  -H "Content-Type: application/json" \
  -d '{"name":"SNMP Monitoring"}'
Delete a skill
curl -X DELETE http://localhost:8090/skills/1
A successful DELETE returns:
204 No Content
💾 Database
The application uses SQLite for persistent data storage.
On startup, the API:
Opens the configured SQLite database.
Creates the database if necessary.
Creates the skills table if it does not already exist.
Starts the HTTP server.
The default database file is:
rest-api.db
The database location can also be configured using the DB_PATH environment variable.
⚙️ Configuration
The API supports environment-based configuration for deployment.
PORT
The server uses the PORT environment variable when provided.
If PORT is not set, it defaults to:
8090
DB_PATH
The SQLite database location can be configured using:
DB_PATH=rest-api.db
If DB_PATH is not set, the application uses:
rest-api.db
This allows the same application to run locally and on a cloud platform with persistent storage.
▶️ Running Locally
1. Clone the repository
git clone https://github.com/olaiyadigitech/rest-api.git
cd rest-api
2. Download dependencies
go mod download
3. Run the API
go run main.go
The API will start on:
http://localhost:8090
4. Test the API
curl http://localhost:8090/portfolio
or:
curl http://localhost:8090/skills
🏗️ Project Structure
rest-api/
├── main.go
├── go.mod
├── go.sum
├── README.md
└── .gitignore
The SQLite database and compiled binary are intentionally excluded from version control.
🔄 API Architecture
Client
  │
  │ HTTP Request
  ▼
Go HTTP Server
  │
  ├── /portfolio
  │
  └── /skills
       │
       ├── GET
       ├── POST
       ├── PUT
       ├── PATCH
       └── DELETE
              │
              ▼
          SQLite Database
🎯 Project Purpose
This project demonstrates practical experience building a backend service with Go, including:
RESTful API design
HTTP request handling
JSON serialization
CRUD operations
SQLite database integration
Resource-based routing
Error handling
Environment-based configuration
Git/GitHub workflow
It also provides a foundation for extending the API with authentication, authorization, additional resources, automated testing, API documentation, and production deployment.
🔮 Future Improvements
Potential future enhancements include:
JWT authentication
API authorization
Request validation
Automated unit and integration tests
Structured logging
Middleware
CORS configuration
API documentation with OpenAPI/Swagger
PostgreSQL support
Docker containerization
CI/CD automation
Public cloud deployment
Health-check endpoint
👨‍💻 Author
Oluwafemi Ajayi
Network Engineer | NOC Engineer | Infrastructure & Automation | Go Backend Developer
GitHub: olaiyadigitech
📄 License
This project is available for learning, demonstration, and portfolio purposes.
