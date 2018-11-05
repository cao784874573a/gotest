package main

import (
	"fmt"
	"flag"
	_ "strings"
	

	"github.com/yanyiwu/gosimhash"
)


var sentence = flag.String("sentence", "我来到北京清华大学", "")
var top_n = flag.Int("top_n", 4, "")

func main(){

 flag.Parse()
	// simhah
   hasher := gosimhash.New("../dict/jieba.dict.utf8", "../dict/hmm_model.utf8", "../dict/idf.utf8", "../dict/stop_words.utf8")
  
   
   fingerprint := hasher.MakeSimhash(*sentence, *top_n)

  

	

  
	fmt.Println(fingerprint)

}



