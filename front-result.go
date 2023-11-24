package wetest

func (h weTest) resultTests(err string) {
	h.Log("- RESULTADO tests -")

	if err != "" {
		h.Log("error", err)
	} else {
		h.Log("OK")
	}

}
