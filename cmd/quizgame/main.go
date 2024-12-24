package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Question struct {
	question string
	answer   string
}

func findAbsolutePath(relativePath string) (string, error) {
	absolutePath, err := filepath.Abs(relativePath)
	if err != nil {
		println("Error while converting relative path to absolute path", err)
		return "", err
	}
	return absolutePath, nil
}

func readCSV(csvPath string) ([]Question, error) {
	absolutePath, err := findAbsolutePath(csvPath)
	if nil != err {
		return nil, err
	}
	file, err := os.Open(absolutePath)
	if nil != err {
		println("Error while opening the file on path ", csvPath)
		return nil, err
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if nil != err {
		println("Error while reading csv file ", csvPath)
		return nil, err
	}
	var data []Question
	for _, quizRecord := range records {
		ques := Question{}
		ques.question = quizRecord[0]
		ques.answer = quizRecord[1]
		data = append(data, ques)
	}
	return data, nil
}

func readArguments() (string, int) {
	fileName := flag.String("quizInput", "default.csv", "A csv file in the format of {Question, answer} which defaults to default.csv")
	limit := flag.Int("limit", 30, "Time limit for each question")
	flag.Parse()
	return *fileName, *limit
}

func readInput(input chan string) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input <- scanner.Text()
	}
}

func askQuestion(quizRecord Question,
	timer <-chan time.Time,
	answerPrompt <-chan string) (int, error) {
	fmt.Printf("%s: ", quizRecord.question)
	for {
		select {
		case <-timer:
			return -1, fmt.Errorf("TIME OUT")
		case answer := <-answerPrompt:
			score := 0
			if strings.Compare(
				strings.Trim(strings.ToLower(answer), "\n"),
				quizRecord.answer) == 0 {
				score = 1
			} else {
				return 0, nil
			}
			return score, nil
		}
	}
}

func askQuestions(questions []Question, timelimit int) (int, error) {
	finalScore := 0
	timer := time.NewTimer(time.Duration(timelimit) * time.Second)
	answerPrompt := make(chan string)
	go readInput(answerPrompt)

	for _, question := range questions {
		result, err := askQuestion(question, timer.C, answerPrompt)
		if nil != err && -1 == result {
			return finalScore, err
		}
		finalScore += result
	}
	return finalScore, nil
}

func main() {
	fileName, limit := readArguments()
	Questions, err := readCSV(fileName)
	if nil != err {
		println(err.Error())
		os.Exit(3)
	}
	if nil == Questions {
		println("No questions found in the file !!! Exiting")
		os.Exit(3)
	}
	println("Starting your quiz now be ready!!!")
	finalScore, err := askQuestions(Questions, limit)
	fmt.Printf("You scored %d out of %d", finalScore, len(Questions))
}
