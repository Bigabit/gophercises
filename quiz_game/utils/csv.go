package quizgame

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

type Result interface {
	Update(int, bool)
}

type Test interface {
	Print()
	Test(int)
}

//	Stores indices of questions
type TestResults struct {
	results []bool
}

type Question struct {
	problem string
	answer  string
}

type Questions struct {
	questions map[int]Question
	length    int
}

func input() (in string, err error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(">> ")
	in, err = reader.ReadString('\n')
	return
}

// Initialize questions
func NewQuestions(path string) (Test, error) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	questions := Questions{
		questions: make(map[int]Question),
		length:    0,
	}
	for index, record := range records {
		questions.questions[index] = Question{
			problem: record[0],
			answer:  record[1],
		}
		questions.length += 1
	}

	return &questions, err
}

func NewResult(length int) (Result, error) {
	return &TestResults{
		results: make([]bool, length),
	}, nil
}

func (q Questions) Print() {
	qs := q.questions
	for i := 0; i < q.length; i++ {
		fmt.Print(i, ".")
		line := qs[i]
		fmt.Println(line.problem, "=", line.answer)
	}
}

func (q Questions) Test(seconds int) {
	qs := q.questions
	results, err := NewResult(q.length)
	if err != nil {
		fmt.Println(err)
	}
	pass := 0

	timer := time.NewTimer(time.Duration(seconds) * time.Second)

problemloop:
	for i := 0; i < q.length; i++ {
		line := qs[i]
		fmt.Println(line.problem)

		ans := make(chan string)
		//	prevent blocking program so that timer expiration can execute
		go func() {
			answer, err := input()
			if err != nil {
				fmt.Println(err)
			}
			ans <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("You've run out of time")
			break problemloop
		case ans := <-ans:
			orig := strings.TrimSpace(line.answer)
			ans = strings.Trim(ans, "\n")

			if ans == orig {
				results.Update(i, true)
				pass += 1
			} else {
				results.Update(i, false)
			}
		}
	}
	fmt.Println("Passed:", pass)
	fmt.Println("Failed:", q.length-pass)
}

func (tr TestResults) Update(i int, correct bool) {
	tr.results[i] = correct
}
