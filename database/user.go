package database

import "os"

var Users map[string]Passwords

type Passwords struct {
    Password string `json:"password"`
}

func Init() {
    Users = make(map[string]Passwords)

    passwords := Passwords{
        Password: os.Getenv("VERCEL_PASSWORD"),
    }
    Users[os.Getenv("VERCEL_EMIAL")] = passwords
}
