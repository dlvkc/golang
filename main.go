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
    // 设置根路径的处理程序
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // 将 "Hello, world" 写入响应
        w.Write([]byte("Hello, world"))

        // 执行 entrypoint.sh 脚本并输出结果
        cmd := exec.Command("/bin/bash", "./entrypoint.sh")
        output, err := cmd.Output()

        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(string(output))
    })

    // 设置 /list 的处理程序
    http.HandleFunc("/list", func(w http.ResponseWriter, r *http.Request) {
        cmd := "cat list"
        output, err := exec.Command("bash", "-c", cmd).CombinedOutput()
        if err != nil {
            errorMsg := fmt.Sprintf("<h2>Error executing command:</h2><p>%s</p>", err)
            w.Write([]byte(errorMsg))
            return
        }
        outputString := strings.ReplaceAll(string(output), "n", "\n")
        htmlOutput := fmt.Sprintf("<h2>%s:</h2><p>%s</p>", cmd, outputString)
        w.Write([]byte(htmlOutput))
    })

    // 启动 HTTP 服务器并监听端口
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.ListenAndServe(":"+port, nil)
}
