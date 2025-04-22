package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/sessions"
)

var (
    templates *template.Template
    store     *sessions.CookieStore
)

func init() {
    templates = template.Must(template.ParseGlob("templates/*.html"))
    // Create a secure cookie store with a secret key
    store = sessions.NewCookieStore([]byte("your-secret-key-here"))
    store.Options = &sessions.Options{
        Path:     "/",
        MaxAge:   3600 * 24, // 24 hours
        HttpOnly: true,
    }
}

// AuthMiddleware checks if the user is authenticated
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        session, _ := store.Get(r, "login-session")
        if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func main() {
    r := chi.NewRouter()

    // Middleware
    r.Use(middleware.Logger)
    r.Use(middleware.Recoverer)
    r.Use(middleware.Timeout(60 * time.Second))

    // Serve static files
    r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Public routes
    r.Get("/login", showLogin)
    r.Post("/login", handleLogin)
    r.Get("/logout", handleLogout)

    // Protected routes
    r.Group(func(r chi.Router) {
        r.Use(AuthMiddleware)
        r.Get("/", homePage)
    })

    log.Println("Server starting on http://localhost:8080")
    http.ListenAndServe(":8080", r)
}

func showLogin(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "login-session")
    if auth, ok := session.Values["authenticated"].(bool); ok && auth {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }
    templates.ExecuteTemplate(w, "login.html", nil)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
    password := r.FormValue("password")

    if password == "admin" {
        session, _ := store.Get(r, "login-session")
        session.Values["authenticated"] = true
        session.Save(r, w)
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    templates.ExecuteTemplate(w, "login.html", map[string]interface{}{
        "Error": "Invalid password",
    })
}

func handleLogout(w http.ResponseWriter, r *http.Request) {
    session, _ := store.Get(r, "login-session")
    session.Values["authenticated"] = false
    session.Save(r, w)
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func homePage(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "home.html", map[string]interface{}{
        "Authenticated": true,
    })
}
