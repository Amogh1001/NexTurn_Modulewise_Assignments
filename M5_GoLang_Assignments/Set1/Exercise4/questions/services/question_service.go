package services

import (
	"fmt"
	"question-management/questions/models"
)

// Questions slice to store employee data
var Questions []models.Question

func AddQuestions() {
	question1 := models.Question{
		ID: 1,
		Q:  "What kind of language is Python?",
		Opts: map[string]string{
			"A": "Interpreted",
			"B": "Compiled",
			"C": "Both",
			"D": "None of the above",
		},
		Correct: "A",
	}
	question2 := models.Question{
		ID: 2,
		Q:  "What is the capital of India?",
		Opts: map[string]string{
			"A": "Mumbai",
			"B": "Delhi",
			"C": "Kolkata",
			"D": "Chennai",
		},
		Correct: "B",
	}
	Questions = append(Questions, question1, question2)
}

func ShowQuestion(id int) {
	for _, question := range Questions {
		if question.ID == id {
			fmt.Println(question.Q)
			for k, v := range question.Opts {
				fmt.Println(k, v)
			}
		}
	}
}

// GetAllQuestions returns all employees
func GetAllQuestions() {
	totalQuestions := len(Questions)
	var correct int = 0
	for _, question := range Questions {
		ShowQuestion(question.ID)
		var ans string
		fmt.Scanln(&ans)
		if ans == question.Correct {
			correct++
		}
	}
	fmt.Printf("You got %d out of %d questions correct\n", correct, totalQuestions)
}
