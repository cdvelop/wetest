package wetest

import (
	"strconv"
)

func (h *WeTest) ExecBackendActions() {
	const this = "ExecBackendActions error "

	for i, uc := range h.backend_uses_case {

		if _, test_required := h.name_required_tests[uc.TestName]; test_required {

			var this = "-backend-test " + strconv.Itoa(i) + " " + uc.TestName + " "

			for _, t := range uc.TestActions {

				if t.Set_backend_date != "" {
					h.Log(this+h.Set_backend_date+":", t.Set_backend_date)
					h.SetDate(t.Set_backend_date)
				}
			}

		} else {
			h.Log(this+"nombre test:"+uc.TestName, "no existe")
		}

	}
}
