package wetest

import "github.com/cdvelop/model"

func (h weTest) execBackendActions(uc_actions []model.Response) {

	for _, uc := range uc_actions {
		for _, a := range uc.Data {
			if a[h.Set_server_date] != "" {
				h.SetDate(a[h.Set_server_date])
			}

		}
	}
}
