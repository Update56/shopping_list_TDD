package list

import "testing"

//Тест на создание списка
func TestSendMessage(t *testing.T) {
	var shp_ls shopping_list
	if &shp_ls == nil {
		t.Errorf("Экземпляр структуры не создан")
	}
}
