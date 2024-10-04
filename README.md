# URI (URL Shortener)

![URI Img](assets/URI.png)

**DEMO**:  [HERE](https://clever-youth-production.up.railway.app/)

This is a brief description of my project.

URL Shortener is a web application that allows users to shorten long URLs into shorter, more shareable URLs. This project is built using Go with the Fiber framework and utilizes Redis for caching to enhance performance.

## Technologies Used

- **Programming Language**: Go
- **Framework**: Fiber
- **Database Caching**: Redis
- **Template Engine**: HTML
- **CSS Framework**: Tailwind CSS

## Features

- **Shorten URL**: Users can enter a long URL and receive the generated shortened URL.
- **Caching with Redis**: Utilizes Redis to store and retrieve already shortened URLs, speeding up access and reducing the load on the main database.
  (NOTE: The application in this repo is not yet integrated with the main DB.)

## Prerequisites

Before running this application, ensure you have installed:

- Go (version 1.18 or higher)
- Redis
- PostgreSQL (if used as the main storage)
- Go dependencies

## Installation

1. **Clone this repository**:

   ```bash
   git clone https://github.com/Rakhulsr/URI.git

   ```

2. **Create a .env file to store your environment variables. Example:**:

- REDIS_HOST=
- REDIS_PORT=
- REDIS_USERNAME=
- REDIS_PASSWORD=
- REDIS_URL=redis://localhost:6379
- BASE_URL=localhost:8080
- DATABASE*URL=postgres://username:password@localhost:5432/database_name (\_optional*)

3. **Install dependencies:**
   ```bash
   go mod tidy
   ```
4. **Migrate the database (If you are using a migration tool, add this step according to the tool you choose).**
5. **Run the application:**
   ```bash
    go run main.go
   ```

## Usage

1. Enter the long URL you want to shorten in the provided form.
2. Click the "Submit" button.
3. The shortened URL will be displayed. You can copy the URL.
