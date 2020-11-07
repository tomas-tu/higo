package main

import (
"fmt"
"os"
"net/http"
"log"
)

func main() {
        http.HandleFunc("/foo", higoHandler)
        log.Fatal(http.ListenAndServe(":8088", nil))
}

func higoHandler(w http.ResponseWriter, r *http.Request) {       
        name, err := os.Hostname()
        if err != nil {
                panic(err.Error())
        }       
        fmt.Fprintln(w, "higo, "+name)
}
