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

	file, err := os.Open("version.php")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, " * The WordPress version") {
			Name := "WordPress"
			fmt.Println("This is the Name ", Name)
		} else if strings.HasPrefix(line, "$wp_version") {

			versionext := strings.TrimPrefix(line, "$wp_version = '")
			version := strings.TrimSuffix(versionext, "';")

			fmt.Println("This is the version ", version)

		}

	}
}
