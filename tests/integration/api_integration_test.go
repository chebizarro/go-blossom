package api_test

import (
    "bytes"
    "io"
    "goblossom/api"
    "goblossom/internal/blob"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/go-chi/chi/v5"
    "github.com/stretchr/testify/assert"
)

func TestUploadPutIntegration(t *testing.T) {
    mockRepo := &blob.MockBlobRepository{Blobs: make(map[string]blob.BlobDescriptor)}
    handler := api.HandlerImpl{BlobRepo: mockRepo}

    router := chi.NewRouter()
    router.Put("/upload", handler.UploadPut)

    reqBody := bytes.NewBuffer([]byte("mock file data"))
    req := httptest.NewRequest("PUT", "/upload", reqBody)
    req.Header.Set("Content-Type", "application/pdf")
    req.Header.Set("Content-Length", "100")
    req.Header.Set("X-SHA-256", "mockhash")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

func TestSHA256GetIntegration(t *testing.T) {
    mockRepo := &blob.MockBlobRepository{
        Blobs: map[string]blob.BlobDescriptor{
            "mockhash": {Sha256: "mockhash", URL: "http://localhost/mockhash"},
        },
    }
    handler := api.HandlerImpl{BlobRepo: mockRepo}

    router := chi.NewRouter()
    router.Get("/{sha256}", handler.SHA256Get)

    req := httptest.NewRequest("GET", "/mockhash", nil)
    w := httptest.NewRecorder()

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    body, _ := io.ReadAll(w.Body)
    assert.Contains(t, string(body), "file data")
}

func TestSHA256DeleteIntegration(t *testing.T) {
    mockRepo := &blob.MockBlobRepository{
        Blobs: map[string]blob.BlobDescriptor{
            "mockhash": {Sha256: "mockhash"},
        },
    }
    handler := api.HandlerImpl{BlobRepo: mockRepo}

    router := chi.NewRouter()
    router.Delete("/{sha256}", handler.SHA256Delete)

    req := httptest.NewRequest("DELETE", "/mockhash", nil)
    w := httptest.NewRecorder()

    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)

    _, _, err := mockRepo.GetBlob("mockhash")
    assert.Error(t, err) // Blob should be deleted
}
