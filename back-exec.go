package wetest

import (
	"strconv"

	"github.com/cdvelop/model"
)

func (h weTest) execBackendActions(uc_actions ...model.Response) {

	for i, uc := range uc_actions {
		var this = "-test " + strconv.Itoa(i) + " "

		for _, a := range uc.Data {
			h.Log(this+h.Set_server_date+":", a[h.Set_server_date])
			if a[h.Set_server_date] != "" {
				h.SetDate(a[h.Set_server_date])
			}

		}
	}
}
