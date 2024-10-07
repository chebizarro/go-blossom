package auth

import (
    "encoding/base64"
    "encoding/json"
    "errors"
    "fmt"
    "time"
	"strings"

    "github.com/nbd-wtf/go-nostr"
)

// VerifyNostrEventFromHeader verifies a Nostr event from the Authorization header
func VerifyNostrEventFromHeader(authHeader string) error {
    if !strings.HasPrefix(authHeader, "Nostr ") {
        return errors.New("invalid Authorization scheme")
    }

    // Decode the Nostr event from the Authorization header
    encodedEvent := strings.TrimPrefix(authHeader, "Nostr ")
    decodedEventBytes, err := base64.StdEncoding.DecodeString(encodedEvent)
    if err != nil {
        return fmt.Errorf("failed to decode Nostr event: %v", err)
    }

    var event *nostr.Event
    if err := json.Unmarshal(decodedEventBytes, &event); err != nil {
        return fmt.Errorf("failed to unmarshal Nostr event: %v", err)
    }

    // Verify the Nostr event
    return VerifyNostrEvent(event)
}

// VerifyNostrEvent verifies if the provided Nostr event is valid and signed correctly by the sender's private key.
func VerifyNostrEvent(event *nostr.Event) error {
    // Step 1: Check the event kind
    if event.Kind != nostr.KindBlobs {
        return errors.New("invalid event kind")
    }

    // Step 2: Check if the event has expired
    if !isValidTimestamp(event.CreatedAt) {
        return errors.New("nostr event is expired")
    }

    // Step 3: Verify the signature using the public key
    valid, err := event.CheckSignature()
    if err != nil {
        return err
    }

    if !valid {
        return errors.New("invalid nostr event signature")
    }

    return nil
}

// isValidTimestamp checks if the event's created_at timestamp is valid and not expired
func isValidTimestamp(createdAt nostr.Timestamp) bool {
    expirationTag := getExpirationTag(createdAt)
    return time.Now().Unix() < expirationTag
}

// getExpirationTag extracts the expiration timestamp from the event's tags.
func getExpirationTag(createdAt nostr.Timestamp) int64 {
    expirationTime := time.Unix(int64(createdAt), 0).Add(24 * time.Hour) // Default 24-hour expiration
    return expirationTime.Unix()
}
