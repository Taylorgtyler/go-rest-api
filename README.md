# Go Starter: JWT-Authenticated To-Do App

Welcome to the Go Starter repository for a JWT-authenticated To-Do application. This boilerplate code is built to kickstart your journey in creating a robust, secure, and fully functional To-Do application. It is designed with modern best practices, focusing on JWT authentication and basic CRUD operations to manage to-do items.

## Features

- **JWT Authentication**: Securely manage user sessions using JSON Web Tokens (JWT). Sign up, log in, and secure your endpoints effortlessly.
- **User Registration and Login**: Ready-to-use user registration and login endpoints with hashed passwords.
- **To-Do Operations**:
    - Create new to-do items.
    - Read, update, and delete existing to-do items.
- **Structured Code**: Organized code structure that ensures separation of concerns and modularity, making the codebase easy to navigate and maintain.

## Getting Started

### Prerequisites

- Make sure you have Go installed on your machine.
- A MySQL database setup and configured.

### Installation

1. **Clone the Repository:**
    ```bash
    git clone https://github.com/yourusername/go-starter-todo.git
    ```

2. **Navigate to the Project Directory:**
    ```bash
    cd go-starter-todo
    ```

3. **Install Dependencies:**
    - This project does not depend on external dependencies apart from the standard library, as it uses Go modules.

### Configuration

- Update the database configuration by setting the relevant environment variables:
    - `DB_HOST`
    - `DB_PORT`
    - `DB_USER`
    - `DB_PASS`
    - `DB_NAME`

### Running the Application

1. **Start the Application:**
    ```bash
    go run main.go
    ```

2. **Access the Application:**
    - Open your browser or API client like Postman and access `http://localhost:8080`.

## Usage

### Endpoints

- **User Registration**: `POST /register`
- **User Login**: `POST /login`
- **Create To-Do Item**: `POST /activity`
- **Update To-Do Item**: `PUT /activity/:id`
- **Delete To-Do Item**: `DELETE /activity/:id`
- **Get To-Do Items**: `GET /activity`

### Authenticating Requests

- For endpoints that require authentication, make sure to include the JWT token in the Authorization header as a Bearer token.

## Contributing

Feel free to fork the project, make changes, and submit pull requests. Your contributions are welcome!

## License

This project is open-source and available under the MIT License.
