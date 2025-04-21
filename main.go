package main

import (
	"fmt"
	"shopping_list/mlist"
)

func main() {

	shopList := make([]mlist.ShoppingList, 1) //создаём слайс из листов
	shopList[0].AddItemFromInst(InputItem())  //в первый вставляем новый элемент
	fmt.Println("Первый список создан")

	for {
		fmt.Println("0. Добавить новый список")
		fmt.Printf("Выберите список (1 - %d): ", len(shopList))
		var choice int
		fmt.Scan(&choice)
		if choice >= 1 && choice <= len(shopList) {
			fmt.Println(shopList[choice-1])
		} else if choice == 0 {
			shopList = append(shopList, mlist.ShoppingList{})
			shopList[len(shopList)-1].AddItemFromInst(InputItem())
			fmt.Printf("Список № %d создан", len(shopList)-1)
		}
	}
}

func InputItem() mlist.ShoppingListItem {
	fmt.Println("Введите название продукта")
	var name string
	fmt.Scan(&name)

	fmt.Println("Введите количество продукта")
	var amount float64
	fmt.Scan(&amount)

	fmt.Println("Введите ед. измерения")
	var unit string
	fmt.Scan(&unit)

	return mlist.ShoppingListItem{Name: name, Amount: amount, Unit: unit}
}
