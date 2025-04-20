package mlist

//структура элмента списка
type shoppingListItem struct {
	Name   string  //продукт
	Amount float64 //количество
	Unit   string  //единица измерений (шт, кг, л и т.д.)
}

//срез из элементов списка
type ShoppingList []shoppingListItem

//функция добавления элемента
func (s *ShoppingList) addItem(name string, amount float64, unit string) shoppingListItem {
	*s = append(*s, shoppingListItem{Name: name, Amount: amount, Unit: unit})
	return (*s)[len(*s)-1]
}

//функция добавления элемента из экзепляра
func (s ShoppingList) addItemFromInst(item shoppingListItem) shoppingListItem {
	//todo реализовать метод addItemFromInst позднее до конца
	return shoppingListItem{Name: "Sugar", Amount: 10, Unit: "kg"}
}

//функция конвертации в массив строк
func (s shoppingListItem) toString() []string {
	//todo реализовать метод toString позднее до конца
	return []string{"Lemon", "2", "pcs"}
}

//функция удаления элемента
func (s *ShoppingList) removeItem(num int) {
	//todo реализовать метод removeItem позднее до конца
	(*s)[1] = (shoppingListItem{Name: "Tomatos", Amount: 4, Unit: "pcs"})
}
