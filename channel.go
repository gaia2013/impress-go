
// チャネルを閉じる
package main

import (
  "fmt"
)

func callerA(c chan string) {
  c <- "Hello World!"
  close(c) // 1.関数が呼び出されたらチャネルを閉じる
}

func callerB(c chan string) {
  c <- "Hola Mundo!"
  close(c) // 1.関数が呼び出されたらチャネルを閉じる
}

func main() {
  a, b := make(chan string), make(chan string)
  go callerA(a)
  go callerB(b)

  var msg string
  ok1, ok2 := true, true
  for ok1 || ok2 {
    select {
    case msg, ok1 = <-a: // 2.チャネルが閉じているとok1とok2はfalseになる
      if ok1 {
	fmt.Printf("%s from A\n", msg)
      }
    case msg, ok2 = <-b: // 2.チャネルを閉じているとok1とok2はfalseになる
      if ok2 {
	fmt.Printf("%s from B\n", msg)
      }
    }
  }
}
