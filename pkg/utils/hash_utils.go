package utils

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "io"
    "os"
)

// HashFile calculates the SHA-256 hash of a file at the given path.
func HashFile(filePath string) (string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return "", fmt.Errorf("failed to open file: %v", err)
    }
    defer file.Close()

    // Initialize a SHA-256 hash generator
    hash := sha256.New()

    // Copy the file contents into the hash generator
    if _, err := io.Copy(hash, file); err != nil {
        return "", fmt.Errorf("failed to hash file: %v", err)
    }

    // Convert the hash to a string
    return hex.EncodeToString(hash.Sum(nil)), nil
}

// HashReader calculates the SHA-256 hash of data from an io.Reader.
// This is useful for streaming data, like HTTP uploads.
func HashReader(reader io.Reader) (string, error) {
    // Initialize a SHA-256 hash generator
    hash := sha256.New()

    // Copy the reader data into the hash generator
    if _, err := io.Copy(hash, reader); err != nil {
        return "", fmt.Errorf("failed to hash reader data: %v", err)
    }

    // Convert the hash to a string
    return hex.EncodeToString(hash.Sum(nil)), nil
}

// VerifyHash compares the calculated SHA-256 hash of data from an io.Reader with a provided hash.
func VerifyHash(reader io.Reader, expectedHash string) (bool, error) {
    calculatedHash, err := HashReader(reader)
    if err != nil {
        return false, err
    }

    return calculatedHash == expectedHash, nil
}

// HashBytes calculates the SHA-256 hash of a byte array.
func HashBytes(data []byte) string {
    hash := sha256.Sum256(data)
    return hex.EncodeToString(hash[:])
}

// HashString calculates the SHA-256 hash of a string.
func HashString(data string) string {
    return HashBytes([]byte(data))
}
