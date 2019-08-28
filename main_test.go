package main  //  テストファイルはテストされる関数と同じパッケージにおく

import (
  "testing"
  "time"
)

func TestDecode(t *testing.T) {
  post, err := decode("post.json")  //	テストされる関数の呼び出し
  if err != nil {
    t.Error(err)
  }
  if post.Id != 1 { //	結果が予想通りかチェックし、違えばエラーメッセージを設定
    t.Error("Wrong id, was expecting 1 but got", post.Id)
  }
  if post.Content != "Hello World!" {
    t.Error("Wrong content, was expecting 'Hello World!' but get", post.Content)
  }
}

func TestEncode(t *testing.T) {
  t.Skip("Skipping encoding for now") //  テストを全て省略
}

func TestLongRunningTest(t *testing.T) {
  if testing.Short() {
    t.Skip("Skipping long-running test in short mode")
  }
  time.Sleep(10 * time.Second)
}
