package main

import (
    "fmt"
    "net/http"
    "os"
    "time"
)

func main() {
    url := os.Getenv("HEALTHCHECK_URL")
    if url == "" {
        url = "http://localhost:13133/"
    }
    timeout := 3 * time.Second
    client := http.Client{Timeout: timeout}

    resp, err := client.Get(url)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
    defer resp.Body.Close()

    if resp.StatusCode == 200 {
        os.Exit(0)
    }
    fmt.Fprintf(os.Stderr, "Unexpected status code: %d\n", resp.StatusCode)
    os.Exit(1)
}
