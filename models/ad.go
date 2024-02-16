package models

type Ad struct {
	Id         int
	Title      string
	StartAt    string
	EndAt      string
	Conditions AdCondition
}

type AdCondition struct {
	AgeStart int
	AgeEnd   int
	Gender   string
	Platform []string
	Country  []string
}
