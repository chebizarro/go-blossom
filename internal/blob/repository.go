package blob

import (
    "io"
    "time"
)

type BlobDescriptor struct {
    URL      string    // The URL for accessing the blob
    Sha256   string    // The SHA-256 hash of the blob
    Size     int64     // Size of the blob in bytes
    Type     string    // MIME type of the blob (optional)
    Uploaded time.Time // Time when the blob was uploaded
}

// BlobRepository defines an interface for handling blob operations.
type BlobRepository interface {
    // SaveBlob saves a blob and returns a descriptor for the saved blob.
    SaveBlob(sha256 string, data io.Reader, size int64, mimeType string) (*BlobDescriptor, error)

    // GetBlob retrieves the blob data for the given SHA-256 hash.
    GetBlob(sha256 string) (io.Reader, *BlobDescriptor, error)

    // DeleteBlob deletes the blob with the given SHA-256 hash.
    DeleteBlob(sha256 string) error

    // HasBlob checks if a blob exists for the given SHA-256 hash.
    HasBlob(sha256 string) (bool, error)

    // ListBlobsByPubKey lists all blobs associated with a specific public key.
    ListBlobsByPubKey(pubKey string, since, until time.Time) ([]*BlobDescriptor, error)
}
