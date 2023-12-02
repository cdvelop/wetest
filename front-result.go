package wetest

func (h WeTest) resultTests(err string) {
	h.Log("- RESULTADO tests -")

	if err != "" {
		h.Log(err)
	} else {
		h.Log("OK")
	}

}
