package names_test

import (
	"testing"

	"github.com/Icyvexen/rpgscripts/helpers"
	"github.com/Icyvexen/rpgscripts/names"
)

func TestNewShipName(t *testing.T) {

	//This seed is defined as such for consistent testing purposes
	var seed int64 = 7868455476845
	var seeded helpers.Parameters
	seeded = seeded.NewParameters(seed)
	notSeeded := helpers.Parameters{}

	tests := []helpers.TCStrings{
		//NOTE: to validate the tests more quickly, check if 1 word from each section is present
		{Name: "Not Seeded - s1", Ar: notSeeded, Want: "Northern"},
		{Name: "Not Seeded - s1", Ar: notSeeded, Want: "Hangman's"},
		{Name: "Not Seeded - s1", Ar: notSeeded, Want: "Maiden"},
		{Name: "Not Seeded - s2", Ar: notSeeded, Want: "Glory"},
		{Name: "seeded", Ar: seeded, Want: "Obedient Sky's Gypsy"},
	}
	helpers.LoopForTestStrings(t, names.NewShipName, tests)
}
