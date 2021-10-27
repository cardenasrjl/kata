package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

type Kata struct {
	Path string
	Date string
}

func main() {
	
	fmt.Printf("The files are %+v",GetFiles("./") )
}


func GetFiles(root string) []Kata {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	

	localDirs := []Kata{}
	for _, f := range files {
		if f.IsDir() {
			localDirs = append(localDirs,GetFiles(root + "/" + f.Name())...)
		} else {
			//get the date of the last modification date
			f.ModTime()
		}
	}
	
	if len(localDirs) == 0 {
		localDirs = append( localDirs,Kata{
			Path: root,
		})
	}
	
	return localDirs
	
}




