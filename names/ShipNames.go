/*
names - This package is the name generation package.
The following can be generated:
- Shipnames
*/

package names

import (
	"math/rand"
	"strings"
	"time"
)

//List of categories
var (
	descriptive = map[int]string{
		1:  "Vigilant",
		2:  "Azure",
		3:  "Briny",
		4:  "Northern",
		5:  "Wishful",
		6:  "Flying",
		7:  "Elysian",
		8:  "Obedient",
		9:  "Impetuous",
		10: "Deep",
		11: "Stubborn",
		12: "Undying",
		13: "Magic",
		14: "Drifting",
		15: "Fair",
		16: "Dawn",
		17: "Island",
		18: "Deep",
		19: "Midnight",
		20: "Maelstrom",
	}

	possessive = map[int]string{
		1:  "Wave's",
		2:  "Sea's",
		3:  "Morning's",
		4:  "Wind's",
		5:  "Derelict's",
		6:  "Freedom's",
		7:  "Night's",
		8:  "Curmudgeon's",
		9:  "Crusader's",
		10: "Castaway's",
		11: "Admiral's",
		12: "Bastard's",
		13: "Mariner's",
		14: "Hangman's",
		15: "Distance's",
		16: "Water's",
		17: "Spirit's",
		18: "Maiden's",
		19: "Sky's",
		20: "Tempest's",
	}

	subject1 = map[int]string{
		1:  "Mariner",
		2:  "Siren",
		3:  "Saber",
		4:  "Brigand",
		5:  "Maiden",
		6:  "Drifter",
		7:  "Gypsy",
		8:  "Strider",
		9:  "Witch",
		10: "Pearl",
		11: "Angel",
		12: "Wind",
		13: "Queen",
		14: "Mariner",
		15: "Jewel",
		16: "Lady",
		17: "Hunter",
		18: "Star",
		19: "Voyager",
		20: "Rider",
	}

	subject2 = map[int]string{
		1:  "Bride",
		2:  "Daughter",
		3:  "Bounty",
		4:  "Hunger",
		5:  "Grace",
		6:  "Folly",
		7:  "Laughter",
		8:  "Glory",
		9:  "Treachery",
		10: "Voice",
		11: "Embrace",
		12: "Pride",
		13: "Peace",
		14: "Faith",
		15: "Sin",
		16: "Dream",
		17: "Longing",
		18: "Wrath",
		19: "Caress",
		20: "Kiss",
	}
)

//NewShipName - returns a ship name
func NewShipName(params ...int64) string {
	var rs rand.Source
	var sb strings.Builder
	if len(params) > 0 {
		rs = rand.NewSource(params[0])
	} else {
		rs = rand.NewSource(time.Now().UnixNano())
	}
	rng := rand.New(rs)
	desc := rng.Intn(3)
	if desc == 0 {
		d := descriptive[rng.Intn(20)+1]
		sb.WriteString(d)
		sb.WriteString(" ")
	} else if desc == 1 {
		p := possessive[rng.Intn(20)+1]
		sb.WriteString(p)
		sb.WriteString(" ")
	} else {
		d := descriptive[rng.Intn(20)+1]
		sb.WriteString(d)
		sb.WriteString(" ")
		p := possessive[rng.Intn(20)+1]
		sb.WriteString(p)
		sb.WriteString(" ")
	}
	sub := rng.Intn(2)
	if sub == 0 {
		s1 := subject1[rng.Intn(20)+1]
		sb.WriteString(s1)
	} else {
		s2 := subject2[rng.Intn(20)+1]
		sb.WriteString(s2)
	}
	ret := sb.String()
	return ret
}
