package wetest

import "github.com/cdvelop/model"

func (t *WeTest) ExecuteBeforeAction(h *model.MainHandler) {
	t.Module.MainHandler = h
	t.Log("- INICIANDO Before View tests -")

	t.runTestsSequentially(t.front_uc_before_view, 0, 0, t.resultTests)
}

func (t *WeTest) RunE2EfrontTestsAfterView() {

	t.Log("- INICIANDO After View tests -")

	t.runTestsSequentially(t.front_uc_after_view, 0, 0, t.resultTests)
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

	h.processUnitTest(&test, dataIndex, test.TestActions[dataIndex], func(err string) {
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
