package wetest

func (h WeTest) Object_Post_Update(this string, t TestAction, result func(err string)) {
	const a = "Object_Post_Update "

	h.Log(this + a + h.obj.ObjectName)

	// enviamos la petición al servidor para contar los elementos
	h.SendOneRequest("POST", "update", h.obj.ObjectName, t.Data, func(serverResp []map[string]string, err string) {

		if err != "" {
			result(a + err)
			return
		}

		h.Log("TEST POST UPDATE DATA RESPUESTA SERVIDOR:", serverResp)

		// if len(serverResp) != t.Count.Expected.DBBackend {
		// 	result(a + "se esperaba:" + strconv.Itoa(t.Count.Expected.DBBackend) + " pero se obtuvo:" + strconv.Itoa(len(serverResp)) + " elementos en la base de datos del servidor")
		// 	return
		// }

		result("")
	})

}
