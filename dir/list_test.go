package list

import "testing"

//Тест на создание списка
func TestShopListCreation(t *testing.T) {
	var shp_ls shopping_list
	if &shp_ls == nil {
		t.Errorf("Экземпляр структуры не создан")
	}
}

//Тест проверку полей структуры списка
func TestShopListFields(t *testing.T) {
	var shp_ls shopping_list
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
