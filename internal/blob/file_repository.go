package blob

import (
    "errors"
    "fmt"
    "io"
    "os"
    "path/filepath"
    "time"
)

type FileBlobRepository struct {
    baseDir string // Directory where blobs will be stored
}

// NewFileBlobRepository creates a new instance of FileBlobRepository.
func NewFileBlobRepository(baseDir string) *FileBlobRepository {
    return &FileBlobRepository{baseDir: baseDir}
}

// SaveBlob saves the blob in the file system and returns a BlobDescriptor.
func (repo *FileBlobRepository) SaveBlob(sha256 string, data io.Reader, size int64, mimeType string) (*BlobDescriptor, error) {
    blobPath := filepath.Join(repo.baseDir, sha256)
    file, err := os.Create(blobPath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    // Copy data to the file
    n, err := io.Copy(file, data)
    if err != nil {
        return nil, err
    }

    // Check if the size matches
    if n != size {
        return nil, errors.New("size mismatch when saving blob")
    }

    return &BlobDescriptor{
        URL:      fmt.Sprintf("/%s", sha256),
        Sha256:   sha256,
        Size:     n,
        Type:     mimeType,
        Uploaded: time.Now(),
    }, nil
}

// GetBlob retrieves the blob data for the given SHA-256 hash.
func (repo *FileBlobRepository) GetBlob(sha256 string) (io.Reader, *BlobDescriptor, error) {
    blobPath := filepath.Join(repo.baseDir, sha256)
    file, err := os.Open(blobPath)
    if err != nil {
        if os.IsNotExist(err) {
            return nil, nil, errors.New("blob not found")
        }
        return nil, nil, err
    }

    // Get file info for size and last modified time
    fileInfo, err := file.Stat()
    if err != nil {
        return nil, nil, err
    }

    descriptor := &BlobDescriptor{
        URL:      fmt.Sprintf("/%s", sha256),
        Sha256:   sha256,
        Size:     fileInfo.Size(),
        Uploaded: fileInfo.ModTime(),
    }

    return file, descriptor, nil
}

// DeleteBlob deletes the blob with the given SHA-256 hash.
func (repo *FileBlobRepository) DeleteBlob(sha256 string) error {
    blobPath := filepath.Join(repo.baseDir, sha256)
    if err := os.Remove(blobPath); err != nil {
        if os.IsNotExist(err) {
            return errors.New("blob not found")
        }
        return err
    }
    return nil
}

// HasBlob checks if a blob exists for the given SHA-256 hash.
func (repo *FileBlobRepository) HasBlob(sha256 string) (bool, error) {
    blobPath := filepath.Join(repo.baseDir, sha256)
    if _, err := os.Stat(blobPath); err != nil {
        if os.IsNotExist(err) {
            return false, nil
        }
        return false, err
    }
    return true, nil
}

// ListBlobsByPubKey lists blobs uploaded by a user.
func (repo *FileBlobRepository) ListBlobsByPubKey(pubKey string, since, until time.Time) ([]*BlobDescriptor, error) {
    // Implementation left to your design (e.g., metadata stored in a separate file or database)
    return nil, errors.New("not implemented")
}
