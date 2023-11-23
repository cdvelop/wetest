package wetest

import "github.com/cdvelop/model"

func (h *weTest) RunModuleTests(usesCaseTests []model.Response, result func(err string)) {
	h.Log("- INICIANDO tests -")
	h.runTestsSequentially(usesCaseTests, 0, 0, result)
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
