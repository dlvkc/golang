package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
)

func handler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/list" {
        out, err := exec.Command("cat", "list").Output()
        if err != nil {
            fmt.Fprintf(w, "<html><body>Error: %s</body></html>", err.Error())
        } else {
            fmt.Fprintf(w, "<html><body>%s</body></html>", string(out))
        }
    } else {
        fmt.Fprintln(w, "Hello, World!")
    }
}

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello, world"))

        cmd := exec.Command("/bin/sh", "./entrypoint.sh")
        output, err := cmd.Output()

        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(string(output))
    })

    http.HandleFunc("/list", handler)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.ListenAndServe(":"+port, nil)
}
