package wetest

import "github.com/cdvelop/model"

func (h *weTest) RunE2EfrontTests() {
	const this = "RunE2EfrontTests "

	string_content, err := h.SelectContent("meta[name='JsonBootTests']", "content")
	if err != "" {
		h.Log(this + err)
		return
	}
	// h.Log("CONTENIDO:", string_content)

	if string_content == "none" {
		h.Log(this + "sin tests para ejecutar")
		return
	}
	h.Log("- INICIANDO tests -")

	usesCaseTests, err := h.DecodeResponses([]byte(string_content))
	if err != "" {
		h.Log(this + err)
		return
	}

	h.runTestsSequentially(usesCaseTests, 0, 0, h.resultTests)
}

func (h *weTest) runTestsSequentially(useCaseTests []model.Response, testIndex int, dataIndex int, result func(err string)) {
	if testIndex >= len(useCaseTests) {
		// Todas las pruebas han sido completadas
		result("")
		return
	}

	test := useCaseTests[testIndex]

	if dataIndex >= len(test.Data) {
		// Todas las pruebas de este caso de uso han sido completadas
		h.runTestsSequentially(useCaseTests, testIndex+1, 0, result)
		return
	}

	h.processUnitTest(dataIndex, test.Data[dataIndex], func(err string) {
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
