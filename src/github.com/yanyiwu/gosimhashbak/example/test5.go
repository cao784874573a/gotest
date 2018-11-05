package main

import (
	"fmt"
	"io/ioutil"
	_ "strings"
	
	"github.com/yanyiwu/gojieba"
	
)
func main(){

	b, err := ioutil.ReadFile("/home/python/webroot/zuo2.txt")

    if err != nil {
        fmt.Print(err)
    }
    
   
    str := string(b)


	extt := gojieba.NewJieba()
	
	defer extt.Free()
	
	word_weights := extt.ExtractWithWeight(str, 1600)
	
	fmt.Println("关键词抽取:", word_weights)
}





