package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parse(str string) (res []string) {

	res = strings.Split(str, ",")

	return
}

func main() {

	data, err := os.Open("problems.csv")

	if err != nil {
		fmt.Printf("ERROR! The file does not exist")
	}

	defer data.Close()

	scanner := bufio.NewScanner(data)
	var answered int = 0
	var questions int = 0

	var msg string

	for scanner.Scan() {
		res := parse(scanner.Text())
		msg = "Question " + fmt.Sprint(questions+1) + ": " + res[0]
		fmt.Println(msg)
		questions += 1

		var answer int
		fmt.Scan(&answer)

		intres, _ := strconv.Atoi(res[1])

		if answer == intres {
			answered += 1
		}

	}

	fmt.Println(" ")
	fmt.Println("Score: ", answered, " out of ", questions)

}
