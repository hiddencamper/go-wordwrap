package wordwrap

import (
	"testing"
)

func TestWordWrap(t *testing.T) {
	s := "\nSpaceTraders is a headless API and fleet-management game where players can work together or against each other to trade, explore, expand, and conquer in a dynamic and growing universe. Build your own UI, write automated scripts, or just play the game from the comfort of your terminal. The game is currently in alpha and is under active development."
	s = WordWrap(s, 80)
	t.Log(s)
	s = "\nSpaceTraders is a headless API and fleet-management game\nwhere players can work together or against each other to trade, explore, expand, and conquer in a dynamic and growing universe. Build your own UI, write automated scripts, or just play the game from the comfort of your terminal. The game is currently in alpha and is under active development."
	s = WordWrap(s, 80)
	t.Log(s)
	s = "\nabcde fghjiklmnopqrstuvwsyz"
	s = WordWrap(s, 5)
	t.Log(s)
	s = WordWrap(s, 1)
	t.Log(s)
	s = WordWrap(s, 0)
	t.Log(s)
}
