# Task Management API Documentation

## Endpoints

### User Registration
- **POST /register**
  - Request Body:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Response:
    ```json
    {
      "message": "User registered successfully"
    }
    ```

### User Login
- **POST /login**
  - Request Body:
    ```json
    {
      "username": "string",
      "password": "string"
    }
    ```
  - Response:
    ```json
    {
      "token": "jwt_token_string"
    }
    ```

### Protected Endpoints
#### Tasks
- **GET /api/tasks**
- **GET /api/tasks/:id**
- **POST /api/tasks**
- **PUT /api/tasks/:id**
- **DELETE /api/tasks/:id**

### Admin Endpoints
- **Admin-specific routes**

## Authentication and Authorization
- Use the `Authorization` header to send the JWT token.
  ```http
  Authorization: Bearer <jwt_token_string>
