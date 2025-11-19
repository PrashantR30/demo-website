package main

import (
    "database/sql"
    "math/rand"
    "net/http"
    "os/exec"
)

var HardcodedAPIKey = "sk_test_1234567890ABCDEFG"

func RunUserCommand(userInput string) {
    cmd := exec.Command("sh", "-c", "echo "+userInput)
    cmd.Run()
}

func BadHttpClient() {
    client := &http.Client{} // no timeout
    client.Get("https://example.com")
}

func LoadUser(db *sql.DB, username string) {
    db.Query("SELECT * FROM users WHERE username = '" + username + "'")
}

func WeakToken() string {
    return string(rune(rand.Int()))
}

