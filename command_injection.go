package main

import "os/exec"

func RunUserCommand(userInput string) {
    cmd := exec.Command("sh", "-c", "echo "+userInput) // vulnerable
    cmd.Run()
}

