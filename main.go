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
			EditList(shopList[choice-1])
		} else if choice == 0 {
			shopList = append(shopList, mlist.ShoppingList{})
			shopList[len(shopList)-1].AddItemFromInst(InputItem())
			fmt.Printf("Список № %d создан", len(shopList)-1)
		} else {
			fmt.Println("Неверный пункт меню")
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

func EditList(list mlist.ShoppingList) {
	for {
		list.Print()
		fmt.Println("1.Добавить новый элемент\n2.Удалить элемент\n0.Выйти в меню")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 0:
			return
		case 1:
			list.AddItemFromInst(InputItem())
		case 2:
			fmt.Print("Введите номер элемента: ")
			fmt.Scan(&choice)
			list.RemoveItem(choice)
		default:
			fmt.Println("Неверный пункт меню")
		}
	}
}
