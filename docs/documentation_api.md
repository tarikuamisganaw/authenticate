# Task Management API Documentation
This is the link for the postman api documenation:
https://documenter.getpostman.com/view/37345989/2sA3rzKYbas

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
### Users
-**Get /api/users**

## Authentication and Authorization
- Use the `Authorization` header to send the JWT token.
  ```http
  Authorization: Bearer <jwt_token_string>
