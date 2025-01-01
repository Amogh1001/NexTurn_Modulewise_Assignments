package employees

import (
	"question-management/questions/services"
)

func AddQuestions() {
	services.AddQuestions()
}

func ShowQuestion(id int) {
	services.ShowQuestion(id)
}

// GetAllQuestions returns all employees
func GetAllQuestions() {
	services.GetAllQuestions()
}
