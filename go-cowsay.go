package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	berta := `
        \   ^__^
         \  (oo)\_______
            (__)\       )\/T
                ||----w |
                ||     ||
`
	say := "You suck"
	stat, _ := os.Stdin.Stat()
	if len(os.Args) >= 2 {
		say = os.Args[1]
	} else if (stat.Mode() & os.ModeNamedPipe) != 0 {
		reader, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		say = string(reader)
	}
	say = strings.TrimSuffix(say, "\n")
	lineMax := int(40)
	lines := float64(len(say)) / float64(lineMax)
	if lines < 1 {
		fmt.Print(" ")
		for i := 0; i < len(say)+2; i++ {
			fmt.Print("-")
		}
		fmt.Print("\n< " + say + " >\n")
		fmt.Print(" ")
		for i := 0; i < len(say)+2; i++ {
			fmt.Print("-")
		}
	} else {
		words := strings.Fields(say)
		wordLength := ""
		var s []string
		for _, word := range words {
			if len(wordLength+" "+word) >= lineMax {
				s = append(s, wordLength)
				wordLength = ""
			}
			if wordLength == "" {
				wordLength = wordLength + word
			} else {
				wordLength = wordLength + " " + word
			}
		}
		s = append(s, wordLength) // Append last sentence to slice
		longestString := 0
		for _, w := range s {
			if len(w) > longestString {
				longestString = len(w)
			}
		}
		fmt.Print(" ")
		for i := 0; i < longestString+2; i++ {
			fmt.Print("-")
		}
		fmt.Print("\n")
		lineCount := 0
		for _, line := range s {
			if lineCount == 0 {
				fmt.Print("/ " + line)
				for i := 0; i < longestString-len(line); i++ {
					fmt.Print(" ")
				}
				fmt.Print(" \\\n")
			} else if lineCount == len(s)-1 {
				fmt.Print("\\ " + line)
				for i := 0; i < longestString-len(line); i++ {
					fmt.Print(" ")
				}
				fmt.Print(" /\n")
			} else {
				fmt.Print("| " + line)
				for i := 0; i < longestString-len(line); i++ {
					fmt.Print(" ")
				}
				fmt.Print(" |\n")
			}
			lineCount++
		}
		fmt.Print(" ")
		for i := 0; i < longestString+2; i++ {
			fmt.Print("-")

		}
	}
	fmt.Print(berta)
}
