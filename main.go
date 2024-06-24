package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strings"
	"unicode/utf8"
)

type Flags struct {
	bytes bool
	lines bool
	words bool
	runes bool
}

func (flags *Flags) Empty() bool {
	return !flags.bytes && !flags.lines && !flags.words && !flags.runes 
}

func (flags *Flags) Count() (count int) {
	if flags.bytes {
		count++
	}
	if flags.lines {
		count++
	}
	if flags.words {
		count++
	}
	if flags.runes {
		count++
	}
	return count
}

func GetFileContent(filename string) string {
	dat, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return string(dat)
}

func GetInputContent() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal("Error reading the input: ", err)
	}

	return string(input)
}

func CountBytes(content string) int {
	return len(content)
}

func CountLines(content string) int {
	return strings.Count(content, "\n")
}

func CountWords(content string) int {
	words := strings.Fields(content)

	return len(words)
}

func CountRunes(content string) int {
	return utf8.RuneCountInString(content)
}

func main() {

	var filename string
	var output string
	var content string
	args := os.Args[1:]
	
	flags := Flags{
		bytes: slices.Contains(args, "-c"),
		lines: slices.Contains(args, "-l"),
		words: slices.Contains(args, "-w"),
		runes: slices.Contains(args, "-m"),
	}

	if (len(args) > flags.Count()) {
		filename = args[len(args) - 1]
		content = GetFileContent(filename)
	} else {
		content = GetInputContent()
	}

	if flags.bytes {
		output += fmt.Sprintf("%d ", CountBytes(content))
	}

	if flags.lines {
		output += fmt.Sprintf("%d ", CountLines(content))
	}

	if flags.words {
		output += fmt.Sprintf("%d ", CountWords(content))
	}

	if flags.runes {
		output += fmt.Sprintf("%d ", CountRunes(content))
	}

	if flags.Empty() {
		output += fmt.Sprintf("%d %d %d ", CountLines(content), CountWords(content), CountBytes(content))
	}

	output += filename

	fmt.Println(output)
}
