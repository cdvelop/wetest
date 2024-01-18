package wetest

import (
	"strconv"
)

func (h *WeTest) ExecBackendActions() (err string) {
	const e = "ExecBackendActions error "

	for i, uc := range h.backend_uses_case {

		if _, test_required := h.name_required_tests[uc.TestName]; test_required {

			var this = "-backend-test " + strconv.Itoa(i) + " " + uc.TestName + " "

			for _, t := range uc.TestActions {

				if t.Set_backend_date != "" {
					h.Log(this+h.Set_backend_date+":", t.Set_backend_date)
					h.SetDate(t.Set_backend_date)
				} else if t.DB_Insert_Data_Table != "" {

					// fmt.Println("DB_Insert_Data_Table DATA:", t.Data)

					err = h.DB_Insert_Data_Table(this, t)
				}

				if err != "" {
					return this + e + err
				}
			}

		} else {
			h.Log(e+"nombre test:"+uc.TestName, "no existe")
		}

	}

	return
}
