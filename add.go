package wetest

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

func AddE2ETestAdapter(h *model.Handlers) (err string) {

	m := &model.Module{ModuleName: "wetest"}

	t := &weTest{}

	object, err := object.BuildObjectFromStruct(t)
	if err != "" {
		return "AddE2ETestAdapter " + err
	}
	t.Object.Module = m
	m.Handlers = h
	t.Object = object

	h.TestAdapter = t

	return
}

func (h *weTest) AddUsesCaseTest(uses_cases ...model.Response) {

	var uc_frontend []model.Response
	var uc_backend []model.Response

	// separamos los casos de uso
	for _, uc := range uses_cases {
		if uc.Action == h.Backend_action {
			h.execBackendActions(uc)

		} else if uc.Action == h.Frontend_action {

			uc_frontend = append(uc_frontend, uc)
			h.frontend_uses_case = append(h.frontend_uses_case, uc)

			uc_backend = append(uc_backend, uc)
		}
	}

}

func (h weTest) TestActions() (model.BackEndActions, model.FrontEndActions) {
	return h.BackEndActions, h.FrontEndActions
}
