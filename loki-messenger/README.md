# Loki-Messenger

This is a secure messaging application built with Go, PostgreSQL, and SvelteKit.

## How to Run

1.  **Create a `.env` file:**
    *   Navigate to the `backend-go` directory.
    *   Create a `.env` file by copying the example: `cp .env.example .env`
    *   Update the `JWT_SECRET` in the `.env` file with a secure, private key.

2.  **Run with Docker Compose:**
    *   From the root `loki-messenger` directory, run the following command:
        ```bash
        docker-compose up --build
        ```

The application will be available at `http://localhost:8080`.