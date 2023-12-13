package wetest

func (h WeTest) resultTests(err string) {

	var OK = "OK"
	if err != "" {
		OK = ""
	}

	h.Log("- RESULTADO tests -", OK, err)

}
