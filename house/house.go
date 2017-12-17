package house

import (
	"fmt"
)

type Creature struct {
	Name   string
	Action string
}

var creatures = []Creature{
	Creature{"house that Jack built", "is"},
	Creature{"malt", "lay in"},
	Creature{"rat", "ate"},
	Creature{"cat", "killed"},
	Creature{"dog", "worried"},
	Creature{"cow with the crumpled horn", "tossed"},
	Creature{"maiden all forlorn", "milked"},
	Creature{"man all tattered and torn", "kissed"},
	Creature{"priest all shaven and shorn", "married"},
	Creature{"rooster that crowed in the morn", "woke"},
	Creature{"farmer sowing his corn", "kept"},
	Creature{"horse and the hound and the horn", "belonged to"},
}

func verse(number int) string {
	c1 := creatures[number]
	c2 := creatures[number-1]
	if number == 1 {
		return fmt.Sprintf("that %s the %s.", c1.Action, c2.Name)
	}
	return fmt.Sprintf("that %s the %s\n%s", c1.Action, c2.Name, verse(number-1))
}

func Verse(number int) string {
	c := creatures[number-1]
	if number == 1 {
		return fmt.Sprintf("This %s the %s.", c.Action, c.Name)
	}
	return fmt.Sprintf("This is the %s\n%s", c.Name, verse(number-1))
}

func Song() string {
	res := ""
	for i := 1; i < len(creatures); i++ {
		res += Verse(i)
		res += "\n\n"
	}
	res += Verse(len(creatures))
	return res
}
