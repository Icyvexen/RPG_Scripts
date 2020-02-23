package items_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/Icyvexen/rpgscripts/items"
)

//TCMagicItems - stores Name (string), Ar(guments, []int64), and Want (MagicItem)
type TCMagicItems struct {
	Name string
	Ar   []int64
	Want items.Item
}

//LoopForTestMagicItems - test each condition provided
func LoopForTestMagicItems(t *testing.T, tf func(...int64) items.Item, tests []TCMagicItems) {
	t.Helper()
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			if strings.Contains(tt.Name, "Not Seeded") {
				i := 10000
				for i > 0 {
					got := tf(tt.Ar...)
					if !reflect.DeepEqual(got, tt.Want) {
						fmt.Println("SUCCESS +++ ", tt.Name)
						break
					}
					i--
				}
				if i == 0 {
					t.Errorf("%v --- Could not generate \"%v\" after 10000 iterations. Maybe the test is broken?", tt.Name, tt.Want)
				}
			} else {
				got := tf(tt.Ar...)
				if !reflect.DeepEqual(got, tt.Want) {
					t.Errorf("%v --- Generated \"%v\", want \"%v\" .", tt.Name, got, tt.Want)
				}
			}
		})
	}
}

func TestCreateMagicItem(t *testing.T) {
	//This seed is defined as such for consistent testing purposes
	var seed int64 = 7868455476845
	seeded := []int64{seed}
	notSeeded := []int64{}

	tests := []TCMagicItems{
		// TODO: Add test cases.
		{Name: "Seeded - 1", Ar: seeded, Want: items.MagicItem{"wand", "Gold which can cast Fireball"}},
		{Name: "Seeded - 2", Ar: seeded, Want: items.MagicItem{"wand", "Bone which can cast Cone of Cold"}},
		{Name: "Seeded - 3", Ar: seeded, Want: items.MagicItem{"wand", "Wood which can cast Fireball"}},
		{Name: "Not Seeded - Weapon", Ar: notSeeded, Want: items.MagicItem{"weapon", "Bow +4"}},
		{Name: "Not Seeded - Wand", Ar: notSeeded, Want: items.MagicItem{"wand", "Copper which can cast Paralyzer"}},
		{Name: "Not Seeded - Armor", Ar: notSeeded, Want: items.MagicItem{"armor", "Studded Leather Armbraces +2"}},
		{Name: "Not Seeded - Ring (Weapon)", Ar: notSeeded, Want: items.MagicItem{"ring", "Attacking +3"}},
		{Name: "Not Seeded - Ring (Wand)", Ar: notSeeded, Want: items.MagicItem{"ring", "Glass which can cast Cone of Cold"}},
		{Name: "Not Seeded - Ring (Armor)", Ar: notSeeded, Want: items.MagicItem{"ring", "Defense +2"}},
		{Name: "Not Seeded - Ring (Potion)", Ar: notSeeded, Want: items.MagicItem{"ring", "Telekenisis"}},
		{Name: "Not Seeded - Potion", Ar: notSeeded, Want: items.MagicItem{"potion", "Speed"}},
	}

	LoopForTestMagicItems(t, items.CreateMagicItem, tests)
}
