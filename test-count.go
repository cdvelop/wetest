package wetest

import (
	"strconv"

	"github.com/cdvelop/model"
)

func (h WeTest) Count_Elements(this string, t TestAction, result func(err string)) {
	const a = "Count_Elements "
	var err string
	h.Log(this+".Count_Elements:", t.Expected)

	expected_value, ok := t.Expected.(int)
	if !ok {
		err = a + "la expectativa debe ser un valor tipo int para poder contar"
	} else {
		var ex = "se esperaba:" + strconv.Itoa(expected_value) + " pero se obtuvo:"
		// h.Log("SE ESPERA:", expected_value)

		var total_elements int
		total_elements, err = h.obj.CountViewElements()
		if err == "" {

			if total_elements != expected_value {
				err = a + ex + strconv.Itoa(total_elements) + " elementos en la vista"
			} else {

				// consultamos la base de datos
				h.ReadAsyncDataDB(t.ReadDBParams, func(r model.ReadResult) {

					if r.Error != "" {
						result(a + r.Error)
						return
					}

					if len(r.DataString) != expected_value {
						result(a + ex + strconv.Itoa(len(r.DataString)) + " elementos en la base de datos del cliente")
					} else {
						// enviamos la petición al servidor para contar los elementos
						h.SendOneRequest("POST", "read", h.obj.ObjectName, map[string]string{t.ReadDBParams.WHERE[0]: t.ReadDBParams.SEARCH_ARGUMENT}, func(serverResp []map[string]string, err string) {

							if err != "" {
								result(a + err)
								return
							}

							if len(serverResp) != expected_value {
								result(a + ex + strconv.Itoa(len(serverResp)) + " elementos en la base de datos en el servidor")
							}
						})
					}
				})
			}
		}
	}
}
