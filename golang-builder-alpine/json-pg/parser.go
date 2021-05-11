package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ProcessFile(file string) {
	info, err := os.Stat(file)

	if err != nil {
		log.Printf("Error: %v", err)
	}
	if os.IsNotExist(err) {
		log.Printf("Error: %v", err)
	}

	filename := info.Name()
	//fmt.Println("Name:" + filename)
	//fmt.Printf("Size: %d \n", info.Size())
	dir := filepath.Dir(file)

	//fmt.Printf("Dir: %s \n", dir)
	abs, err := filepath.Abs(dir)
	check(err)
	//fmt.Printf("Abs file path: %s \n", abs)

	//var contents string
	full := filepath.Join(abs, filename)
	//fmt.Printf("Full file name: %s \n", full)
	dat, err := ioutil.ReadFile(full)
	check(err)
	//fmt.Print(string(dat))
	contents := string(dat)

	newContents := strings.Replace(contents, "json:\"-\"", "json:\"-\" pg:\"-\"", -1)

	//newContents = strings.Replace(newContents, ",omitempty", "", -1)

	err = ioutil.WriteFile(file, []byte(newContents), 0)
	check(err)

}
