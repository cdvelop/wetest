package wetest

func (t *WeTest) RunE2EfrontTests() {
	const this = "RunE2EfrontTests "

	// string_content, err := h.SelectContent("meta[name='JsonBootTests']", "content")
	// if err != "" {
	// 	h.Log(this + err)
	// 	return
	// }
	// // h.Log("CONTENIDO:", string_content)

	// if string_content == "none" {
	// 	h.Log(this + "sin tests para ejecutar")
	// 	return
	// }
	t.Log("- INICIANDO tests -")

	for _, v := range t.uses_case {

		t.Log(v.Description)
	}
	// t.FetchAdapter.SendOneRequest("GET", "read", t.ObjectName+"2", map[string]string{t.Required_tests: "ok"}, func(result []map[string]string, err string) {

	// 	if err != "" {
	// 		t.Log(this + err)
	// 		return
	// 	}

	// })
	// t.FetchAdapter.SendOneRequest("POST", "read", t.ObjectName, map[string]string{t.Required_tests: "ok"}, func(result []map[string]string, err string) {

	// 	if err != "" {
	// 		t.Log(this + err)
	// 		return
	// 	}

	// })

	// usesCaseTests, err := h.DecodeResponses([]byte(string_content))
	// if err != "" {
	// 	h.Log(this + err)
	// 	return
	// }

	t.runTestsSequentially(t.uses_case, 0, 0, t.resultTests)
}

func (h *WeTest) runTestsSequentially(useCaseTests []UseCase, testIndex int, dataIndex int, result func(err string)) {
	if testIndex >= len(useCaseTests) {
		// Todas las pruebas han sido completadas
		result("")
		return
	}

	test := useCaseTests[testIndex]

	if dataIndex >= len(test.TestActions) {
		// Todas las pruebas de este caso de uso han sido completadas
		h.runTestsSequentially(useCaseTests, testIndex+1, 0, result)
		return
	}

	h.processUnitTest(dataIndex, test.TestActions[dataIndex], func(err string) {
		if err != "" {
			// h.Log("NO CONTINUAR TEST:", dataIndex)
			result(err)
			return
		}

		// h.Log("CONTINUANDO TEST:", dataIndex)

		// Procesar la siguiente prueba de manera recursiva
		h.runTestsSequentially(useCaseTests, testIndex, dataIndex+1, result)
	})
}
