package main

import (
	"fmt"
	_"flag"
	"os"
	"io/ioutil"
	"strings"
	"path"
	"log"
	"baliance.com/gooxml/document"
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"time"
	
)
func main(){
	var basenames string
	var str string
	//日志记录
	fileName := "Info_First.log"
	logFile,err:= os.Create(fileName)

	defer logFile.Close()

    if err != nil {
        log.Fatalln("open file error")
	}
	
	files, dirs, _ := GetFilesAndDirs("./simplemath")//获取文件

	for _, table := range dirs {
        temp, _, _ := GetFilesAndDirs(table)
        for _, temp1 := range temp {
            files = append(files, temp1)
        }
    }

    for _, table1 := range files {
        fmt.Printf("获取的文件为[%s]\n", table1)
    }



	xfiles, _ := GetAllFiles("/home/shaonianlang/下载")

    for _, file := range xfiles {
		
		filenameWithSuffix := path.Base(file)
		fileSuffix := path.Ext(file)
		basenames = strings.TrimSuffix(filenameWithSuffix, fileSuffix)

		log.Println("开始转换文件%s\n",filenameWithSuffix)
		

			//读取doc文档
		doc, err := document.Open(file)
		if err != nil {
			log.Fatalf("文件出错 %s", err)
		}
		
		//生成doc文档
		
		//doc.Paragraphs()得到包含文档所有的段落的切片
		for i, para := range doc.Paragraphs() {
			//run为每个段落相同格式的文字组成的片段
			_ =i
			for j, run := range para.Runs() {
				_ =j
				str +=run.Text()

			}

		}
    }

	

	file6, error := os.Create("/home/shaonianlang/下载/"+basenames+".txt")
	if error != nil {
        fmt.Println(error)
    }
	file6.WriteString(str)
	file6.Close()
	str =""
	
	log.Println("文件生成ok")
	StartCac()

	
}


func GetAllFiles(dirPth string) (files []string, err error) {
    var dirs []string
    dir, err := ioutil.ReadDir(dirPth)
    if err != nil {
        return nil, err
    }

    PthSep := string(os.PathSeparator)
    //suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

    for _, fi := range dir {
        if fi.IsDir() { // 目录, 递归遍历
            dirs = append(dirs, dirPth+PthSep+fi.Name())
            GetAllFiles(dirPth + PthSep + fi.Name())
        } else {
            // 过滤指定格式
            ok := strings.HasSuffix(fi.Name(), ".doc")
            if ok {
                files = append(files, dirPth+PthSep+fi.Name())
            }
        }
    }

    // 读取子目录下文件
    for _, table := range dirs {
        temp, _ := GetAllFiles(table)
        for _, temp1 := range temp {
            files = append(files, temp1)
        }
    }

    return files, nil
}


func StartCac() {
    t1 := time.Now() // get current time
    //logic handlers
    for i := 0; i < 1000; i++ {
        fmt.Print("*")
    }
    elapsed := time.Since(t1)
    log.Println("时间花费位:\n" , elapsed)
}


//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
    dir, err := ioutil.ReadDir(dirPth)
    if err != nil {
        return nil, nil, err
    }

    PthSep := string(os.PathSeparator)
    //suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

    for _, fi := range dir {
        if fi.IsDir() { // 目录, 递归遍历
            dirs = append(dirs, dirPth+PthSep+fi.Name())
            GetFilesAndDirs(dirPth + PthSep + fi.Name())
        } else {
            // 过滤指定格式
            ok := strings.HasSuffix(fi.Name(), ".go")
            if ok {
                files = append(files, dirPth+PthSep+fi.Name())
            }
        }
    }

    return files, dirs, nil
}








