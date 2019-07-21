package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
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
	data, err := ioutil.ReadFile("./article.txt")
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

	pl := make(PairList, len(wordcount))
	i := 0
	for k, v := range wordcount {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	for _, p := range pl {
		fmt.Printf("%s,%d\n", p.Key, p.Value)
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
