package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "os/exec"
)

func main() {
    // 设置根路径的处理程序
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // 将 "Hello, world" 写入响应
        w.Write([]byte("Hello, world"))

        // 执行 test.sh 脚本并输出结果
        cmd := exec.Command("/bin/sh", "./entrypoint.sh")
        output, err := cmd.Output()

        if err != nil {
            log.Fatal(err)
        }

        fmt.Println(string(output))
    })

    // 启动 HTTP 服务器并监听端口
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    http.ListenAndServe(":"+port, nil)
}
