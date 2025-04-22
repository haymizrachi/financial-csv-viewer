# Login Server

A simple and secure Go web application that implements password-based authentication with session management.

## Features

- Single password authentication system
- Secure session management using cookies
- Protected routes with middleware authentication
- Clean and responsive UI
- 24-hour session duration
- Secure cookie storage

## Tech Stack

- **Backend**: Go
- **Router**: [Chi](https://github.com/go-chi/chi)
- **Session Management**: [Gorilla Sessions](https://github.com/gorilla/sessions)
- **Template Engine**: Go's built-in HTML templating
- **Frontend**: HTML & CSS

## Project Structure

```
login-server/
├── main.go        # Main application file
├── templates/     # HTML templates
│   ├── login.html # Login page template
│   └── home.html  # Protected home page template
└── static/        # Static assets
    └── styles.css # CSS styles
```

## Key Components

- **Authentication Middleware**: Protects routes by validating session cookies
- **Session Management**: Uses secure cookies with 24-hour expiration
- **Route Handlers**:
  - `/login` - Handles login form display and authentication
  - `/logout` - Manages user logout and session clearing
  - `/` - Protected home page (requires authentication)

## Getting Started

1. Ensure you have Go installed on your system
2. Clone this repository
3. Run the server:
   ```bash
   go run main.go
   ```
4. Access the application at `http://localhost:8080`
5. Use the password "admin" to log in

## Security Features

- HttpOnly session cookies
- Secure password validation
- Protected route middleware
- Session-based authentication
- Automatic redirect to login for unauthorized access
