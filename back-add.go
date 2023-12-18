package wetest

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

// ej: os.Args
func AddBackendApiE2E(h *model.MainHandler, osArgs []string) (out *WeTest, err string) {
	const this = "AddBackendE2EtestsAPI error "
	var name_required_tests = map[string]bool{}
	for _, arg := range osArgs {
		if strings.Contains(arg, "test:") != 0 {
			var new_test string
			err = strings.ExtractTwoPointArgument(arg, &new_test)
			if err != "" {
				return nil, this + err
			}
			name_required_tests[new_test] = true
		}
	}

	// h.Log("EN AddBackendE2EtestsAPI:", add_backend_api)

	if len(name_required_tests) != 0 {

		t, err := AddE2ETestHandler(h.Logger)
		if err != "" {
			return nil, this + err
		}
		t.AddHandlersToWetest(h)

		t.name_required_tests = name_required_tests

		h.MainHandlerAddModules(t.Object.Module)

		out = t
	}

	return
}
