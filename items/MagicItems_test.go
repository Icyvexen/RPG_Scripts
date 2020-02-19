package items_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/Icyvexen/rpgscripts/items"
)

func TestCreateMagicItem(t *testing.T) {
	type args struct {
		params []int64
	}
	//This seed is defined as such for consistent testing purposes
	var seed int64 = 7868455476845
	seeded := args{params: []int64{seed}}
	notSeeded := args{}

	tests := []struct {
		name string
		args args
		want items.MagicItem
	}{
		// TODO: Add test cases.
		{"Seeded - 1", seeded, items.MagicItem{"wand", "Gold which can cast Fireball"}},
		{"Seeded - 2", seeded, items.MagicItem{"wand", "Bone which can cast Cone of Cold"}},
		{"Seeded - 3", seeded, items.MagicItem{"wand", "Wood which can cast Fireball"}},
		{"Not Seeded - Weapon", notSeeded, items.MagicItem{"weapon", "Bow +4"}},
		{"Not Seeded - Wand", notSeeded, items.MagicItem{"wand", "Copper which can cast Paralyzer"}},
		{"Not Seeded - Armor", notSeeded, items.MagicItem{"armor", "Studded Leather Armbraces +2"}},
		{"Not Seeded - Ring (Weapon)", notSeeded, items.MagicItem{"ring", "Attacking +3"}},
		{"Not Seeded - Ring (Wand)", notSeeded, items.MagicItem{"ring", "Glass which can cast Cone of Cold"}},
		{"Not Seeded - Ring (Armor)", notSeeded, items.MagicItem{"ring", "Defense +2"}},
		{"Not Seeded - Ring (Potion)", notSeeded, items.MagicItem{"ring", "Telekenisis"}},
		{"Not Seeded - Potion", notSeeded, items.MagicItem{"potion", "Speed"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := items.CreateMagicItem(tt.args.params...); !reflect.DeepEqual(got, tt.want) {
				if strings.Contains(tt.name, "Not Seeded") {
					for i := 10000; i > 0; i-- {
						if !reflect.DeepEqual(got, tt.want) {
							got = items.CreateMagicItem()
						} else {
							break
						}
					}
					if !reflect.DeepEqual(got, tt.want) {
						t.Errorf("Was unable to generate desired item in 10000 iterations. This may indicate a faulty test condition.")
					}
				} else {
					t.Errorf("CreateMagicItem() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
