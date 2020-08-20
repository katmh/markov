package main

import (
	"fmt"
	"flag"
	"strings"
	"os"
	"bufio"
	"io/ioutil"
	"encoding/json"
	"github.com/mb-14/gomarkov"
)

func main() {
	train := flag.Bool("train", false,  "Train the Markov chain")
	flag.Parse()
	if *train {
		chain := buildModel()
		saveModel(chain)
	} else {
		chain, err := loadModel()
		if err != nil {
			fmt.Println(err)
			return
		}
		generateText(chain)
	}
}

func buildModel() (*gomarkov.Chain) {
	chain := gomarkov.NewChain(1)
	file, _ := os.Open("data.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		chain.Add(split)
	}
	return chain
}

func saveModel(chain *gomarkov.Chain) {
	jsonObj, _ := json.Marshal(chain)
	err := ioutil.WriteFile("model.json", jsonObj, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func loadModel() (*gomarkov.Chain, error) {
	var chain gomarkov.Chain
	data, err := ioutil.ReadFile("model.json")
	if err != nil {
		return &chain, err
	}
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return &chain, err
	}
	return &chain, nil
}

func generateText(chain *gomarkov.Chain) {
	tokens := []string{gomarkov.StartToken}
	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}
	fmt.Println(strings.Join(tokens[1:len(tokens)-1], " "))
}