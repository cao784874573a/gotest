package main

import (
	"fmt"
	"flag"
	"os"
	_ "io/ioutil"
	_ "strings"
	_ "path"
	"log"
	"baliance.com/gooxml/document"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "time"
	
)
var sentence = flag.String("sentence", "我来到北京清华大学", "")

const TAG = "main: "

func main(){
	flag.Parse()
	
	var str string
	doc, err := document.Open(*sentence)
	Mylog(doc)
		if err != nil {
			Mylog(err)
		}
		
		//生成doc文档
		
		//doc.Paragraphs()得到包含文档所有的段落的切片
		for i, para := range doc.Paragraphs() {
			//run为每个段落相同格式的文字组成的片段
			_ =i
			for j, run := range para.Runs() {
				_ =j
				str+=run.Text()
				
			}

		}

		fmt.Println(str)
		

	
	

}

func Mylog(v ...interface{}) {
    f, err := os.OpenFile("20181105go.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		Mylog(err)
	}
    defer f.Close()
    logger := log.New(f, TAG, log.Ldate|log.Ltime|log.Lmicroseconds)
    logger.Println(v...)
}


















