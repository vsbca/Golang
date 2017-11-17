package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	//compile a regular expression to match

	//Open a file in container it needs to be located in layer

	file, err := os.Open("joomla.xml")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasSuffix(line, "joomla.org</authorUrl>") {
			Name := "Joomla"
			fmt.Println("This is the Name ", Name)
		} else if strings.HasSuffix(line, "</version>") {

			versionext := strings.TrimPrefix(line, "	<version>")
			version := strings.TrimSuffix(versionext, "</version>")

			fmt.Println("This is the version ", version)

		}

	}
}
