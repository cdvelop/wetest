package wetest

import "github.com/cdvelop/model"

//ej: os.Args
func (h *weTest) LoadE2Etests(osArgs []string) (e2eTests string) {
	const this = "LoadE2Etests error "
	var load_tests bool
	for _, arg := range osArgs {
		if arg == "test" {
			h.Log("ARGUMENTO test encontrado")
			load_tests = true
		}
	}

	if load_tests {

		var uc_frontend []model.Response
		var uc_backend []model.Response

		// separamos los casos de uso
		for _, uc := range h.uses_cases {
			if uc.Action == h.Frontend_action {
				uc_frontend = append(uc_frontend, uc)
			} else if uc.Action == h.Backend_action {
				uc_backend = append(uc_backend, uc)
			}
		}

		// codificamos la data a json
		front_data, err := h.EncodeResponses(uc_frontend...)
		if err != "" {
			h.Log(this + err)
			return
		}

		e2eTests = string(front_data)

		h.execBackendActions(uc_backend)
	}

	return
}

func (h weTest) LoadE2EtestsNONE() (uses_case []model.Response) {
	return h.uses_cases
}
