package main

import (
	"fmt"
	"os"
	"sync"
	"io/ioutil"
	"strings"
)

var wg sync.WaitGroup
var search_string string = ""

func walk_r(dir string){
	f, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range list{
		if strings.HasPrefix(v.Name(), ".")==false{ //ignore hidden ones
			path := fmt.Sprintf("%s%s%s", dir, string(os.PathSeparator), v.Name())
			if v.IsDir() {
				wg.Add(1)
				go walk_r(path)
			}else{
				file, err := ioutil.ReadFile(path)
				if err != nil {
					fmt.Println(err)
				}
				str := string(file)
				line := 0
				temp := strings.Split(str, "\n")
				for _, item := range temp {
					if strings.Contains(item, search_string){
						fmt.Println(line, path)
					}
					line++
				}
				
			}
		}
	}
	wg.Done()
}

func main() {
	args := os.Args[1:]
	
	dir := args[0]
	search_string = args[1]
	
	wg.Add(1)
	go walk_r(dir)
	wg.Wait()
}
