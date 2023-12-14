package wetest

import (
	"strconv"

	"github.com/cdvelop/model"
)

func (h WeTest) Count_Elements(this string, t TestAction, result func(err string)) {
	const a = "Count Elements "

	h.Log(this + a + h.obj.ObjectName)

	// var ex = "se esperaba:" + strconv.Itoa(expected_value) + " pero se obtuvo:"
	// h.Log("SE ESPERA:", expected_value)

	ui_elements, err := h.obj.CountViewElements()
	if err == "" {

		if ui_elements != t.Count.Expected.UI {
			result(a + "se esperaba:" + strconv.Itoa(t.Count.Expected.UI) + " pero se obtuvo:" + strconv.Itoa(ui_elements) + " elementos en la interfaz del usuario")
			return
		} else {

			// consultamos la base de datos frontend
			h.ReadAsyncDataDB(model.ReadParams{
				FROM_TABLE:      h.obj.Table,
				WHERE:           []string{t.Count.WHERE},
				SEARCH_ARGUMENT: t.Count.ARGUMENT,
			}, func(F model.ReadResult) {

				if F.Error != "" {
					result(a + F.Error)
					return
				}

				if len(F.DataString) != t.Count.Expected.DBFrontend {
					result(a + "se esperaba:" + strconv.Itoa(t.Count.Expected.DBFrontend) + " pero se obtuvo:" + strconv.Itoa(len(F.DataString)) + " elementos en la db del frontend")
				} else {
					// enviamos la petición al servidor para contar los elementos
					h.SendOneRequest("POST", "read", h.obj.ObjectName, map[string]string{t.Count.WHERE: t.Count.ARGUMENT}, func(serverResp []map[string]string, err string) {

						if err != "" {
							result(a + err)
							return
						}

						if len(serverResp) != t.Count.Expected.DBBackend {
							result(a + "se esperaba:" + strconv.Itoa(t.Count.Expected.DBBackend) + " pero se obtuvo:" + strconv.Itoa(len(serverResp)) + " elementos en la base de datos del servidor")
						}
					})
				}
			})
		}
	}

}
