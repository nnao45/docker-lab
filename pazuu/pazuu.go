package main

import (
        "fmt"
        "io/ioutil"
        "net/http"
        "time"
)

func バルス(filename string) string {
        buff, err := ioutil.ReadFile(filename)
        if err != nil {
                return "目がああああああああああああああああ"
        }
        return string(buff)
}

func handler(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, バルス("/go/src/buls.txt"))
}

func main() {
        http.HandleFunc("/", handler)
        time.Sleep(10 * time.Second)
        http.ListenAndServe(":8080", nil)
}

