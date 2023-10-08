package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
)

func byteCount(r io.Reader, wg *sync.WaitGroup) {
	defer wg.Done()

	var count int

	fileScanner := bufio.NewReader(r)

	for {
		_, err := fileScanner.ReadByte()
		if err != nil {
			break
		} else {
			count++
		}
	}

	fmt.Print(count, " ")
}

func lineCount(r io.Reader, wg *sync.WaitGroup) {
	defer wg.Done()

	var count int

	fileScanner := bufio.NewScanner(r)

	for fileScanner.Scan() {
		count++
	}

	fmt.Print(count, " ")
}

func wordCount(r io.Reader, wg *sync.WaitGroup) {
	defer wg.Done()

	var count int

	fileScanner := bufio.NewScanner(r)
	fileScanner.Split(bufio.ScanWords)

	for fileScanner.Scan() {
		count++
	}

	fmt.Print(count, " ")
}

func charCount(r io.Reader, wg *sync.WaitGroup) {
	defer wg.Done()

	var count int

	fileScanner := bufio.NewScanner(r)

	for fileScanner.Scan() {
		count += len(fileScanner.Text())
	}

	fmt.Print(count, " ")
}

func main() {
	var byteFlag, lineFlag, wordFlag, charFlag bool

	flag.BoolVar(&byteFlag, "c", false, "print the byte counts")
	flag.BoolVar(&lineFlag, "l", false, "print the newline counts")
	flag.BoolVar(&wordFlag, "w", false, "print the word counts")
	flag.BoolVar(&charFlag, "m", false, "print the character counts")

	flag.Parse()

	filePath := flag.Arg(0)

	var fileReader io.Reader = os.Stdin

	if len(filePath) > 0 {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("wc: %s: No such file or directory\n", filePath)
			return
		}
		defer file.Close()

		fileReader = file
	}

	wg := new(sync.WaitGroup)

	byteReader, byteWriter := io.Pipe()
	lineReader, lineWriter := io.Pipe()
	wordReader, wordWriter := io.Pipe()
	charReader, charWriter := io.Pipe()

	var writers []io.Writer

	if byteFlag {
		wg.Add(1)
		writers = append(writers, byteWriter)
		go byteCount(byteReader, wg)
	}

	if lineFlag {
		wg.Add(1)
		writers = append(writers, lineWriter)
		go lineCount(lineReader, wg)
	}

	if wordFlag {
		wg.Add(1)
		writers = append(writers, wordWriter)
		go wordCount(wordReader, wg)
	}

	if charFlag {
		wg.Add(1)
		writers = append(writers, charWriter)
		go charCount(charReader, wg)
	}

	io.Copy(io.MultiWriter(writers...), fileReader)

	byteWriter.Close()
	lineWriter.Close()
	wordWriter.Close()
	charWriter.Close()

	wg.Wait()

	fmt.Print(filePath)

}
