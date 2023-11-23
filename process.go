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

	if t[h.Name_Object] != "" {
		h.current_object, err = h.GetObjectByName(t[h.Name_Object])
		if err != "" {
			result(err)
			return
		}
	}

	// test asynchronous actions:
	h.WaitFor(h.milliseconds, func() {
		var err string

		if t[h.Click_menu_module] != "" {
			h.Log(this+"CLICK MODULO:", t[h.Click_menu_module], i)
			err = h.ElementClicking(h.QuerySelectorMenuModule(t[h.Click_menu_module]))

		} else if t[h.Clicking_ID] != "" {
			h.Log(this+"CLICK OBJETO ID:", t[h.Clicking_ID], i)

			err = h.current_object.ClickingID(t[h.Clicking_ID])

		} else if t[h.Click_object_element] != "" {
			h.Log(this+"ClickObjectElement:", t[h.Click_object_element], i)

			err = h.current_object.ClickObjectElement(t[h.Click_object_element])

		}

		result(err)
	})

}
