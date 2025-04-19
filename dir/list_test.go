package list

import "testing"

//Тест на создание списка
func TestShopListCreation(t *testing.T) {
	var shp_ls shoppingListItem
	if &shp_ls == nil {
		t.Errorf("Экземпляр структуры не создан")
	}
}

//Тест проверку полей структуры списка
func TestShopListFields(t *testing.T) {
	var shp_ls shoppingListItem
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
	item := shopListSlice.addItem("Milk", 4.5, "L")
	if item != (shoppingListItem{Name: "Milk", Amount: 4.5, Unit: "L"}) {
		t.Errorf("Ошибка в добавлении элемента")
	} else {
		t.Log(item)
	}
}
