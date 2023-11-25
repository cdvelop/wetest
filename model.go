package wetest

import (
	"github.com/cdvelop/model"
)

type weTest struct {
	*model.Object

	// casos de usos frontend registrados
	frontend_uses_case []model.Response

	milliseconds   int
	current_object *model.Object

	model.BackEndActions

	model.FrontEndActions
}
