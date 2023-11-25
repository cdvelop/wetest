package wetest

import "strconv"

func (h *weTest) processUnitTest(i int, t map[string]string, result func(err string)) {
	var this = "-test " + strconv.Itoa(i) + " "
	h.Log("DATA PARA EL TEST:", t)

	h.milliseconds += 100
	var err string

	if t[h.Wait] != "" {
		wait, e := strconv.Atoi(t[h.Wait])
		if e == nil {
			h.milliseconds += wait
		}
	}

	if t[h.Name_object_use] != "" {
		h.current_object, err = h.GetObjectByName(t[h.Name_object_use])
		if err != "" {
			result(err)
			return
		}
	}

	// test asynchronous actions:
	h.WaitFor(h.milliseconds, func() {
		var err string

		if t[h.Click_menu_module] != "" {
			h.Log(this+h.Click_menu_module+":", t[h.Click_menu_module])
			err = h.ElementClicking(h.QuerySelectorMenuModule(t[h.Click_menu_module]))

		} else if t[h.Clicking_ID] != "" {
			h.Log(this+h.Clicking_ID+":", t[h.Clicking_ID])
			err = h.current_object.ClickingID(t[h.Clicking_ID])

		} else if t[h.Click_object_element] != "" {
			h.Log(this+h.Click_object_element+":", t[h.Click_object_element])
			err = h.current_object.ClickObjectElement(t[h.Click_object_element])

		} else if t[h.Form_complete] != "" {
			h.Log(this+h.Form_complete+":", t[h.Form_complete])

			if len(t) < 2 {
				err = this + "no llego data para completar formulario"
				h.Log(t)
			} else {
				err = h.FormComplete(t[h.Form_complete], t)
			}

		}

		result(err)
	})

}
