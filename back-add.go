package wetest

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/strings"
)

// ej: os.Args
func AddBackendApiE2E(h *model.Handlers, osArgs []string) (out *WeTest, err string) {
	const this = "AddBackendE2EtestsAPI error "
	var required_tests []string
	for _, arg := range osArgs {
		if strings.Contains(arg, "test:") != 0 {
			var new_test string
			err = strings.ExtractTwoPointArgument(arg, &new_test)
			if err != "" {
				return nil, this + err
			}
			required_tests = append(required_tests, new_test)
		}
	}

	// h.Log("EN AddBackendE2EtestsAPI:", add_backend_api)

	if len(required_tests) != 0 {

		t, err := AddE2ETestHandler(h)
		if err != "" {
			return nil, this + err
		}
		t.required_tests = required_tests

		h.AddObjects(t.Object)

		out = t
	}

	return
}
