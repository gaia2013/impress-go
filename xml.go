package main

import (
  "encoding/xml"
  "fmt"
  "os"
)

type Post struct {
  XMLName xml.Name  `xml:"post"`
  Id	  string    `xml:"id,attr"`
  Content string    `xml:"content"`
  Author  Author    `xml:"author"`
}

type Author struct {
  Id	string	`xml:"id,attr"`
  Name	string	`xml:",chardata"`
}

func main() {
  post	:=  Post{ //  データの入った構造体を作成する
    Id:	      "1",
    Content:  "Hello World!",
    Author:   Author{
      Id:   "2",
      Name: "Sau Sheong",
    },
  }

  xmlFile, err := os.Create("post.xml")	//  データを保存するXMLファイルを作成する
  if err != nil {
    fmt.Println("Error creating XML file:", err)
    return
  }
  encoder := xml.NewEncoder(xmlFile)  //  XMLファイルに対してエンコーダを生成する
  encoder.Indent("", "\t")
  err = encoder.Encode(&post) //  構造体をファイルにエンコードする
  if err != nil {
    fmt.Println("Error encoding XML to file:", err)
    return
  }
}

