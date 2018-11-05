package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/yanyiwu/gosimhash"
	"github.com/yanyiwu/gojieba"
	

)

var sentence = flag.String("sentence", "我来到北京清华大学", "")
var top_n = flag.Int("top_n", 4000, "")

func main() {
	flag.Parse()
 	// simhah
	hasher := gosimhash.New("../dict/jieba.dict.utf8", "../dict/hmm_model.utf8", "../dict/idf.utf8", "../dict/stop_words.utf8")
	
	defer hasher.Free()
	
	fingerprint := hasher.MakeSimhash(*sentence, *top_n)




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

	//fmt.Printf("%s\n",string(bs))
	
	fmt.Println(string(bs))
	//fmt.Printf("%s\n",*sentence)

}
