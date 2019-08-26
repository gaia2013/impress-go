//  フォーラムの投稿を作成し、読み出し、更新することができる、RESTベースの単純なWebサービス作成

package main

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
)

type Post struct {
  Id  int
  Content string
  Author  string
}

var Db *sql.DB

func init() { //  Dbに接続
  var err error
  Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
  if err != nil {
    panic(err)
  }
}

func retrieve(id int) (post Post, err error) { // 投稿を１つだけ取り出す
  post = Post{}
  err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
  return
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
  _, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
  return
}

func (post *Post) delete() (err error) { // 投稿の削除
  _, err = Db.Exec("delete from posts where id = $1", post.Id)
  return
}

func main() {
  post := Post{Content: "Hello World!", Author: "Sau Sheong"}

  fmt.Println(post)
  post.create()
  fmt.Println(post)

  readPost, _ := retrieve(post.Id)
  fmt.Println(readPost)

  readPost.Content = "Bonjour Monde!"
  readPost.Author = "Pierre"
  readPost.update()


  readPost.delete()
}
