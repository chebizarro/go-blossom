package auth_test

import (
	"encoding/base64"
	"encoding/json"
	"testing"
	"goblossom/internal/auth"
	"github.com/nbd-wtf/go-nostr"
	"github.com/stretchr/testify/assert"
)

func TestVerifyNostrEvent(t *testing.T) {
    validEvent := &nostr.Event{
        PubKey:    "validpubkey",
        Sig:       "validsig",
        Kind:      24242,
        CreatedAt: 1625155000,
        Tags:      nostr.Tags{nostr.Tag{"t", "get"}},
    }

    err := auth.VerifyNostrEvent(validEvent)
    assert.NoError(t, err)

    invalidEvent := &nostr.Event{
        PubKey:    "validpubkey",
        Sig:       "invalidsig",
        Kind:      24242,
        CreatedAt: 1625155000,
        Tags:      nostr.Tags{nostr.Tag{"t", "get"}},
    }

    err = auth.VerifyNostrEvent(invalidEvent)
    assert.Error(t, err)
}

func TestVerifyNostrEventFromHeader(t *testing.T) {
    validEvent := &nostr.Event{
        PubKey:    "validpubkey",
        Sig:       "validsig",
        Kind:      24242,
        CreatedAt: 1625155000,
        Tags:      nostr.Tags{nostr.Tag{"t", "get"}},
    }

    eventBytes, _ := json.Marshal(validEvent)
    encodedEvent := base64.StdEncoding.EncodeToString(eventBytes)

    err := auth.VerifyNostrEventFromHeader("Nostr " + encodedEvent)
    assert.NoError(t, err)
}
