package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func handleCommandLine() (algorithm int, minSize, maxSize int64,
	suffixes, files []string) {
	flag.IntVar(&algorithm, "algorithm", 1, "1 or 2")
	flag.Int64Var(&minSize, "min", -1,
		"minimum file size (-1 means no minimum)")
	flag.Int64Var(&maxSize, "max", -1,
		"maximum file size (-1 means no maximum)")
	var suffixesOpt = flag.String("suffixes", "",
		"comma-separated list of file suffixes")
	flag.Parse()
	if algorithm != 1 && algorithm != 2 {
		algorithm = 1
	}
	if minSize > maxSize && maxSize != -1 {
		log.Fatalln("minimum size must be < maximum size")
	}
	suffixes = []string{}
	if *suffixesOpt != "" {
		suffixes = strings.Split(*suffixesOpt, ",")
	}
	files = flag.Args()
	return algorithm, minSize, maxSize, suffixes, files
}

func sink(in <-chan string) {
	for filename := range in {
		fmt.Println(filename)
	}
}

func filterSize(min, max int64, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if min == -1 || max == -1 {
				out <- filename
				continue
			}

			finfo, err := os.Stat(filename)
			if err != nil {
				continue
			}

			size := finfo.Size()
			if (min == -1 || min > -1 && min <= size) &&
				(max == -1 || max > -1 && max >= size) {
				out <- filename
			}
		}
		close(out)
	}()
	return out
}

func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if len(suffixes) == 0 {
				out <- filename
				continue
			}

			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <-filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}

func source(files []string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()
	return out
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(0)

	algorithm, minSize, maxSize, suffixes, files := handleCommandLine()

	if algorithm == 1 {
		sink(
			filterSize(
				minSize,
				maxSize,
				filterSuffixes(suffixes, source(files)),
			),
		)

		return
	}

	channel1 := source(files)
	channel2 := filterSuffixes(suffixes, channel1)
	channel3 := filterSize(minSize, maxSize, channel2)
	sink(channel3)
}
