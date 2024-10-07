package blob

import (
    "errors"
    "io"
    "strings"
)

// MockBlobRepository simulates the behavior of a real BlobRepository for testing purposes.
type MockBlobRepository struct {
    Blobs map[string]BlobDescriptor // A map to store blobs by their SHA-256 hash
}

// NewMockBlobRepository creates a new instance of the MockBlobRepository.
func NewMockBlobRepository() *MockBlobRepository {
    return &MockBlobRepository{
        Blobs: make(map[string]BlobDescriptor),
    }
}

// SaveBlob simulates saving a blob to the repository.
func (m *MockBlobRepository) SaveBlob(sha256 string, data io.Reader, size int64, mimeType string) (*BlobDescriptor, error) {
    if _, exists := m.Blobs[sha256]; exists {
        return nil, errors.New("blob already exists")
    }

    // Simulate saving the blob by storing its descriptor
    descriptor := BlobDescriptor{
        URL:    "http://localhost/" + sha256,
        Sha256: sha256,
        Size:   size,
        Type:   mimeType,
    }
    m.Blobs[sha256] = descriptor
    return &descriptor, nil
}

// GetBlob simulates retrieving a blob from the repository.
func (m *MockBlobRepository) GetBlob(sha256 string) (io.Reader, *BlobDescriptor, error) {
    descriptor, exists := m.Blobs[sha256]
    if !exists {
        return nil, nil, errors.New("blob not found")
    }

    // Simulate returning the blob data as a string reader
    data := strings.NewReader("mock file data")
    return data, &descriptor, nil
}

// DeleteBlob simulates deleting a blob from the repository.
func (m *MockBlobRepository) DeleteBlob(sha256 string) error {
    if _, exists := m.Blobs[sha256]; !exists {
        return errors.New("blob not found")
    }
    delete(m.Blobs, sha256)
    return nil
}

// HasBlob simulates checking if a blob exists in the repository.
func (m *MockBlobRepository) HasBlob(sha256 string) (bool, error) {
    _, exists := m.Blobs[sha256]
    return exists, nil
}

// ListBlobsByPubKey simulates listing blobs uploaded by a specific user.
func (m *MockBlobRepository) ListBlobsByPubKey(pubKey string, since, until int64) ([]*BlobDescriptor, error) {
    // In a real implementation, you would filter blobs based on the pubKey, since, and until fields.
    // For simplicity, we just return all blobs.
    blobList := []*BlobDescriptor{}
    for _, descriptor := range m.Blobs {
        blobList = append(blobList, &descriptor)
    }
    return blobList, nil
}

// MirrorBlob simulates mirroring a blob from another server.
func (m *MockBlobRepository) MirrorBlob(url string) error {
    // For the mock implementation, we simply assume that the mirror was successful.
    sha256 := "mockmirroredhash"
    descriptor := BlobDescriptor{
        URL:    url,
        Sha256: sha256,
        Size:   1234,
        Type:   "application/pdf",
    }
    m.Blobs[sha256] = descriptor
    return nil
}
