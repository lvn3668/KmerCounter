package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {

	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)

	if (len(os.Args) <2) {
		log.Fatalf("This code takes 2 arguments: path to directory and kmer counter")
	}
	var kmerlength int
	kmerlength, err := strconv.Atoi(os.Args[2])

	if kmerlength < 0 {
		log.Fatalf("Kmer length cannot be less than zero")
	} else {
		fmt.Println("Desired kmer length is %s", kmerlength)
	}

	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	var path string
	path, err = os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {

		if strings.HasSuffix(f.Name(), "fasta") {
			file, err := os.Open(f.Name())
			if err != nil {
				log.Fatal(err)
			} else
			{
				fmt.Println(filepath.Join(path, f.Name()))
			}
			defer file.Close()
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var (
				text      []string
				totaltext string
			)

			var linenumber int
			var linereadin string
			linenumber = 1
			for scanner.Scan() {
				linereadin = strings.ToLower(scanner.Text())
				if linenumber > 1 {
				text = append(text, linereadin)
				totaltext = totaltext + strings.ToLower(scanner.Text())
				}
				linenumber++
			}

			// The method os.File.Close() is called
			// on the os.File object to close the file
			file.Close()

			totaltext = strings.ReplaceAll(totaltext, ">", "")
			fmt.Println(totaltext)
			// check if the fasta file contains only A T G C
			isDNA := regexp.MustCompile(`^[atgc]`).MatchString

			if !isDNA(totaltext) {
				fmt.Printf("%q is not valid\n", f.Name())
				log.Fatalf("Not valid DNA file")
			} else {
				fmt.Println("Valid Fasta File")
			}

			var (
				kmerlengths = make(map[string]int, 1000)
			)

			fmt.Printf("Length of sequence in file %s is %d\n", f.Name(), len(text))
			// if length is greater than zero, then count the kimer length
			// check if kmer length is greater than length of file
			if kmerlength > len(text) {
				log.Fatalf("Kmer length cannot be greater than file length")
			}
			if len(text) > 0 {
				for i := 0; i < len(totaltext)-kmerlength; i = i + kmerlength {
					var _, checkkey = kmerlengths[totaltext[i:i+kmerlength]]
					if checkkey {
						kmerlengths[totaltext[i:i+kmerlength]]++
					} else {
						kmerlengths[totaltext[i:i+kmerlength]] = 1
					}
				}
			} else
			{
				log.Fatalf("Empty fasta file")
			}
			type kv struct {
				Key   string
				Value int
			}
			var ss []kv
			for k, v := range kmerlengths {
				ss = append(ss, kv{k, v})
			}

			sort.Slice(ss, func(i, j int) bool {
				return ss[i].Value > ss[j].Value
			})


			//Write it to file
			opfile, err := os.Create(f.Name()+"."+os.Args[2]+".counts.txt")

			if (err != nil) {
				log.Fatalf("Unable to open output file %s", f.Name())
			} else {
				fmt.Printf("Writing to file %s", opfile.Name())
			}

			defer opfile.Close()

			for _, kv := range ss {
				fmt.Println("Key:", kv.Key, "Value:", kv.Value)
				_, err2 := opfile.WriteString(
					"Kmer: "+
					kv.Key+
					"Count: "+
					string(kv.Value)+
				 "\n")

				if err2 != nil {
					log.Fatal(err2)
				}

			}
			fmt.Printf("Finished writing to file %s ", opfile.Name())
		}

	}

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	//fmt.Println("Contents of file:", string(data))
}
