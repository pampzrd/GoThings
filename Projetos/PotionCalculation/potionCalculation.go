/* Utility tool to calculate how many potions can we buy on Ragnarok Online :3 */
/* Our Lord is good! */
package main

import "fmt"

func main() {
	entry := 0
	ismerchant := true
	tpOfPotion := 0
	potionPriceArr := [4]int{10, 50, 7, 38}
	potionPrice := 0
	amountOfZenys := 0
	spentZenys := 0
	totalPotions := 0

	fmt.Println("Are you a merchant ? 1-YES 2-NO")
	fmt.Scanln(&entry)
	switch entry {
	case 1:
		ismerchant = true
	case 2:
		ismerchant = false
	default:
		ismerchant = true
	}

	fmt.Println("1-Red Potion 2-Orange Potion")
	fmt.Scanln(&tpOfPotion)
	if ismerchant == true && tpOfPotion == 1 {
		potionPrice = potionPriceArr[2]
	} else if ismerchant == true && tpOfPotion == 2 {
		potionPrice = potionPriceArr[3]
	} else if ismerchant != true && tpOfPotion == 1 {
		potionPrice = potionPriceArr[0]
	} else if ismerchant != true && tpOfPotion == 2 {
		potionPrice = potionPriceArr[1]
	}

	// It receives a amount of Zenys available to spend

	fmt.Println("Insert amount of Zenys:")
	fmt.Scanln(&amountOfZenys)

	i := 0
	for {
		spentZenys += potionPrice

		if spentZenys > amountOfZenys {
			break
		}
		totalPotions++
		i++
	}

	// Calculates change amount of Zenys.
	ZenysChange := amountOfZenys - (totalPotions * potionPrice)
	fmt.Println("The quantity of potions is ", totalPotions, ".")
	fmt.Println(ZenysChange, " change in Zenys.")

}
