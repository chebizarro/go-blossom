package api

import (
    "goblossom/internal/auth"
    "net/http"
)

// AuthMiddleware enforces Nostr authentication for protected routes
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }

        // Verify the Nostr event
        err := auth.VerifyNostrEventFromHeader(authHeader)
        if err != nil {
            http.Error(w, "Invalid Nostr event: "+err.Error(), http.StatusUnauthorized)
            return
        }

        // If the event is valid, proceed with the next handler
        next.ServeHTTP(w, r)
    })
}
