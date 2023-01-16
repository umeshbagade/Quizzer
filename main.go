package main

import (
	"fmt"
	"time"
)

func main() {
	url := "http://localhost:3000/questions"

	problems := questionPuller(url)

	tobj := time.NewTimer(10 * time.Second * time.Duration(len(problems))) // Time for all the questions --> 1 question => 10 seconds

	correctAns := 0
ProblemLoop:

	for i, problem := range problems {
		var answer string
		fmt.Printf("\nProblem %d: %s", i+1, problem.Question)
		fmt.Printf("\n a. %s \n b. %s \n c. %s \n d. %s \n Select Option 'a','b','c','d' : ",
			problem.Options.A, problem.Options.B, problem.Options.C, problem.Options.D)
		ansC := make(chan string)

		go func() {
			fmt.Scanf("%s", &answer)
			ansC <- answer
		}()
		select {
		case <-tobj.C:
			fmt.Print("timer done")
			break ProblemLoop
		case iAns := <-ansC:
			if iAns == problem.Answer {
				correctAns++
			}
			if i == len(problems)-1 {
				fmt.Print("Done all Questions ")
			}

		}

	}

	fmt.Printf("Correct %d out of %d", correctAns, len(problems))

}
