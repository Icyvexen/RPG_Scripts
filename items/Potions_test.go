package items_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/Icyvexen/rpgscripts/helpers"
	"github.com/Icyvexen/rpgscripts/items"
)

func TestPotionGenerate(t *testing.T) {
	var seed int64 = 7868455476845
	var seeded helpers.Parameters
	seeded = seeded.NewParameters(seed)
	notSeeded := helpers.Parameters{}

	tests := []helpers.TCStrings{
		{Name: "Not Seeded - Effect", Ar: notSeeded, Want: "Levitation"},
		{Name: "Not Seeded - Container", Ar: notSeeded, Want: "Clay Pot"},
		{Name: "Not Seeded - Color", Ar: notSeeded, Want: "Ivory"},
		{Name: "Not Seeded - Liquid", Ar: notSeeded, Want: "Ichorous"},
		{Name: "Not Seeded - Taste", Ar: notSeeded, Want: "Mutton"},
	}

	pot := items.Potion{}
	LoopForTestPotions(t, pot.Generate, tests)
}

func LoopForTestPotions(t *testing.T, tf func(helpers.Parameters) items.Potion, tests []helpers.TCStrings) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if strings.Contains(tt.Name, "Not Seeded") {
				i := 10000
				for i > 0 {
					got := tf(tt.Ar)
					if strings.Contains(got.Describe(), tt.Want) {
						fmt.Println("SUCCESS +++ ", tt.Name)
						break
					}
					i--
				}
				if i == 0 {
					t.Errorf("%v --- Could not generate \"%v\" after 10000 iterations. Maybe the test is broken?", tt.Name, tt.Want)
				}
			} else {
				got := tf(tt.Ar)
				if strings.Contains(got.Describe(), tt.Want) {
					t.Errorf("%+v --- Generated \"%+v\", want \"%+v\" .", tt.Name, got, tt.Want)
				}
			}
		})
	}
}
