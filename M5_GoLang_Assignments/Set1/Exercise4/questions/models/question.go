package models

type Question struct {
	ID      int
	Q       string
	Opts    map[string]string
	Correct string
}
