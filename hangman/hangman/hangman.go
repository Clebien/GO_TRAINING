package hangman

import "strings"

type Game struct {
	State        string   // game state
	Letters      []string // Letters in the world to find
	FoundLetters []string //Good letters
	UsedLetters  []string
	TurnsLeft    int
}

func New(turns int, word string) *Game {
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	g := &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}

	return g

}
