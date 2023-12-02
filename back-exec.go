package wetest

import (
	"strconv"
)

func (h WeTest) execBackendActions(uses_case ...UseCase) {
	h.Log("EJECUTANDO execBackendActions....")
	for i, uc := range uses_case {
		var this = "-test " + strconv.Itoa(i) + " "

		for _, t := range uc.TestActions {

			h.Log(this+h.Set_server_date+":", t.Set_server_date)

			if t.Set_server_date != "" {
				h.SetDate(t.Set_server_date)
			}

		}
	}
}
