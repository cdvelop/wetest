package wetest

import "github.com/cdvelop/model"

func (h *weTest) AddUsesCaseTest(uses_cases ...model.Response) {

	// separamos los casos de uso
	for _, uc := range uses_cases {
		if uc.Action == h.Backend_action && !h.FrontendExecution {
			h.execBackendActions(uc)

		} else if uc.Action == h.Frontend_action && h.FrontendExecution {

			h.frontend_uses_case = append(h.frontend_uses_case, uc)

		}
	}

}

func (h weTest) TestActions() (model.BackEndActions, model.FrontEndActions) {
	return h.BackEndActions, h.FrontEndActions
}
