package main

import (
	"fmt"
	"os"

	"enghimanshu.tech/go/files/fileutils"
)

func main(){

	rootPath,_:=os.Getwd()
	filepath := rootPath + "/data/text.txt"
    content,error := fileutils.ReadTextFile(filepath)
    if error == nil {
	     	fmt.Println(content)
			newContent := fmt.Sprintf("Original: %v\n Double: %v%v", content, content, content)
	        fileutils.WriteTextFile(filepath, newContent)
			}else{
		fmt.Println(error)
	}

}