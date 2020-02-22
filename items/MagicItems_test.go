package items_test

import (
	"testing"

	"github.com/Icyvexen/rpgscripts/helpers"
	"github.com/Icyvexen/rpgscripts/items"
)

func TestCreateMagicItem(t *testing.T) {
	//This seed is defined as such for consistent testing purposes
	var seed int64 = 7868455476845
	seeded := []int64{seed}
	notSeeded := []int64{}

	tests := []helpers.TCMagicItems{
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

	helpers.LoopForTestMagicItems(t, items.CreateMagicItem, tests)
}
