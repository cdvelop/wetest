package wetest

func (h *WeTest) AddUsesCaseTest(uses_cases ...UseCase) {

	for _, uc := range uses_cases {
		if uc.BackendExecute { //BACKEND

			h.backend_uses_case = append(h.backend_uses_case, uc)

		} else { //FRONTEND
			if uc.RunBeforeView {
				h.front_uc_before_view = append(h.front_uc_before_view, uc)

			} else {
				h.front_uc_after_view = append(h.front_uc_after_view, uc)
			}

		}
	}
}
