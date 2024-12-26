package dbmodel

type Quest struct {
	ID          int
	Title       string
	Description string
	Objective   string
	Reward      string
	PlayerName  string
	PlayerClass string
}
