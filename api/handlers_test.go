package api_test

import (
    "context"
    "errors"
    "io"
    "strings"
    "testing"

    "goblossom/api"
    "goblossom/internal/blob"
    "goblossom/utils"

    "github.com/stretchr/testify/assert"
)

type MockBlobRepository struct {
    blobs map[string]blob.BlobDescriptor
}

func (m *MockBlobRepository) SaveBlob(sha256 string, data io.Reader, size int64, mimeType string) (*blob.BlobDescriptor, error) {
    m.blobs[sha256] = blob.BlobDescriptor{
        URL:    "http://localhost/" + sha256,
        Sha256: sha256,
        Size:   size,
        Type:   mimeType,
    }
    return &m.blobs[sha256], nil
}

func (m *MockBlobRepository) GetBlob(sha256 string) (io.Reader, *blob.BlobDescriptor, error) {
    descriptor, exists := m.blobs[sha256]
    if !exists {
        return nil, nil, errors.New("blob not found")
    }
    return strings.NewReader("file data"), &descriptor, nil
}

func (m *MockBlobRepository) DeleteBlob(sha256 string) error {
    delete(m.blobs, sha256)
    return nil
}

func (m *MockBlobRepository) ListBlobsByPubKey(pubKey string, since, until int64) ([]*blob.BlobDescriptor, error) {
    // Return a mock list of blobs
    return []*blob.BlobDescriptor{
        {Sha256: "mockhash", URL: "http://localhost/mockhash"},
    }, nil
}

func TestUploadPut(t *testing.T) {
    mockRepo := &MockBlobRepository{blobs: make(map[string]blob.BlobDescriptor)}
    handler := api.HandlerImpl{BlobRepo: mockRepo}

    req := api.UploadPutReq{
        File:        strings.NewReader("mock file data"),
        Size:        int64(100),
        XContentType: "application/pdf",
    }

    res, err := handler.UploadPut(context.Background(), req)
    assert.NoError(t, err)
    assert.Equal(t, "mockhash", res.Sha256)
    assert.Equal(t, "http://localhost/mockhash", res.Url)
}

func TestSHA256Get(t *testing.T) {
    mockRepo := &MockBlobRepository{
        blobs: map[string]blob.BlobDescriptor{
            "mockhash": {URL: "http://localhost/mockhash", Sha256: "mockhash"},
        },
    }
    handler := api.HandlerImpl{BlobRepo: mockRepo}

    params := api.SHA256GetParams{Sha256: "mockhash"}
    res, err := handler.SHA256Get(context.Background(), params)
    assert.NoError(t, err)
    assert.NotNil(t, res.Data)
    assert.Equal(t, "mockhash", res.ContentType)
}

func TestSHA256Delete(t *testing.T) {
    mockRepo := &MockBlobRepository{
        blobs: map[string]blob.BlobDescriptor{
            "mockhash": {Sha256: "mockhash"},
        },
    }
    handler := api.HandlerImpl{BlobRepo: mockRepo}

    params := api.SHA256DeleteParams{Sha256: "mockhash"}
    res, err := handler.SHA256Delete(context.Background(), params)
    assert.NoError(t, err)

    _, _, err = mockRepo.GetBlob("mockhash")
    assert.Error(t, err) // Should return an error since the blob has been deleted
}

func TestListPubkeyGet(t *testing.T) {
    mockRepo := &MockBlobRepository{}
    handler := api.HandlerImpl{BlobRepo: mockRepo}

    params := api.ListPubkeyGetParams{Pubkey: "mockpubkey"}
    res, err := handler.ListPubkeyGet(context.Background(), params)
    assert.NoError(t, err)
    assert.Len(t, res, 1)
    assert.Equal(t, "mockhash", res[0].Sha256)
}
