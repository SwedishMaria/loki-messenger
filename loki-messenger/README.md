# Loki-Messenger

This is a secure and simple messenger application built with Go (Echo) for the backend, PostgreSQL for the database, and SvelteKit for the frontend (to be added in a later phase).

## Phase 1: Secure Backend Foundation

This phase focuses on setting up the project structure and building a secure backend foundation.

### Prerequisites

-   Docker
-   Docker Compose

### Getting Started

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/loki-messenger.git
    cd loki-messenger
    ```

2.  **Create the environment file:**

    Navigate to the `backend-go` directory and create a `.env` file by copying the example:

    ```bash
    cp backend-go/.env.example backend-go/.env
    ```

    Update the `.env` file with your desired credentials and settings.

3.  **Build and run the application:**

    From the root `loki-messenger` directory, run the following command:

    ```bash
    docker-compose up --build
    ```

    This will build the backend and database services and start the application. The backend will be accessible at `http://localhost:8080`.

### API Endpoints

-   `POST /api/auth/register`: Register a new user.
-   `POST /api/auth/login`: Log in and receive a JWT.

### Project Structure

```
loki-messenger/
├── backend-go/
│   ├── database/
│   │   ├── db.go
│   │   └── init.sql
│   ├── handlers/
│   │   └── auth_handler.go
│   ├── models/
│   │   └── user.go
│   ├── .env.example
│   ├── Dockerfile
│   ├── go.mod
│   ├── go.sum
│   └── main.go
├── frontend-svelte/
├── docker-compose.yml
└── README.md
```