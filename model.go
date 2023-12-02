package wetest

import (
	"github.com/cdvelop/model"
)

type WeTest struct {
	*model.Object

	Required_tests string
	required_tests []string // nombre de la pruebas solicitada a ejecutar

	// casos de usos a evaluar
	uses_case []UseCase

	milliseconds   int
	current_object *model.Object

	TestAction
}

type UseCase struct {
	BackendExecute bool
	Module         string
	Description    string
	TestActions    []TestAction
}

type TestAction struct {
	//Espera en milisegundos ej 200,2000 = 2 seg
	Wait string
	// setear la fecha del servidor a un dia especifico ej: 2023-12-24
	Set_server_date string
	// objeto a usar para la acción
	Name_object_use string
	// click id de un objeto especifico
	Clicking_ID string // ej Click
	// click en le menu del modulo
	Click_menu_module string
	// click en elementos del modulo del objeto ej: "button[name='btn_cancel']"
	Click_object_element string
	//Completar formulario
	Form_complete string
	// añadir data a la existente
	Form_existing_add string
	//data pare realizar la prueba
	Data map[string]string
}
