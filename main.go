package main

import (
  "fmt"
  "net/http"
)

type HelloHandler struct{}

// func hello(w http.ResponseWriter, r *http.Request) {
//   fmt.Fprintf(w, "Hello!")
// }

func (h HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello!")
}


// func log(h http.HandlerFunc) http.HandlerFunc{
//   return func(w http.ResponseWriter, r *http.Request) {
//     name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
//     fmt.Println("Handler function called - " + name)
//     h(w, r)
//   }
// }

func log(h http.Handler) http.Handler {
  return http.HandlerFunc (func(w http.ResponseWriter, r *http.Request) {
    fmt.Printf("Handler called - %T\n", h)
    h.ServeHTTP (w, r)
  })
}


func protect(h http.Handler) http.Handler {
  return http.HandlerFunc (func(w http.ResponseWriter, r *http.Request) {
    // ...
    h.ServeHTTP (w, r)
  })
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  hello := HelloHandler{}
  http.Handle("/hello", protect(log(hello)))
  server.ListenAndServe()
}

// func main() {
//   server := http.Server{
//     Addr: "127.0.0.1:8080",
//   }
//   http.HandleFunc("/hello", log(hello))
//   server.ListenAndServe()
// }
