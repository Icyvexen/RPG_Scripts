package names

import (
	"math/rand"
	"strings"
	"time"
)

type dictionary map[int]string

func (d dictionary) String(r *rand.Rand) string {
	ret := r.Intn(len(d) + 1)
	return d[ret]
}

//List of book parts
var (
	cover = dictionary{
		1:  "Wood",
		2:  "Leather",
		3:  "Bone",
		4:  "Exotic Hide",
		5:  "Ivory",
		6:  "Steel",
		7:  "Coral",
		8:  "Leaves",
		9:  "Oilskin",
		10: "Tree Bark",
		11: "Sheepskin",
		12: "Bronze",
		13: "Human Skin",
		14: "Copper",
		15: "Seashell",
		16: "Stone",
		17: "Cloth",
		18: "Bamboo",
		19: "Snakeskin",
		20: "Prismatic Glass",
	}

	pageComposition = dictionary{
		1:  "Paper",
		2:  "Parchment",
		3:  "Vellum",
		4:  "Copper Sheet",
		5:  "Paper",
		6:  "Onion Skin",
		7:  "Linen",
		8:  "Clay Tablet",
		9:  "Plastic Sheet",
		10: "Paper",
		11: "Parchment",
		12: "Rice Paper",
		13: "Erasable Surface",
		14: "Papyrus",
		15: "Paper",
		16: "Wax Tablet",
		17: "Etched Glass",
		18: "Display Screen",
		19: "Parchment",
		20: "Paper",
	}

	pov = dictionary{
		1:  "Impoverished",
		2:  "Halfling",
		3:  "Wizard's",
		4:  "Eastern",
		5:  "Fey",
		6:  "Tyrant's",
		7:  "Elven",
		8:  "Dragon's",
		9:  "Human",
		10: "Northern",
		11: "Middle Class",
		12: "Woman's",
		13: "Southern",
		14: "Dwarven",
		15: "Lich's",
		16: "Orcish",
		17: "Rebel's",
		18: "Western",
		19: "Gnomish",
		20: "Wealthy",
	}

	bookTopic = dictionary{
		1:  "Astronomy",
		2:  "Cookbook",
		3:  "Mathematics",
		4:  "Ethics",
		5:  "Alchemy",
		6:  "Comedic Novel",
		7:  "Almenac",
		8:  "Philosophy",
		9:  "Naturalism",
		10: "Politics",
		11: "Agriculture",
		12: "Geology",
		13: "Occultism",
		14: "History",
		15: "Demonology",
		16: "Religion",
		17: "Music",
		18: "Drama",
		19: "Etiquette",
		20: "Astrology",
	}
)

//NewBook - returns a ship name
func NewBook(params ...int64) string {
	var rs rand.Source
	var sb strings.Builder
	if len(params) > 0 {
		rs = rand.NewSource(params[0])
	} else {
		rs = rand.NewSource(time.Now().UnixNano())
	}
	rng := rand.New(rs)

	comp := rand.Intn(11)
	switch comp {
	case 0: //cover pages
		sb.WriteString("A book bound by ")
		sb.WriteString(cover.String(rng))
		sb.WriteString(", with ")
		sb.WriteString(pageComposition.String(rng))
		sb.WriteString(" pages.")
	case 1: //cover pov
		sb.WriteString("A book with ")
		sb.WriteString(cover.String(rng))
		sb.WriteString(" binding, as told from a(n) ")
		sb.WriteString(pov.String(rng))
		sb.WriteString(" perspective.")
	case 2: //cover topic
		sb.WriteString("A book is here with ")
		sb.WriteString(cover.String(rng))
		sb.WriteString(" binding about ")
		sb.WriteString(bookTopic.String(rng))
		sb.WriteString(".")
	case 3: //pages pov
		sb.WriteString("A book with ")
		sb.WriteString(pageComposition.String(rng))
		sb.WriteString(" pages, from a(n) ")
		sb.WriteString(pov.String(rng))
		sb.WriteString(" perspective.")
	case 4: //pages topic
		sb.WriteString("A book with ")
		sb.WriteString(pageComposition.String(rng))
		sb.WriteString(" pages about ")
		sb.WriteString(bookTopic.String(rng))
		sb.WriteString(".")
	case 5: //pov topic
		sb.WriteString("A book from a(n) ")
		sb.WriteString(pov.String(rng))
		sb.WriteString(" perspective for ")
		sb.WriteString(bookTopic.String(rng))
		sb.WriteString(".")
	case 6: //cov pages pov
		sb.WriteString("Here is a book in")
		sb.WriteString(cover.String(rng))
		sb.WriteString(" binding, pages made of ")
		sb.WriteString(pageComposition.String(rng))
		sb.WriteString(", from a(n) ")
		sb.WriteString(pov.String(rng))
		sb.WriteString(" perspective.")
	case 7: //cov pages topic
		sb.WriteString("A book with ")
		sb.WriteString(cover.String(rng))
		sb.WriteString(" binding and set on ")
		sb.WriteString(pageComposition.String(rng))
		sb.WriteString(" pages, on the topic of ")
		sb.WriteString(bookTopic.String(rng))
		sb.WriteString(".")
	case 8: //cov pov topic
		sb.WriteString("This book has ")
		sb.WriteString(cover.String(rng))
		sb.WriteString(" binding from a(n) ")
		sb.WriteString(pov.String(rng))
		sb.WriteString(" perspective, on the topic of ")
		sb.WriteString(bookTopic.String(rng))
		sb.WriteString(".")
	case 9: //pages pov topic
		sb.WriteString("A book with ")
		sb.WriteString(pageComposition.String(rng))
		sb.WriteString(" pages from a(n) ")
		sb.WriteString(pov.String(rng))
		sb.WriteString(" perspective, detailing thoughts on ")
		sb.WriteString(bookTopic.String(rng))
		sb.WriteString(".")
	case 10: //all
		sb.WriteString("A book with ")
		sb.WriteString(cover.String(rng))
		sb.WriteString(" binding, pages made of ")
		sb.WriteString(pageComposition.String(rng))
		sb.WriteString(", over a(n) ")
		sb.WriteString(pov.String(rng))
		sb.WriteString(" perspective on the subjects of ")
		sb.WriteString(bookTopic.String(rng))
		sb.WriteString(".")
	}

	ret := sb.String()
	time.Sleep(time.Microsecond)
	return ret
}
