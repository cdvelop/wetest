package wetest

func (h *WeTest) AddUsesCaseTest(uses_cases ...UseCase) {
	h.uses_case = append(h.uses_case, uses_cases...)
}
