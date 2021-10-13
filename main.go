package main

import (
"fmt"
"os"
"runtime"
)

func main() {
        fmt.Println("Hi,go")
        fmt.Println(runtime.GOARCH)
        name, err := os.Hostname()
        if err == nil {
                fmt.Println(name)
        }
}
