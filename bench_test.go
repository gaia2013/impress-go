package main

import (
  "testing"
)

// デコード関数のベンチマーク
func BenchmarkDecode(b *testing.B) {
  for i	:=  0;	i < b.N;  i++ {
    decode("post.json")
  }
}
