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
	return shoppingListItem{Name: "Milk", Amount: 4.5, Unit: "L"}
}
