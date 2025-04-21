package mlist

import "testing"

//Тест на создание списка
func TestShopListCreation(t *testing.T) {
	var shp_ls ShoppingListItem
	if &shp_ls == nil {
		t.Errorf("Экземпляр структуры не создан")
	}
}

//Тест проверку полей структуры списка
func TestShopListFields(t *testing.T) {
	var shp_ls ShoppingListItem
	if &shp_ls.Name == nil {
		t.Errorf("Поле \"Name\" не существует")
	}
	if &shp_ls.Amount == nil {
		t.Errorf("Поле \"Amount\" не существует")
	}
	if &shp_ls.Unit == nil {
		t.Errorf("Поле \"Unit\" не существует")
	}
}

//Тест на создание среза с элементами списка
func TestShopListSliceCreation(t *testing.T) {
	var shpLsSlice ShoppingList
	if &shpLsSlice == nil {
		t.Errorf("Массив не создан")
	}
}

//Тест на добавление элемента в срез
func TestShopListSliceAdding(t *testing.T) {
	var shopListSlice ShoppingList
	item := shopListSlice.AddItem("Milk", 4.5, "L")
	if item != (ShoppingListItem{Name: "Milk", Amount: 4.5, Unit: "L"}) {
		t.Errorf("Ошибка в добавлении элемента")
	} else {
		t.Log(item)
	}
}

//Тест на добавление элемента в срез из экземпляра
func TestShopListSliceAddingFromInst(t *testing.T) {
	var shopListSlice ShoppingList
	item := shopListSlice.AddItemFromInst(ShoppingListItem{Name: "Sugar", Amount: 10, Unit: "kg"})
	if item != (ShoppingListItem{Name: "Sugar", Amount: 10, Unit: "kg"}) {
		t.Errorf("Ошибка в добавлении элемента")
	} else {
		t.Log(item)
	}
}

//Тест на конвертацию элемента списка в массив строк
func TestShopListItemToStringSlice(t *testing.T) {
	shopListItemString := ShoppingListItem{Name: "Lemon", Amount: 2, Unit: "pcs"}.ToString()
	if shopListItemString[0] != "Lemon" ||
		shopListItemString[1] != "2" ||
		shopListItemString[2] != "pcs" {
		t.Errorf("Ошибка конвертиции в строку")
	} else {
		t.Log(shopListItemString)
	}
}

//Тест на удаление элемента из массива
func TestShopListRemoveItem(t *testing.T) {
	var shopListSlice ShoppingList
	shopListSlice.AddItem("Meat", 2.5, "kg")
	shopListSlice.AddItem("Soda", 1.5, "L")
	shopListSlice.AddItem("Tomatos", 4, "pcs")
	shopListSlice.RemoveItem(1)

	if shopListSlice[0] != (ShoppingListItem{Name: "Meat", Amount: 2.5, Unit: "kg"}) ||
		shopListSlice[1] != (ShoppingListItem{Name: "Tomatos", Amount: 4, Unit: "pcs"}) {
		t.Errorf("Ошибка удаления элемента")
	} else {
		t.Log(shopListSlice)
	}
}
