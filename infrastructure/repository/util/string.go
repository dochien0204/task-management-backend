package util

import "crypto/rand"

func GenerateRandomString(n int) (string, error) {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%&*"
    b := make([]byte, n) // Create a slice to hold the bytes.

    if _, err := rand.Read(b); err != nil {
        return "", err // Return the error if rand.Read fails.
    }

    for i := 0; i < n; i++ {
        b[i] = charset[b[i]%byte(len(charset))] // Convert each byte to a corresponding character from the charset.
    }

    return string(b), nil
}