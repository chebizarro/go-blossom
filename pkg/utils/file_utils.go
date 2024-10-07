package utils

import (
    "fmt"
    "strconv"
    "strings"
)

// ParseFileSize converts a string like "50MB" to an integer representing the size in bytes.
func ParseFileSize(sizeStr string) (int64, error) {
    units := map[string]int64{
        "KB": 1024,
        "MB": 1024 * 1024,
        "GB": 1024 * 1024 * 1024,
    }
    for unit, multiplier := range units {
        if strings.HasSuffix(sizeStr, unit) {
            base, err := strconv.ParseInt(strings.TrimSuffix(sizeStr, unit), 10, 64)
            if err != nil {
                return 0, fmt.Errorf("invalid size: %v", err)
            }
            return base * multiplier, nil
        }
    }
    return 0, fmt.Errorf("unknown size unit")
}

// IsAllowedMimeType checks if a MIME type is in the allowed list
func IsAllowedMimeType(mimeType string, allowed []string) bool {
    for _, allowedType := range allowed {
        if mimeType == allowedType {
            return true
        }
    }
    return false
}
