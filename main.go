package main
import (
  "net/http"
  "strings"
)
func status(w http.ResponseWriter, r *http.Request) {
  message := "OK"
  w.Write([]byte(message))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}

func main() {
  http.HandleFunc("/_status", status)
  http.HandleFunc("/", sayHello)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    panic(err)
  }
}
