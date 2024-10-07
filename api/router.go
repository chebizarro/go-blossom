package api

import (
    "github.com/go-chi/chi/v5"
    "net/http"
)

// NewRouter initializes the router and applies routes for the API
func NewRouter(handler HandlerImpl) http.Handler {
    router := chi.NewRouter()

    // Add the Nostr authentication middleware globally
    router.Use(AuthMiddleware)

    // Define routes for blob operations
    router.Get("/list/{pubkey}", handler.ListPubkeyGet)
    router.Put("/mirror", handler.MirrorPut)
    router.Delete("/{sha256}", handler.SHA256Delete)
    router.Get("/{sha256}", handler.SHA256Get)
    router.Head("/{sha256}", handler.SHA256Head)
    router.Head("/upload", handler.UploadHead)
    router.Put("/upload", handler.UploadPut)

    return router
}
