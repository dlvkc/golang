package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
    "strings"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if r := recover(); r != nil {
                log.Printf("Recovering from panic: %v", r)
                http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
            }
        }()

        // 执行 entrypoint.sh 脚本并输出结果
        cmd := exec.Command("/bin/bash", "./entrypoint.sh")
        output, err := cmd.Output()
        if err != nil {
            log.Fatal(err)
        }

        // Write response at the end
        defer func() {
            if w.err == nil {
                w.Write([]byte("Hello, world"))
            }
        }()

        // Prepend output to response
        w.Write(output)
    })

    http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
        cmd := "cat list"
        output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
        if err != nil {
            errorMsg := fmt.Sprintf("<h2>Error executing command:</h2><p>%s</p>", err)
            w.Write([]byte(errorMsg))
            return
        }
        outputString := strings.ReplaceAll(string(output), "n", "n")
        htmlOutput := fmt.Sprintf("<h2>%s:</h2><p>%s</p>", cmd, outputString)
        w.Write([]byte(htmlOutput))
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.ListenAndServe(":"+port, nil)
}
