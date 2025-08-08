# Loki Messenger

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Loki Messenger is a secure, real-time messaging application. This project is a monorepo containing the frontend and backend services.

## Features

*   **Real-time Messaging:** Instant messaging with WebSockets.
*   **Secure Communication:** End-to-end encryption for all messages.
*   **User Authentication:** Secure user registration and login.
*   **Scalable Architecture:** Containerized services for easy scaling.

## Technology Stack

The project is built with the following technologies:

*   **Backend:** [Go](https://golang.org/), [PostgreSQL](https://www.postgresql.org/), [gorilla/websocket](https://github.com/gorilla/websocket)
*   **Frontend:** [Svelte](https://svelte.dev/), [Vite](https://vitejs.dev/)
*   **Containerization:** [Docker](https://www.docker.com/), [Docker Compose](https://docs.docker.com/compose/)

## Project Structure

The project is organized as a monorepo:

-   `docs/`: Project documentation.
-   `scripts/`: Utility scripts for development and deployment.
-   `services/`:
    -   `backend-go/`: The Go backend service.
    -   `frontend-svelte/`: The Svelte frontend service.

## Getting Started

To get started, you'll need [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/) installed.

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/loki-messenger.git
    cd loki-messenger
    ```

2.  **Configuration:**

    The backend service requires a `JWT_SECRET` for signing authentication tokens. The `docker-compose.yml` file is configured to use the value `your-secret-key` by default. For a production environment, you should use a more secure, randomly generated key.

3.  **Run with Docker Compose:**

    From the root directory, run the following command:

    ```bash
    docker-compose up --build
    ```

    The application will be available at:

    -   **Frontend:** `http://localhost:5173`
    -   **Backend:** `http://localhost:8080`

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

1.  Fork the repository.
2.  Create your feature branch (`git checkout -b feature/your-feature`).
3.  Commit your changes (`git commit -m 'Add some feature'`).
4.  Push to the branch (`git push origin feature/your-feature`).
5.  Open a pull request.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.