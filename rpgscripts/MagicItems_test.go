package rpgscripts

import (
	"reflect"
	"strings"
	"testing"
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
		want MagicItem
	}{
		// TODO: Add test cases.
		{"First one", seeded, MagicItem{"wand", "Gold which can cast Fireball"}},
		{"First one", seeded, MagicItem{"wand", "Bone which can cast Cone of Cold"}},
		{"First one", seeded, MagicItem{"wand", "Wood which can cast Fireball"}},
		{"Not Seeded - Weapon", notSeeded, MagicItem{"weapon", "Bow +4"}},
		{"Not Seeded - Wand", notSeeded, MagicItem{"wand", "Copper which can cast Paralyzer"}},
		{"Not Seeded - Armor", notSeeded, MagicItem{"armor", "Studded Leather Armbraces +2"}},
		{"Not Seeded - Ring (Weapon)", notSeeded, MagicItem{"ring", "Attacking +3"}},
		{"Not Seeded - Ring (Wand)", notSeeded, MagicItem{"ring", "Glass which can cast Cone of Cold"}},
		{"Not Seeded - Ring (Armor)", notSeeded, MagicItem{"ring", "Defense +2"}},
		{"Not Seeded - Ring (Potion)", notSeeded, MagicItem{"ring", "Telekenisis"}},
		{"Not Seeded - Potion", notSeeded, MagicItem{"potion", "Speed"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateMagicItem(tt.args.params...); !reflect.DeepEqual(got, tt.want) {
				if strings.Contains(tt.name, "Not Seeded") {
					for i := 10000; i > 0; i-- {
						if !reflect.DeepEqual(got, tt.want) {
							got = CreateMagicItem()
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
