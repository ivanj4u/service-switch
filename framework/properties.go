package framework

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	applicationProperties map[string] string
	databaseProperties    map[string] string
)

func initProperties() {
	applicationProperties = getProperties("/cfg/application.properties")
	databaseProperties = getProperties("/cfg/database.properties")
}

func getProperties(filename string) (map[string] string) {
	properties := map[string] string{}
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + filename)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				properties[key] = value
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	file.Close()
	return properties
}
