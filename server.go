package main

import(
  "encoding/json"
  "net/http"
  "path"
  "database/sql"
  _ "github.com/lib/pq"
  "strconv"
)

type Post struct {
  Id	  int	  `json:"id"`
  Content string  `json:"content"`
  Author  string  `json:"author"`
}

func retrieve(id int) (post Post, err error) { // 投稿を１つだけ取り出す
  post = Post{}
  err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
  return
}

var Db *sql.DB

 func init() { //  Dbに接続
   var err error
   Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
   if err != nil {
     panic(err)
   }
 }
 
func (post *Post) create() (err error) { // 新しい投稿の作成
  statement := "insert into posts (content, author) values ($1, $2) returning id"
  stmt, err := Db.Prepare(statement)
  if err != nil {
    return
  }
  defer stmt.Close()
  err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
  return
}

func (post *Post) update() (err error) { // 投稿の更新
  _, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content    , post.Author)
  return
}
 
 func (post *Post) delete() (err error) { // 投稿の削除
   _, err = Db.Exec("delete from posts where id = $1", post.Id)
   return
 }




func main() {
  server  :=  http.Server{
    Addr: "127.0.0.1:8080",
  }
  http.HandleFunc("/post/", handleRequest)
  server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r * http.Request) {  // リクエストを正しい関数に振り分けるためのハンドラ関数
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

func handleGet(w http.ResponseWriter, r *http.Request) (err error) { //  投稿の取り出し
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

func handlePost(w http.ResponseWriter, r *http.Request) (err error) { // 投稿の作成
  len := r.ContentLength
  body := make([]byte, len) // バイト列を作成
  r.Body.Read(body)  // バイト列にリクエストの本体を読み込み
  var post Post
  json.Unmarshal(body, &post) // バイト列を構造体Postに組み替え
  err = post.create() // データベースのレコードを作成
  if err != nil {
    return
  }
  w.WriteHeader(200)
  return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) { // 投稿の更新
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
  r.Body.Read(body)
  json.Unmarshal(body, &post)
  err = post.update()
  if err != nil {
    return
  }
  w.WriteHeader(200)
  return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) { // 投稿の削除
  id, err := strconv.Atoi(path.Base(r.URL.Path))
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
