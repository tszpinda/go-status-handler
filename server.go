package main

import (
    "net/http"
    "regexp"
    "io"
    "os"
    "bufio"
    "log"
)

var contentTypes = map[string]bool{
                            "application/json": true,
                            "text/json": true,
                            "application/csv": true,
                            "text/csv": true,
                            "text/html": true,}
var regex = regexp.MustCompile(".*/")

func handleAll(w http.ResponseWriter, r *http.Request) {
        contentType := r.Header.Get("Content-Type")
        if !contentTypes[contentType] {
            http.Error(w, "Unsupported content type", http.StatusNotAcceptable)
            return
        }

        status := r.Header.Get("x-status-code")
        if(status == "") {
            status = "404"
        }
        filename := status + "." + regex.ReplaceAllString(contentType, "")

        f, _ := os.Open(filename)
        defer f.Close()
        b := bufio.NewReader(f)

        if _, err := io.Copy(w, b); err != nil {
             log.Fatal(err)
        }
}

func main() {
    http.HandleFunc("/", handleAll)
    http.ListenAndServe(":8080", nil)
}
