# 🛒 E-Commerce Backend in Go

A simple e-commerce backend API built with Go, MySQL, and Gorilla Mux. This project demonstrates user authentication, product management, cart checkout, and order processing using RESTful APIs.

## Features

- User registration and login with JWT authentication
- Password hashing with bcrypt
- Product listing and management
- Cart checkout and order creation
- Input validation with go-playground/validator
- Database migrations with golang-migrate
- Environment variable management with godotenv

## Project Structure

```
.
├── cmd/
│   ├── api/         # API server entry point
│   └── migrate/     # Database migration entry point and migrations
├── config/          # Environment variable and config management
├── db/              # Database connection logic
├── service/
│   ├── auth/        # JWT and password utilities
│   ├── cart/        # Cart and checkout logic
│   ├── order/       # Order storage logic
│   ├── product/     # Product storage and routes
│   └── user/        # User storage and routes
├── types/           # Shared types and interfaces
├── utils/           # Utility functions (JSON, validation, etc.)
├── Makefile         # Build, test, and migration commands
└── README.md
```

## API Endpoints

- `POST /api/v1/register` — Register a new user
- `POST /api/v1/login` — Login and receive JWT token
- `GET /api/v1/products` — List all products
- `POST /api/v1/cart/checkout` — Checkout cart (JWT required)

## Getting Started

### Prerequisites

- Go 1.18+
- MySQL database

### Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/saikumaradapa/ecom.git
   cd ecom
   ```

2. **Configure environment variables:**
   - Copy `.env.example` to `.env` and update values as needed.

3. **Run database migrations:**
   ```sh
   make migrate-up
   ```

4. **Build and run the server:**
   ```sh
   make run
   ```

   The server will start on `http://localhost:8080`.

### Running Tests

```sh
make test
```

## Development

- API entry point: [`cmd/api/main.go`](cmd/main.go)
- Migration entry point: [`cmd/migrate/main.go`](cmd/migrate/main.go)
- Main configuration: [`config/env.go`](config/env.go)

## Useful Commands

- `make build` — Build the binary
- `make run` — Run the application
- `make test` — Run tests
- `make migration name=your_migration` — Create a new migration
- `make migrate-up` — Apply migrations
- `make migrate-down` — Rollback migrations

## Dependencies

- [Gorilla Mux](https://github.com/gorilla/mux)
- [go-playground/validator](https://github.com/go-playground/validator)
- [golang-migrate/migrate](https://github.com/golang-migrate/migrate)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt)
- [godotenv](https://github.com/joho/godotenv)

## License

MIT

---

Built with ❤️ and fun 🤩 by [Sai Kumar Adapa](https://github.com/saikumaradapa)
