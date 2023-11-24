package wetest

import (
	"github.com/cdvelop/model"
)

type weTest struct {
	*model.Object

	// casos de usos registrados
	uses_cases []model.Response

	milliseconds   int
	current_object *model.Object

	model.BackEndActions

	model.FrontEndActions
}
