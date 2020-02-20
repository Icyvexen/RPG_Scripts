package names_test

import (
	"testing"

	"github.com/Icyvexen/rpgscripts/helpers"
	"github.com/Icyvexen/rpgscripts/names"
)

func TestNewBook(t *testing.T) {

	//This seed is defined as such for consistent testing purposes
	var seed int64 = 7868455476845
	seeded := []int64{seed}
	notSeeded := []int64{}

	//Only 1 seeded one due to implementation
	tests := []helpers.TCStrings{
		{Name: "Not Seeded - cov pages", Ar: notSeeded, Want: "A book bound by"},
		{Name: "Not Seeded - cov pov", Ar: notSeeded, Want: " binding, as told from a(n)"},
		{Name: "Not Seeded - cov topic", Ar: notSeeded, Want: "A book is here with"},
		{Name: "Not Seeded - pages pov", Ar: notSeeded, Want: " pages, from a(n) "},
		{Name: "Not Seeded - pages topic", Ar: notSeeded, Want: " pages about "},
		{Name: "Not Seeded - pov topic", Ar: notSeeded, Want: " perspective for "},
		{Name: "Not Seeded - cov pages pov", Ar: notSeeded, Want: "Here is a book in"},
		{Name: "Not Seeded - cov pages topic", Ar: notSeeded, Want: " binding and set on "},
		{Name: "Not Seeded - cov pov topic", Ar: notSeeded, Want: "This book has "},
		{Name: "Not Seeded - pages pov topic", Ar: notSeeded, Want: " perspective, detailing thoughts on "},
		{Name: "Not Seeded - all", Ar: notSeeded, Want: " perspective on the subjects of "},
		{Name: "Seeded - Only one", Ar: seeded, Want: "Human perspective"},
	}
	helpers.LoopForTestStrings(t, names.NewBook, tests)
}
