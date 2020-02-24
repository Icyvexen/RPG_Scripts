package items

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Icyvexen/rpgscripts/helpers"
)

//MagicItem - possesses Type (string) and Text (string).
type MagicItem struct {
	Type string
	Text string
}

//newMagicItem - taking type and text of item, returns pointer to the MagicItem
func newMagicItem(iT string, text string) Item {
	mi := MagicItem{iT, text}
	return mi
}

//Describe - return
func (mi MagicItem) Describe() string {
	ret := fmt.Sprintf("%+v", mi)

	return string(ret)
}

//Declare the "constant" maps up here for cleaner code
var (
	magicItemTypes = map[int]string{
		0: "Potion",
		1: "Armor",
		2: "Wand",
		3: "Weapon",
		4: "Ring",
	}

	armorTypes = map[int]string{
		0: "Hand",
		1: "Chest",
		2: "Head",
		3: "Armbraces",
		4: "Pants",
		5: "Shoes",
		6: "Tower Shield",
		7: "Shield",
	}

	armorComposition = map[int]string{
		0: "Wooden",
		1: "Leather",
		2: "Studded Leather",
		3: "Chainmail",
		4: "Platemail",
		5: "Cloth",
	}

	wandEffects = map[int]string{
		0: "Magic Missile",
		1: "Fireball",
		2: "Paralyzer",
		3: "Cone of Cold",
	}

	weaponTypes = map[int]string{
		0: "Sword",
		1: "Arrows",
		2: "Hammer",
		3: "Mace",
		4: "Dagger",
		5: "Bow",
	}

	wandAndStaffComposition = map[int]string{
		0: "Wood",
		1: "Iron",
		2: "Copper",
		3: "Gold",
		4: "Glass",
		5: "Bone",
	}
)

//Declare the RNG for this section.
var (
	rs  = rand.NewSource(time.Now().UnixNano())
	rng = rand.New(rs)
)

/*
CreateMagicItem - generates a magic item and returns the text description as a MagicItem.
*/
func CreateMagicItem(params helpers.Parameters) Item {
	var iType, iText string
	iType, iText = generateItemType(len(magicItemTypes), params)
	item := newMagicItem(iType, iText)
	return item

}

func generateItemType(n int, params helpers.Parameters) (string, string) {

	itemRandom := rng.Int63n(5)
	var iType, text string
	switch itemRandom {
	case 0:
		iType = "potion"
		var pot Potion
		text = pot.Generate(params).Describe()
	case 1:
		iType = "armor"
		text = generateArmor()
	case 2:
		iType = "wand"
		text = generateWand()
	case 3:
		iType = "weapon"
		text = generateWeapon()
	case 4:
		iType = "ring"
		text = generateRing(params)
	}
	return iType, text
}

func generateArmor() string {
	aT := armorTypes[rand.Intn(len(armorTypes))]
	aC := armorComposition[rand.Intn(len(armorComposition))]
	mod := rand.Intn(4) + 1

	return itemDescriptorType(aC, " ", aT, " +", strconv.Itoa(mod))
}

func generateWand() string {
	e := wandEffects[rand.Intn(len(wandEffects))]
	t := wandAndStaffComposition[rand.Intn(len(wandAndStaffComposition))]
	return itemDescriptorType(t, " which can cast ", e)
}

func generateWeapon() string {
	wt := weaponTypes[rand.Intn(len(weaponTypes))]
	mod := rand.Intn(4) + 1
	return itemDescriptorType(wt, " +", strconv.Itoa(mod))
}

func generateRing(params helpers.Parameters) string {
	mit := magicItemTypes
	comp := mit[rand.Intn(len(mit))]
	mod := rand.Intn(4) + 1
	switch comp {
	case "Weapon":
		var ret string = "Attacking +"
		ret += strconv.Itoa(mod)
		return ret
	case "Wand":
		return generateWand()
	case "Armor":
		var ret string = "Defense +"
		ret += strconv.Itoa(mod)
		return ret
	case "Potion":
		var pot Potion
		return pot.Generate(params).Describe()
	}
	return generateRing(params)
}

func itemDescriptorType(str ...string) string {
	var sb strings.Builder
	for _, s := range str {
		sb.WriteString(s)
	}
	return sb.String()

}
