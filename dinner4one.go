/******
Program: Dinner for one
Author: Naomi
Date Completed: []
Description: []
******/

package main

import "fmt"

// defining the combo
type Combo struct {
	Entree string
	Side   string
	Drink  string
	Size   int
	Price  float64
}

// psuedo constructor
func NewCombo(entree string, side string, drink string, size int, price float64) Combo {
	var c Combo
	c.Entree = entree
	c.Side = side
	c.Drink = drink
	c.Size = size
	c.Price = price

	return c
}

// non prompt memu ad seleect combo - it displays the all the combos and what the user selected and will return what the user selected
// there is supposed to be no parameters
func PromptMenuAndSelectCombo(c Combo) Combo {

	fmt.Println("Welcome to Eclectic Drive-Thru. What would you like to order?")

	fmt.Println("Combo 1:")
	fmt.Println("Entree: Hamburger")
	fmt.Println("Side: Fries")
	fmt.Println("Price: $5.99")

	fmt.Println("Combo 2:")
	fmt.Println("Entree: Burrito")
	fmt.Println("Side: Rice")
	fmt.Println("Price: $4.99")

	fmt.Println("Combo 3:")
	fmt.Println("Entree: Salad")
	fmt.Println("Side: Breadsticks")
	fmt.Println("Price: $4.49")

	fmt.Println("Please select your order number")
	var input int
	fmt.Scanln(&input)

	if input == 1 {
		var c1 Combo = NewCombo("Hamburger", "Rice", 5.00)
		return c1
	} else if input == 2 {
		var c2 Combo = NewCombo("Burrito", "Rice", 4.99)
		return c2
	} else if input == 3 {
		var c3 Combo = NewCombo("Salad", "Breadsticks", 4.49)
		return c3
	}

}

func UpdateSizeAndDrink (c Combo) Combo{
	fmt.Println("Size ungrade prices: Small $0, Medium $2, Large $3")
	fmt.Println("What size would you like? (1 = small, 2 = Medium, 3 = Large)")
	var drinksize int
	fmt.Scanln(&drinksize)
	if drinksize == 1{
		c
	}
}
func(c Combo) SetSize(){ // 

}
func(c Combo) SetDrink(){ // implementation must account for the user leaving the drink selection blank. 
	
}
func (c Combo) Display() { //displayed to the end user to allow them
	fmt.Println("Entree: ", c.Entree, "\nSide: ", c.Side, "\nPrice: ", c.Price)
}

func (c Combo) DisplayFull() {
	fmt.Println("Here is your order:")
	fmt.Println("Entree: ",c.Entree,"\nSide: ",c.Size,"\nDrink: ",\n)
}


func main() {

	//var c1 Combo = PromptMenuAndSelectCombo()

//	PromptMenuAndSelectCombo()
	var  c1 Combo = NewCombo("Hamburger","fries")
	
	fmt.Println("Size upgrade prices: Small $0, Medium $2, Large $3")
	fmt.Println("What size would you like? (1 = small, 2 = Medium, 3 = Large)")
	var drinksize int
	fmt.Scanln(&drinksize)
	if drinksize == 1{
		
	}

}
