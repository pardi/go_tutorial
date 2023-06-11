package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
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
func parseflags() (csv_filename string, timer time.Duration) {

	const (
		default_CSV   = "problems.csv"
		usage_csv     = "The file containing the questions/answers to the quiz. They must be numeric answers!"
		default_timer = 30 * time.Second
		usage_timer   = "The maximum timer for answering all questions"
	)

	flag.StringVar(&csv_filename, "file", default_CSV, usage_csv)
	flag.StringVar(&csv_filename, "f", default_CSV, usage_csv+" (shorthand)")
	flag.DurationVar(&timer, "timer", default_timer, usage_timer)
	flag.DurationVar(&timer, "t", default_timer, usage_timer+" (shorthand)")

	flag.Parse()

	return
}

func main() {

	// Get inputs from the command line
	csv_filename, timer := parseflags()

	// Open file
	data, err := os.Open(csv_filename)

	if err != nil {
		log.Fatal(err)
	}
	defer data.Close()

	csv_reader := csv.NewReader(data)

	var num_answered int = 0
	var num_questions int = 0

	_ = time.AfterFunc(timer, func() {
		fmt.Println("Sorry, the time is up!")
		os.Exit(1)
	})

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
