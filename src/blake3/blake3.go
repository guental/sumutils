package main

import (
	"github.com/zeebo/blake3"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

const chunkSize = 1024*1024

func usage() {
    fmt.Printf("Usage: b3sum [FILE]...\n")
}

func isFile(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if fi.Mode().IsDir() {
		return false
	}
	return true
}


func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}
	if fi.Mode().IsDir() {
		return true
	}
	return false
}

func Blake3SumFile(filename string) string {
	inFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()
	buf := make([]byte, chunkSize)
	hash := blake3.New()
	for {
		nbRead, err := inFile.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if nbRead == 0 {
			break
		}
		hash.Write(buf[:nbRead])
	}
	return hex.EncodeToString(hash.Sum(nil)[:])
}

func main() {
    flag.Usage = usage
	flag.Parse()

	exitCode := 0

	if len(flag.Args()) < 1 {
		os.Exit(0)
	}

	for _, inFilename := range os.Args[1:] {
		if isFile(inFilename) {
			hash := Blake3SumFile(inFilename)
			if len(flag.Args()) < 2 {
				fmt.Printf("%s\n", hash)
				os.Exit(0)
			} else {
				fmt.Printf("%s\t%s\n", hash, inFilename)
			}
			continue
		} else {
			if isDir(inFilename) {
				fmt.Printf("b3sum: %s: Is a directory\n", inFilename)
				exitCode = 1
				continue
			} else {
				fmt.Printf("b3sum: %s: No such file or directory\n", inFilename)
				exitCode = 1
				continue
			}
		}
	}
os.Exit(exitCode)
}
