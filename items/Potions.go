package items

import (
	"bytes"
	"encoding/json"
	"math/rand"

	"github.com/Icyvexen/rpgscripts/helpers"
)

//Potion - has Container, Color, Liquid, Taste, and Effect - all
type Potion struct {
	Container string
	Color     string //TODO: make this color string? Or at least an "enum"?
	Liquid    string
	Taste     string
	Effect    string
}

//Describe -- Item Interface - returns a JSON-ized string of the
func (p Potion) Describe() string {
	ret, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return bytes.NewBuffer(ret).String()
}

//Declare the different parts of a potion to choose from
var (
	potionEffects = map[int]string{
		0:  "Strength",
		1:  "Healing",
		2:  "Speed",
		3:  "Regeneration",
		4:  "Protection from Undead",
		5:  "Protection from Magic",
		6:  "Protection from Fire",
		7:  "Protection from Heat",
		8:  "Protection from Cold",
		9:  "Invisibility",
		10: "Levitation",
		11: "Telekenisis",
		12: "Speed",
		13: "Control Plants",
		14: "Control Humans",
		15: "Control Animals",
		16: "Control Undead",
	}
	potionContainer = map[int]string{
		1:  "Vial",
		2:  "Tankard",
		3:  "Ewer",
		4:  "Glass Bottle",
		5:  "Carafe",
		6:  "Bladder",
		7:  "Cruet",
		8:  "Gourd",
		9:  "Jug",
		10: "Test Tube",
		11: "Phial",
		12: "Flagon",
		13: "Canteen",
		14: "Shot Glass",
		15: "Sponge",
		16: "Metal Flask",
		17: "Skin",
		18: "Clay Pot",
		19: "Eggshell",
		20: "Beaker",
	}

	potionColor = map[int]string{
		1:  "White",
		2:  "Clear",
		3:  "Green",
		4:  "Indigo",
		5:  "Clear Reed",
		6:  "Brown",
		7:  "Purple",
		8:  "Silver",
		9:  "Red",
		10: "Clear Blue",
		11: "Pink",
		12: "Golden",
		13: "Dark Blue",
		14: "Neon Green",
		15: "Gray",
		16: "Orange",
		17: "Ivory",
		18: "Yellow",
		19: "Sapphire",
		20: "Black",
	}

	potionLiquid = map[int]string{
		1:  "Chunky",
		2:  "Bubbling",
		3:  "Boiling",
		4:  "Fizzy",
		5:  "Gritty",
		6:  "Frosty",
		7:  "Smooth",
		8:  "Viscous",
		9:  "Sparkling",
		10: "Smoking",
		11: "Runny",
		12: "Frothy",
		13: "Swirling",
		14: "Ichorous",
		15: "Lumpy",
		16: "Silty",
		17: "Skim",
		18: "Resinous",
		19: "Watery",
		20: "Sticky",
	}

	potionTaste = map[int]string{
		1:  "Chicken",
		2:  "Peppermint",
		3:  "Mud",
		4:  "Salt",
		5:  "Water",
		6:  "Lemons",
		7:  "Chalk",
		8:  "Hot Peppers",
		9:  "Hairy",
		10: "Pickle Brine",
		11: "Sweat",
		12: "Beef",
		13: "Alcohol",
		14: "Syrup",
		15: "Fish Oil",
		16: "Citrus",
		17: "Urine",
		18: "Wine",
		19: "Mutton",
		20: "Vinegar",
	}
)

//Generate - create and return a potion
func (p Potion) Generate(pot *Potion, params ...int64) *Potion {
	var rng *rand.Rand
	if params == nil {
		rng = helpers.SetupRNG()
	} else {
		rng = helpers.SetupRNG(params[0])
	}
	pot.Color = potionColor[rng.Intn(len(potionColor)+1)]
	pot.Container = potionContainer[rng.Intn(len(potionContainer)+1)]
	pot.Liquid = potionLiquid[rng.Intn(len(potionLiquid)+1)]
	pot.Taste = potionTaste[rng.Intn(len(potionTaste)+1)]
	pot.Effect = potionEffects[rng.Intn(len(potionEffects))]
	return pot
}
