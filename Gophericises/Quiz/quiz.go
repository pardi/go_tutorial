package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func compose_question_msg(num_question int, str []string) (msg string) {

	msg = "Question " + fmt.Sprint(num_question) + ": " + str[0]

	return
}

func compare_answer(ans_correct int, ans_given int) (res bool) {

	if ans_correct == ans_given {
		res = true
	} else {
		res = false
	}

	return
}

func main() {

	data, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	csv_reader := csv.NewReader(data)

	var num_answered int = 0
	var num_questions int = 0

	for {
		entry, err := csv_reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(compose_question_msg(num_questions+1, entry))

		var answer int
		fmt.Scan(&answer)

		correct_ans, _ := strconv.Atoi(entry[1])

		if compare_answer(correct_ans, answer) {
			num_answered += 1
		}
		num_questions += 1

	}

	// Print result

	fmt.Println(" ")
	fmt.Println("Score: ", num_answered, " out of ", num_questions)

}
