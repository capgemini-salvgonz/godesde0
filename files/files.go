package files

import (
	"bufio"
	"fmt"
	"os"
)

func WriteStringToFile(content string, path string) {
	file, err := os.Create(path)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Fprintln(file, content)
	file.Close()
}

func AppendStringToFile(content string, path string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Fprintln(file, content)
	file.Close()
}

func ReadFile(path string) {

	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err.Error())
	}

}
