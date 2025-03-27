package main

import (
          "fmt"
          "os/user"
)

func main() {
        u, err := user.Current()
        if err != nil {
                fmt.Println("Cannot get current user:", err)
                os.Exit(1)
        }

        fmt.Printf("Hello %s, welcome!\n", u.Username)
}
