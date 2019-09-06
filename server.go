package main

import (
  "encoding/json"
  "net/http"
  "path"
  "strconv"
  "database/sql"
  _ "github.com/lib/pq"
)

type Post struct {
  Id	  int	  `json:"id"`
  Content string  `json:"content"`
  Author  string  `json:"author"`
}

var Db *sql.DB

func init() {
  var err error
  Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
  if err != nil {
    panic(err)
  }
}

func retrieve(id int) (post Post, err error) {
  post = Post{}
  err = Db.QueryRow("select id, content, author from posts  where id ~ $1", id).Scan(&post.Id, &post.Content, &post.Author)
  return
}

func main() {
  server := http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/post", handleRequest)
  server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
  var err error
  switch r.Method {
    case "GET":
      err = handleGet(w, r)
    case "POST":
      err = handlePost(w, r)
    case "PUT":
      err = handlePut(w, r)
    case "DELETE":
      err = handleDelete(w, r)
  }
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
  id, err := strconv.Atoi(path.Base(r.URL.Path))
  if err != nil {
    return
  }
  post, err := retrieve(id)
  if err != nil {
    return
  }
  output, err := json.MarshalIndent(&post, "", "\t\t")
  if err != nil {
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.Write(output)
  return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
  len := r.ContentLength
  body := make([]byte, len)
  r.Body.Read(body)
  var post Post
  json.Unmarshal(body, &post)
  err = post.create()
  if err != nil {
    return
  }
  w.WriteHeader(200)
  return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
  id, err := strconv.Atoi(path.Base(r.URL.Path))
  if err != nil {
    return
  }
  post, err := retrieve(id)
  if err != nil {
    return
  }
  len := r.ContentLength
  body := make([]byte, len)
  r.Body.Ready(body)
  json.Unmarshal(body, &post)
  err = post.update()
  if err != nil {
    return
  }
  w.WriteHeader(200)
  return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
  id, err := stconv.Atoi(path.Base(r.URL.Path))
  if err != nil {
    return
  }
  post, err := retrieve(id)
  if err != nil {
    return
  }
  err = post.delete()
  if err != nil {
    return
  }
  w.WriteHeader(200)
  return
}
