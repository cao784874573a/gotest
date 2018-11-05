package main

import (
	
	"fmt"
	"io/ioutil"
	"github.com/yanyiwu/gosimhash"
)



func main() {


	b, err := ioutil.ReadFile("/home/python/webroot/test.txt")
    if err != nil {
        fmt.Print(err)
    }
    
   
    str := string(b)

    hasher := gosimhash.New("../dict/jieba.dict.utf8", "../dict/hmm_model.utf8", "../dict/idf.utf8", "../dict/stop_words.utf8")
	defer hasher.Free()
	fingerprint := hasher.MakeSimhash(str, 5)
   	
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

	fmt.Println(string(bs))


}
