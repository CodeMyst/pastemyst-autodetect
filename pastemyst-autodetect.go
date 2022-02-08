package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/go-enry/go-enry/v2"
	"github.com/go-enry/go-enry/v2/data"
)

func main() {
	if len(os.Args) != 2 {
		panic("missing file path argument")
	}

	dat, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}

	langs := make([]string, 0, len(data.LanguagesLogProbabilities))
	for k := range data.LanguagesLogProbabilities {
		langs = append(langs, k)
	}

	lang, _ := enry.GetLanguageByClassifier(dat, langs[:])

	if lang == "" {
		fmt.Println("unknown")
	} else {
		fmt.Println(lang)
	}
}
