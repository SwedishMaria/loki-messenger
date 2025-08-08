# Loki Messenger

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![Discord](https://img.shields.io/discord/YOUR_DISCORD_SERVER_ID?logo=discord&logoColor=white)](https://discord.gg/lindqviist)

Loki Messenger is a cutting-edge, secure messaging application designed for absolute privacy and anonymity. This project leverages a sophisticated architecture, routing all traffic through the Tor network and securing all communications with the Signal Protocol. This combination makes it one of the most secure and privacy-respecting messengers available.

## Security and Privacy

Loki Messenger is built on a foundation of uncompromising security. Here’s how we protect your privacy:

*   **End-to-End Encryption:** We use the **Signal Protocol**, the same strong encryption that powers Signal Messenger, to ensure that your conversations are private. Only the sender and receiver can read the messages—no one else, not even us.
*   **Anonymity with Tor:** All traffic is routed through the **Tor network**, providing a powerful layer of anonymity. This makes it extremely difficult for anyone to trace your communications or know who you are talking to.
*   **No Phone Number Required:** Get started without linking your account to a phone number.
*   **Minimal Metadata:** We are committed to collecting as little metadata as possible.

## Technology Stack

The project is built with a focus on security, performance, and modern development practices.

*   **Backend:**
    *   **Go:** A fast, statically typed, compiled language that is well-suited for building high-performance network services.
    *   **PostgreSQL:** A powerful, open-source object-relational database system.
    *   **gorilla/websocket:** A widely used WebSocket implementation for Go.
*   **Frontend:**
    *   **Svelte:** A radical new approach to building user interfaces.
    *   **Vite:** A next-generation frontend tooling that provides an extremely fast development environment.
    *   **@privacyresearch/libsignal-protocol-typescript:** The official TypeScript implementation of the Signal Protocol.
*   **Containerization:**
    *   **Docker & Docker Compose:** For creating a consistent, isolated, and reproducible development environment.

## Project Structure

The project is organized as a monorepo:

-   `docs/`: Project documentation.
-   `scripts/`: Utility scripts.
-   `services/`:
    -   `backend-go/`: The Go backend service.
    -   `frontend-svelte/`: The Svelte frontend service.

## Getting Started

You'll need [Docker](https://www.docker.com/get-started) and [Docker Compose](https://docs.docker.com/compose/install/) installed.

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/loki-messenger.git
    cd loki-messenger
    ```
2.  **Run with Docker Compose:**
    ```bash
    docker-compose up --build
    ```

The application will be available at:

-   **Frontend:** `http://localhost:5173`
-   **Backend:** `http://localhost:8080`

## Author

This project was summoned into existence by **Maria**, a cute puppygirl from Sweden! ૮ ・ﻌ・ა

Got questions, suggestions, bug reports, or just wanna chat? Hit me up! I don't bite... much. You can find me here:

*   **Website:** [https://puppymaria.com/](https://puppymaria.com/)
*   **Discord:** `lindqviist`

## Donations

The author, Maria, is currently battling cancer. If you find Loki Messenger valuable, please consider supporting the fight against cancer by donating to the Swedish Cancer Foundation (Cancerfonden). Your contribution can make a real difference.

[**Donate to Cancerfonden**](https://www.cancerfonden.se/stod-oss/ge-en-gava)

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.