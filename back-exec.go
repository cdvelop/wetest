package wetest

import (
	"fmt"
	"strconv"
)

func (h *WeTest) ExecBackendActions(uses_case ...UseCase) {
	fmt.Println("EJECUTANDO execBackendActions....")
	for i, uc := range uses_case {
		if uc.BackendExecute {

			var this = "-backend-test " + strconv.Itoa(i) + " "

			for _, t := range uc.TestActions {

				h.Log(this+h.Set_server_date+":", t.Set_server_date)

				if t.Set_server_date != "" {
					h.SetDate(t.Set_server_date)
				}

			}
		}
	}
}
