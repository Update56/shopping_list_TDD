package list

//структура элмента списка
type shoppingListItem struct {
	Name   string  //продукт
	Amount float64 //количество
	Unit   string  //единица измерений (шт, кг, л и т.д.)
}

//срез из элементов списка
type ShoppingList []shoppingListItem

//функция добавления элемента
func (s ShoppingList) addItem(name string, amount float64, unit string) shoppingListItem {
	//todo реализовать метод addItem позднее до конца
	return shoppingListItem{Name: "Milk", Amount: 4.5, Unit: "L"}
}

//функция добавления элемента из экзепляра
func (s ShoppingList) addItemFromInst(item shoppingListItem) shoppingListItem {
	//todo реализовать метод addItemFromInst позднее до конца
	return shoppingListItem{Name: "Sugar", Amount: 10, Unit: "kg"}
}
