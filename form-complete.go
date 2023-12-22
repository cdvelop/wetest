package wetest

func (h WeTest) SiMulateUserFormComplete(t *TestAction, this string) (err string) {

	h.Log(this+h.Form_complete+":", t.Form_complete)

	if len(t.Data) == 0 {
		err = this + "no llego data para completar formulario"
	} else {
		err = h.FormComplete(t.Form_complete, t.Data, true, false)
	}

	return
}
