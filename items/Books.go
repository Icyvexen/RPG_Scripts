package names

import (
	"math/rand"
	"time"

	"github.com/Icyvexen/rpgscripts/helpers"
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
func NewBook(params helpers.Parameters) string {
	var rs rand.Source
	var sb string
	s := params.Seed()
	if s > 0 {
		rs = rand.NewSource(s)
	} else {
		rs = rand.NewSource(time.Now().UnixNano())
	}
	rng := rand.New(rs)

	choice := rand.Intn(11)

	cov := cover.String(rng)
	page := pageComposition.String(rng)
	view := pov.String(rng)
	topic := bookTopic.String(rng)

	switch choice {
	case 0: //cover pages
		sb = helpers.StringCombiner("A book bound by ", cov, ", with ", page, " pages.")
	case 1: //cover pov
		sb = helpers.StringCombiner("A book with ", cov, " binding, as told from a(n) ", view, " perspective.")
	case 2: //cover topic
		sb = helpers.StringCombiner("A book is here with ", cov, " binding about ", topic, ".")
	case 3: //pages pov
		sb = helpers.StringCombiner("A book with ", page, "pages, from a(n) ", view, " perspective.")
	case 4: //pages topic
		sb = helpers.StringCombiner("A book with ", page, " pages about ", topic, ".")
	case 5: //pov topic
		sb = helpers.StringCombiner("A book from a(n) ", view, " perspective for ", topic, ".")
	case 6: //cov pages pov
		sb = helpers.StringCombiner("Here is a book in ", cov, " binding, pages made of ", page, ", from a(n) ", view, " perspective.")
	case 7: //cov pages topic
		sb = helpers.StringCombiner("A book with ", cov, " binding and set on ", page, " pages, on the topic of ", topic, ".")
	case 8: //cov pov topic
		sb = helpers.StringCombiner("This book has ", cov, " binding from a(n) ", view, " perspective, on the topic of ", topic, ".")
	case 9: //pages pov topic
		sb = helpers.StringCombiner("A book with ", page, " pages from a(n) ", view, " perspective, detailing thoughts on ", topic, ".")
	case 10: //all
		sb = helpers.StringCombiner("A book with ", cov, " binding, pages made of ", page, ", over a(n) ", view, " perspective on the subjects of ", topic, ".")
	}

	ret := sb
	time.Sleep(time.Microsecond)
	return ret
}
