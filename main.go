/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"chia-replotter/utils"
	"flag"
	"fmt"
	"log"
	"time"
)

var oldDir = flag.String("old", "", "Old dir would be cleaned.")
var newDir = flag.String("new", "", "Directory will store new files")
var prefix = flag.String("prefix", "plot", "prefix of files will be cleaned")
var suffix = flag.String("suffix", ".plot", "suffix of files will be cleaned")

func main() {
	flag.Parse()
	for i := 0; i != flag.NArg(); i++ {
		fmt.Printf("arg[%d]=%s\n", i, flag.Arg(i))
	}

	if *oldDir == "" || *newDir == "" {
		log.Fatalln("old/new dir is missing. ")
		return
	}
	for true {
		log.Println("Checking files change...")
		utils.RemoveOldFiles(*oldDir, *newDir, *prefix, *suffix)
		time.Sleep(time.Second * 5)
	}
}
