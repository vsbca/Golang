package main

import (
	"fmt"
	"strings"
)

func main() {

	data := map[string]int{
		"usr/share/wp/version.php": 123,
		"testin.txt":               2,
		"var/www/nginx/htdocs/wp/wp-includes/version.php": 234,
	}

	for key := range data {
		if strings.HasSuffix(key, "wp-includes/version.php") {
			fmt.Println("This is the key value", key)
		}
	}
}
