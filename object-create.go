package wetest

func (h WeTest) DB_Insert_Data_Table(this string, t TestAction) (err string) {
	return h.CreateObjectsInDB(t.DB_Insert_Data_Table, false, t.Data)
}

func (h WeTest) Object_Post_Create(this string, t TestAction, result func(err string)) {
	const a = "Object_Post_Create "

	h.Log(this + a + h.obj.ObjectName)

	// enviamos la petición al servidor para contar los elementos
	h.SendOneRequest("POST", "create", h.obj.ObjectName, t.Data, func(serverResp []map[string]string, err string) {

		if err != "" {
			result(a + err)
			return
		}

		h.Log("TEST CREATE DATA RESPUESTA SERVIDOR:", serverResp)

		// if len(serverResp) != t.Count.Expected.DBBackend {
		// 	result(a + "se esperaba:" + strconv.Itoa(t.Count.Expected.DBBackend) + " pero se obtuvo:" + strconv.Itoa(len(serverResp)) + " elementos en la base de datos del servidor")
		// 	return
		// }

		result("")
	})

}
