# Backend

Welcome to the backend documentation, where you will find everything related to the backend side of the project.

If you haven't already, check out the POC of the backend and the comparative study documentation  [*here*](./POC.md#backend).

## Gin Framework and Architecture
For the backend, we chose the Gin framework, a high-performance HTTP web framework for Go. Its lightweight nature, built-in middleware, and focus on simplicity made it an ideal choice for our project. Gin allows us to handle requests efficiently while maintaining clean and readable code.

### Key Features of Gin in Our Project
- ### Routing:
    Gin provides fast and flexible routing with support for route groups, which we use to organize endpoints by functionality (e.g., /auth, /users, /resources).
- ### Middleware:
    Middleware such as logging, authentication, and error handling are seamlessly integrated to ensure a robust and secure API.
- ### Validation:
    Using binding and validator tags in our struct definitions, we validate incoming request data effortlessly.

## Directory Structure
To ensure the backend remains organized and scalable, we adopted a modular architecture with the following structure:

```bash
backend/
├── api/                # API-specific logic and handlers
├── controllers/        # Request handlers for routes
├── database/           # Database models and migrations
├── middlewares/        # Custom middleware (e.g., auth, logging)
├── repository/         # Data access layer
├── schemas/            # Request and response schemas
├── services/           # Business logic layer
├── toolbox/            # Utility functions and helpers
└── main.go             # Entry point for the backend application
```
This structure allows for clear separation of concerns, making the application easier to understand and maintain.

# Authentication and Authorization
We implemented JWT-based authentication to ensure secure access to protected routes. Users must provide a valid JWT token in the Authorization header to interact with these routes.

# Login and Registration
- ### Login:
    The /api/auth/login endpoint verifies user credentials and generates a JWT token upon successful authentication.
- ### Registration:
    The /api/auth/register endpoint allows users to create new accounts by providing the required details.
- ### Middleware
    We created a custom middleware to:
        Validate JWT tokens on protected routes.
        Extract user information from the token for request-specific logic.

# Error Handling
In our backend, error handling is a priority to ensure that API clients receive clear and consistent error messages

- ### Bad Request (400):
    Invalid parameters or request payloads.

- ### Unauthorized (401):
    When the user is not authenticated or the token is invalid.

- ### Forbidden (403):
    When the user is authenticated but does not have permission to access a resource.

- ### Not Found (404):
    When the requested resource does not exist.

- ### Conflict (409):
    When trying to create a resource that already exists.

We use a centralized error-handling function to map errors to appropriate HTTP codes and messages.

## Database Integration
We use GORM, an ORM library for Go, for interacting with our database. This allows us to define models, perform queries, and manage migrations with ease.

## Schemas
Each database table corresponds to a Go struct in the schemas/ directory. For example:

```go
type User struct {
	Id       uint64         `json:"id,omitempty" gorm:"primary_key;auto_increment;"`
	Name     string         `json:"name" gorm:"type:varchar(100)"`
	LastName string         `json:"lastname" gorm:"type:varchar(100)"`
	Username string         `json:"username" gorm:"type:varchar(100);unique"`
	Email    *string        `json:"email" gorm:"type:varchar(100);unique"`
	Password *string        `json:"password" gorm:"type:varchar(100)"`
	Image    string         `json:"image" gorm:"type:BYTEA"`
	IsAdmin  bool           `json:"is_admin" gorm:"type:boolean"`
	Services []ServiceToken `gorm:"many2many:user_service_tokens;constraint:OnDelete:CASCADE,OnUpdate:CASCADE;"`
}
```
## Migrations
Database migrations are automated using GORM's migration tool. We ensure all migrations are version-controlled for easy rollbacks when needed.

## Connections
Database connection details (e.g., host, port, credentials) are stored in environment variables and loaded via a configuration file.

## API Endpoints
`GET` `/about.json`: Give all the informations about the services, the differents actions / reactions.

`DELETE` `/api/user/account`: Permit to a user to delete all his data along with his account.

`PUT` `/api/user/service/logout`: Permit to a user to logout from a service.

`GET` `/api/user/workflows` : Permit to a user to get all his workflows infos.

`GET` `/api/user/services` : Permit to a user to get all his services infos.

`POST` `/api/mobile/token` : Permit to the mobile application to create or bind a user using the token given by the service.

`POST` `/api/auth/login`: Permit to a user to login.

`POST` `/api/auth/register`: Permit to a user to register.

Here serviceName is the name of the service you want to authenticate with (github / microsoft / google, ...).

`GET` `/api/serviceName/auth`: Get the url to authenticate with the service.

`GET` `/api/serviceName/callback`: Permit to a user to authenticate with a service or create an account with the service.

`POST` `/api/workflow` : Permit to a user to create a workflow with the service he want and the corresponding options.

`PUT` `/api/workflow/activation` : Permit to a user to activate or deactivate a workflow.

`GET` `/api/workflow/reactions` : Permit to a user to get all the reactions available for all his workflows.

`GET` `/api/workflow/reaction/latest?workflow_id=id` : Permit to a user to get the latest reaction of a workflow.

`DELETE` `/api/workflow` : Permit to a user to delete a workflow and the corresponding data.

`PUT` `/api/workflow` : Permit to a user to update the workflow option.

## Request and Response Formats
Example: Create a New User
Request:

```json
POST /api/workflow
{
    "action_id": 1,
    "reaction_id": 7,
    "name": "titi",
    "action_option": {"owner": "JsuisSayker", "repo": "TestAreaGithub"},
    "reaction_option": {"city_name": "Bordeaux", "language_code": "FR"}
}

Response:
json
{
    "message": "Workflow created successfully",
}
```

## Logging
We use Gin's built-in logger for tracking incoming requests, along with a custom logger to record application-specific events (e.g., errors, database queries). Logs are formatted in JSON and can be integrated with centralized logging tools for monitoring.

## Useful Links
Here are the links to the markdown documentation for the backend:

- [Brand Guidelines](https://drive.google.com/file/d/1QANYuij2kzZfJMqEzDfA68nXjhAFg6or/view?usp=drive_link)
- [POC](./POC.md#backend)
- [Organization](./Organization.md)
- [Main README](../README.md)
- [Back to top](#frontend)
