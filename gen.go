package main

import (
    "crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "log"
    "os/exec"
    "strings"
    "syscall"
    
    "github.com/zeebo/gpwd/terminal"
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
