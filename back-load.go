package wetest

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

	h.Log("EN API load_tests:", load_tests)

	if load_tests {

		// codificamos la data a json
		front_data, err := h.EncodeResponses(h.frontend_uses_case...)
		if err != "" {
			h.Log(this + err)
			return
		}

		return string(front_data)

	}

	return "none"
}
