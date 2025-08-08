# Loki Messenger

This is a secure messaging application built with Go, PostgreSQL, and Svelte.

## Project Structure

The project is organized into the following directories:

-   `docs/`: Contains project documentation.
-   `scripts/`: Contains utility scripts for development and deployment.
-   `services/`: Contains the backend and frontend services.
    -   `backend-go/`: The Go backend service.
    -   `frontend-svelte/`: The Svelte frontend service.

## Getting Started

To get started with Loki Messenger, you'll need to have [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/) installed.

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/your-username/loki-messenger.git
    cd loki-messenger
    ```

2.  **Run with Docker Compose:**

    From the root `loki-messenger` directory, run the following command:

    ```bash
    docker-compose up --build
    ```

The application will be available at the following URLs:

-   **Frontend:** `http://localhost:5173`
-   **Backend:** `http://localhost:8080`