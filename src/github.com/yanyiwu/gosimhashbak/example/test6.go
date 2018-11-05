package main

import (
	"fmt"
	_ "flag"
	_ "io/ioutil"
	_ "strings"
	"log"
	"github.com/yanyiwu/gosimhash"
	"baliance.com/gooxml/document"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	
)
func main(){

	//读取doc文档
	doc, err := document.Open("/home/wwwroot/niandu/new/admin/hello.docx")
    if err != nil {
        log.Fatalf("error opening document: %s", err)
	}
	
	//生成doc文档
	var str string
    //doc.Paragraphs()得到包含文档所有的段落的切片
    for i, para := range doc.Paragraphs() {
        //run为每个段落相同格式的文字组成的片段
        _ =i
        for j, run := range para.Runs() {
            _ =j
            str +=run.Text()+"\n"

		}

	}

	hasher := gosimhash.New("../dict/jieba.dict.utf8", "../dict/hmm_model.utf8", "../dict/idf.utf8", "../dict/stop_words.utf8")
	
	defer hasher.Free()
	
	fingerprint := hasher.MakeSimhash(str, 2000)




	var s string = "0000000000000000000000000000000000000000000000000000000000000000"
	bs := []byte(s)
						
    for i := 63; i >= 0; i-- {
		
		if (fingerprint&1)==1 {

			bs[i]='1'
		} else {

			bs[i]='0'
		}
		fingerprint >>=1
	}

	
	fmt.Println(str)
	fmt.Println(string(bs))
	
	


	
}





