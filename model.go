package wetest

import (
	"github.com/cdvelop/model"
)

type weTest struct {
	*model.Object

	// casos de usos registrados
	uses_cases []model.Response

	// actions
	Name_Object          string `Legend:"Nombre Objeto"`
	Clicking_ID          string `Legend:"Click en Objeto"` // ej Click
	Click_menu_module    string `Legend:"Click en Modulo"`
	Click_object_element string `Legend:"Click en Elemento html"` // ej: "button[name='btn_cancel']"

	Wait string `Legend:"Espera en milisegundos"` // ej 200

	milliseconds   int
	current_object *model.Object
	continue_test  *bool

	// resultTest     func(err string)
}
