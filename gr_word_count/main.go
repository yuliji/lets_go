package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

func IsDigit(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

func main() {
	files, err := ioutil.ReadDir("./articles")
	if err != nil {
		log.Fatal(err)
	}

	words := make(map[string]int)

	numCpu := runtime.NumCPU()

	outc := make(chan map[string]int, len(files))

	inqueues := make([]chan string, numCpu)

	for i := 0; i < numCpu; i++ {
		inqueue := make(chan string, len(files))
		inqueues[i] = inqueue
		go countWord(inqueue, outc)
	}

	for idx, file := range files {
		filePath := path.Join("./articles", file.Name())
		inqueues[idx%numCpu] <- filePath
	}

	for i := 0; i < numCpu; i++ {
		close(inqueues[i])
	}

	for i := 0; i < len(files); i++ {
		articleWords := <-outc
		for w, c := range articleWords {
			_, ok := words[w]
			if ok {
				words[w] += c
			} else {
				words[w] = c
			}
		}
	}

	pl := make(PairList, len(words))
	i := 0
	for k, v := range words {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	for _, p := range pl {
		fmt.Printf("%s,%d\n", p.Key, p.Value)
	}
}

func countWord(inc chan string, outc chan map[string]int) {
	// fmt.Println("start countWord")

	for filePath := range inc {
		// fmt.Printf("work on %s\n", filePath)
		data, err := ioutil.ReadFile(filePath)
		check(err)
		filecontent := string(data)
		lines := strings.Split(filecontent, "\n")

		r := regexp.MustCompile(`\[\d+\]`)
		punctuations := regexp.MustCompile(`[\(\)';\":/.,<>\?&\[\]\{\}%\-—#\$\+!–––]`)

		var wordcount map[string]int
		wordcount = make(map[string]int)

		for _, line := range lines {
			line = strings.Trim(line, " \n\t\r")
			line = r.ReplaceAllString(line, " ")
			line = punctuations.ReplaceAllString(line, " ")
			if line == "" {
				continue
			}
			line = strings.ToLower(line)
			words := strings.Split(line, " ")
			for _, word := range words {
				word = strings.TrimSpace(word)
				if word == "" {
					continue
				}
				if !IsLetter(word) || IsDigit(word) {
					continue
				}
				cnt, ok := wordcount[word]
				if ok {
					wordcount[word] = cnt + 1
				} else {
					wordcount[word] = 1
				}
			}
		}
		outc <- wordcount
	}
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
