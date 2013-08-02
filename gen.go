package main

import (
    "code.google.com/p/go.crypto/ssh/terminal"
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "log"
    "os/exec"
    "strings"
    "syscall"
)

func main() {
    key, err := terminal.ReadPassword(syscall.Stdin)
    if err != nil {
        log.Fatal(err)
    }
    site, err := terminal.ReadPassword(syscall.Stdin)
    if err != nil {
        log.Fatal(err)
    }
    mac := hmac.New(sha256.New, key)
    mac.Write(site)
    pw := hex.EncodeToString(mac.Sum(nil))
    cmd := exec.Command("pbcopy")
    cmd.Stdin = strings.NewReader(pw)
    if err := cmd.Run(); err != nil {
        log.Fatal(err)
    }
}
