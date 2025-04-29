# Go Gin App

## Overview
This project is a RESTful API built using the Gin web framework in Go. It provides user authentication via phone number and OTP (One-Time Password) and includes basic user management functionalities.

## Project Structure
```
go-gin-app
├── cmd
│   └── main.go               # Entry point of the application
├── config
│   └── config.go             # Configuration settings
├── controllers
│   ├── auth_controller.go     # Handles authentication logic
│   └── user_controller.go     # Manages user-related operations
├── middlewares
│   └── jwt_middleware.go      # JWT token validation middleware
├── models
│   └── user.go                # User model definition
├── routes
│   ├── auth_routes.go         # Authentication routes
│   └── user_routes.go         # User management routes
├── services
│   ├── auth_service.go        # Business logic for authentication
│   └── user_service.go        # Business logic for user management
├── utils
│   ├── otp_generator.go        # OTP generation and validation
│   └── response_util.go        # Utility functions for API responses
├── go.mod                      # Module definition and dependencies
├── go.sum                      # Dependency checksums
└── README.md                   # Project documentation
```

## Setup Instructions
1. **Clone the repository:**
   ```
   git clone https://github.com/yourusername/go-gin-app.git
   cd go-gin-app
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Run the application:**
   ```
   go run cmd/main.go
   ```

## API Endpoints

### Authentication
- **Sign Up**
  - **Endpoint:** `POST /api/auth/signup`
  - **Request Body:** `{ "phone_number": "string" }`
  - **Response:** OTP sent to the provided phone number.

- **Sign In**
  - **Endpoint:** `POST /api/auth/signin`
  - **Request Body:** `{ "phone_number": "string", "otp": "string" }`
  - **Response:** JWT token if successful.

### User Management
- **Get User Details**
  - **Endpoint:** `GET /api/user`
  - **Headers:** `Authorization: Bearer <token>`
  - **Response:** User details.

- **Update User Information**
  - **Endpoint:** `PUT /api/user`
  - **Headers:** `Authorization: Bearer <token>`
  - **Request Body:** `{ "phone_number": "string", "other_fields": "value" }`
  - **Response:** Updated user details.

## Usage Examples
- To sign up a user, send a POST request to `/api/auth/signup` with the user's phone number.
- To sign in, send a POST request to `/api/auth/signin` with the phone number and the received OTP.
- Use the JWT token received upon signing in to access protected routes.

## License
This project is licensed under the MIT License. See the LICENSE file for more details.