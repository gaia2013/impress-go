package main

import (
  "encoding/xml"
  "fmt"
  "io"
  "os"
)

type Post struct {    //#A
  XMLName  xml.Name  `xml:"post"`
  Id	   string    `xml:"id,attr"`
  Content  string    `xml:"content"`
  Author   Author    `xml:"author"`
  Xml	   string    `xml:",innerxml"`
  Comments []Comment `xml:"comments>comment"`
}

type Author struct {
  Id	string  `xml:"id,attr"`
  Name	string	`xml:",chardata"`
}

type Comment struct {
  Id	  string  `xml:"id, attr"`
  Content string  `xml:"content"`
  Author  Author  `xml:"author"`
}

func main() {
  xmlFile, err := os.Open("post.xml")
  if err != nil {
    fmt.Println("Error opening XML file:", err)
    return
  }
  defer xmlFile.Close()

  decoder := xml.NewDecoder(xmlFile)  // XMLデータからデコーダ(decoder）を生成
  for { // decoder内のXMLデータを順次処理
    t, err := decoder.Token()  // 各処理でdecoderからトークン（Token）を取得
    if err == io.EOF {
      break
    }
    if err != nil {
      fmt.Println("Error decoding XML into tokens:", err)
      return
    }

    switch se := t.(type) { // トークンの型をチェック
    case xml.StartElement:

      if se.Name.Local == "comment" {
	var comment Comment
	decoder.DecodeElement(&comment, &se) // XMLデータを構造体にデコード
      }
    }
  }
}
