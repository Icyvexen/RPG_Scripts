package names_test

import (
	"testing"

	"github.com/Icyvexen/rpgscripts/helpers"
)

func TestNewBook(t *testing.T) {

	//This seed is defined as such for consistent testing purposes
	var seed int64 = 7868455476845
	var seeded helpers.Parameters
	seeded = seeded.NewParameters(seed)
	notSeeded := helpers.Parameters{}

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
	}
	helpers.LoopForTestStrings(t, items.NewBook, tests)
}
