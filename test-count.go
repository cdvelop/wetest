package wetest

import (
	"strconv"

	"github.com/cdvelop/model"
)

func (h WeTest) Count_Elements(this string, t TestAction, result func(err string)) {
	const e = "Count Elements "

	h.Log(this + e + h.obj.ObjectName)

	// var ex = "se esperaba:" + strconv.Itoa(expected_value) + " pero se obtuvo:"
	// h.Log("SE ESPERA:", expected_value)

	ui_elements, err := h.obj.CountViewElements()
	if err != "" {
		result(e + err)
		return
	}

	if ui_elements != t.Count.Expected.UI {
		result(e + "se esperaba:" + strconv.Itoa(t.Count.Expected.UI) + " pero se obtuvo:" + strconv.Itoa(ui_elements) + " elementos en la interfaz del usuario")
	} else {

		// consultamos la base de datos frontend
		h.ReadAsyncDataDB(&model.ReadParams{
			FROM_TABLE: h.obj.Table,
			WHERE:      t.Count.ReadParams,
		}, func(F *model.ReadResults, err string) {

			if err != "" {
				result(e + err)
				return
			}

			if len(F.ResultsString) != t.Count.Expected.DBFrontend {
				result(e + "se esperaba:" + strconv.Itoa(t.Count.Expected.DBFrontend) + " pero se obtuvo:" + strconv.Itoa(len(F.ResultsString)) + " elementos en la db del frontend")
				return
			}

			// enviamos la petición al servidor para contar los elementos
			h.SendOneRequest("POST", "read", h.obj.ObjectName, t.Count.ReadParams, func(serverResp []map[string]string, err string) {

				if err != "" {
					result(e + err)
					return
				}

				if len(serverResp) != t.Count.Expected.DBBackend {
					result(e + "se esperaba:" + strconv.Itoa(t.Count.Expected.DBBackend) + " pero se obtuvo:" + strconv.Itoa(len(serverResp)) + " elementos en la base de datos del servidor")
					return
				}

				result("")
			})

		})
	}

}
