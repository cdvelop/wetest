package wetest

import (
	"github.com/cdvelop/model"
)

type WeTest struct {
	*model.Object

	name_required_tests map[string]bool // nombre de la pruebas solicitadas a ejecutar

	// casos de usos a evaluar
	backend_uses_case    []UseCase
	front_uc_before_view []UseCase
	front_uc_after_view  []UseCase

	milliseconds int
	obj          *model.Object

	TestAction
}

type UseCase struct {
	BackendExecute bool
	RunBeforeView  bool   // si se ejecuta antes de cargar la vista
	TestName       string //ej "session" = run params "test:session"
	Description    string
	TestActions    []TestAction
}

type TestAction struct {
	//Espera en milisegundos ej 200,2000 = 2 seg
	Wait int
	// setear la fecha del servidor a un dia especifico ej: 2023-12-24
	Set_backend_date string
	// setear la fecha en el fronted a un dia especifico ej: 2023-12-24
	Set_frontend_date string
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
	// borrar la data de tablas por su nombre
	Clear_all_table_data []string
	// contar cuantos elementos existen en: la vista, db cliente, db servidor
	Count *Count
	//data pare realizar la prueba
	Data map[string]string
}

type Count struct {
	ReadParams []map[string]string //where
	Expected   ExpectedCount
}

type ExpectedCount struct {
	UI         int //interfaz usuario
	DBFrontend int
	DBBackend  int
}
