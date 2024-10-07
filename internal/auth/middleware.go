package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/nbd-wtf/go-nostr"
)

// AuthMiddleware is a middleware function that checks the Authorization header for a valid Nostr event
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Step 1: Extract the Authorization header
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
            return
        }

        // Step 2: Ensure it starts with "Nostr "
        if !strings.HasPrefix(authHeader, "Nostr ") {
            http.Error(w, "Invalid Authorization scheme", http.StatusUnauthorized)
            return
        }

        // Step 3: Extract the base64 encoded event
        encodedEvent := strings.TrimPrefix(authHeader, "Nostr ")
        decodedEventBytes, err := base64.StdEncoding.DecodeString(encodedEvent)
        if err != nil {
            http.Error(w, "Invalid base64 encoding", http.StatusUnauthorized)
            return
        }

        // Step 4: Unmarshal the event
        var event *nostr.Event
        if err := json.Unmarshal(decodedEventBytes, &event); err != nil {
            http.Error(w, "Invalid Nostr event", http.StatusUnauthorized)
            return
        }

        // Step 5: Verify the event
        if err := VerifyNostrEvent(event); err != nil {
            http.Error(w, fmt.Sprintf("Nostr event verification failed: %v", err), http.StatusUnauthorized)
            return
        }

        // Step 6: Pass the request to the next handler if everything is valid
        next.ServeHTTP(w, r)
    })
}
