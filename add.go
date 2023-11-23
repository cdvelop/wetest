package wetest

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

func Add(h *model.Handlers, m *model.Module) (t *weTest, err string) {

	t = &weTest{}

	err = object.AddToHandlerFromStructs(t, m, h)
	if err != "" {
		return nil, "wetest add " + err
	}
	h.TestAdapter = t

	return
}

func (h *weTest) AddTest(uses_case ...model.Response) {
	h.uses_cases = append(h.uses_cases, uses_case...)
}

func (h weTest) GetTests() (uses_case []model.Response) {
	return h.uses_cases
}
