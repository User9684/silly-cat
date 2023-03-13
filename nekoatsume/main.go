package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type ImgData struct {
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

func format(toFormat, split string) string {
	var formattedStrings []string
	for _, str := range strings.Split(toFormat, split) {
		firstCharacter := str[:1]
		everythingElse := str[1:]
		formattedStrings = append(formattedStrings, fmt.Sprintf("%s%s", strings.ToUpper(firstCharacter), everythingElse))
	}

	return strings.Join(formattedStrings, split)
}

func main() {
	items, err := os.ReadDir("./")
	if err != nil {
		fmt.Println(err)
	}
	for _, item := range items {
		nameSplit := strings.Split(item.Name(), ".")
		if nameSplit[1] != "png" {
			continue
		}

		catName := nameSplit[0]
		catNameFormatted := strings.TrimSpace(format(format(catName, " "), "-"))

		file, err := os.Create(fmt.Sprintf("%s.json", catName))
		if err != nil {
			fmt.Println(err)
			continue
		}

		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.SetIndent("", "    ")
		encoder.Encode(&ImgData{
			Name:        catNameFormatted,
			Author:      "Jenku",
			Description: fmt.Sprintf("%s from Neko Atsume, redrawn by Jenku!", catNameFormatted),
		})
	}
}
