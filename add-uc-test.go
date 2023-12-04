package wetest

func (h *WeTest) AddFrontUsesCaseTest(uses_cases ...UseCase) {

	for _, uc := range uses_cases {
		if !uc.BackendExecute {
			h.uses_case = append(h.uses_case, uc)
		}
	}
}
