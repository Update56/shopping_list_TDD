package mlist

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// структура элмента списка
type ShoppingListItem struct {
	Name   string  //продукт
	Amount float64 //количество
	Unit   string  //единица измерений (шт, кг, л и т.д.)
}

// срез из элементов списка
type ShoppingList []ShoppingListItem

// функция добавления элемента
func (s *ShoppingList) AddItem(name string, amount float64, unit string) ShoppingListItem {
	*s = append(*s, ShoppingListItem{Name: name, Amount: amount, Unit: unit})
	return (*s)[len(*s)-1]
}

// функция добавления элемента из экзепляра
func (s *ShoppingList) AddItemFromInst(item ShoppingListItem) ShoppingListItem {
	*s = append(*s, item)
	return (*s)[len(*s)-1]
}

// функция конвертации в массив строк
func (s ShoppingListItem) ToString() []string {
	return []string{s.Name, fmt.Sprintf("%g", s.Amount), s.Unit}
}

// функция удаления элемента
func (s *ShoppingList) RemoveItem(num int) {
	num = num - 1
	if num > len(*s) {
		fmt.Println("Не верный элемент")
		return
	}
	*s = append((*s)[:num], (*s)[num+1:]...)
}

func (s ShoppingList) Print() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Amount", "Unit"})
	for _, v := range s {
		table.Append(v.ToString())
	}
	table.Render()
}
