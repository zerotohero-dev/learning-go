package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var britishAmerican = "british-american.txt"

func init() {
	dir, _ := filepath.Split(os.Args[0])
	britishAmerican = filepath.Join(dir, britishAmerican)
}

func fileNamesFromCommandLine() (inFileName, outFileName string, err error) {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err = fmt.Errorf("usage: %s [<]infile.txt [>]outfile.txt",
			filepath.Base(os.Args[0]))

		return "", "", err
	}

	if len(os.Args) > 1 {
		inFileName = os.Args[1]
		if len(os.Args) > 2 {
			outFileName = os.Args[2]
		}
	}

	if inFileName != "" && inFileName == outFileName {
		log.Fatal("won’t overwrite infile")
	}

	return inFileName, outFileName, nil
}

func makeReplacer(file string) (func(string) string, error) {
	rawBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	text := string(rawBytes)

	usForBritish := make(map[string]string)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) == 2 {
			usForBritish[fields[0]] = fields[1]
		}
	}

	return func(word string) string {
		if usWord, found := usForBritish[word]; found {
			return usWord
		}

		return word
	}, nil
}

func americanize(inFile io.Reader, outFile io.Writer) (err error) {
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	var replacer func(string) string

	if replacer, err = makeReplacer(britishAmerican); err != nil {
		return err
	}

	wordRx := regexp.MustCompile("[A-Za-z]+")
	eof := false

	for !eof {
		var line string
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else if err != nil {
			return err
		}

		line = wordRx.ReplaceAllStringFunc(line, replacer)
		if _, err = writer.WriteString(line); err != nil {
			return nil
		}
	}

	return nil
}

func main() {
	inFileName, outFileName, err := fileNamesFromCommandLine()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inFile, outFile := os.Stdin, os.Stdout

	if inFileName!= "" {
		if inFile, err = os.Open(inFileName); err != nil {
			log.Fatal(err)
		}
		defer func() {
			_ = inFile.Close()
		}()
	}
	if outFileName!= "" {
		if outFile, err = os.Open(outFileName); err != nil {
			log.Fatal(err)
		}
		defer func() {
			_ = outFile.Close()
		}()
	}
	if err = americanize(inFile, outFile); err != nil {
		log.Fatal(err)
	}
}
