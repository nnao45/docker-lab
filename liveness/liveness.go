package main

import (
        "fmt"
        "net/http"
)

func main() {
        i := 0

        http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
                fmt.Fprint(w, "お前はもう、死んでいる")
                if i >= 1 {
                        for {
                        }
                }
                i++
        })
        http.ListenAndServe(":8080", nil)
}

