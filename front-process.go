package wetest

import (
	"strconv"
)

func (h *WeTest) processUnitTest(from *UseCase, i int, t TestAction, result func(err string)) {

	var this = "-test " + strconv.Itoa(i) + " "

	if from.RunBeforeView && t.Set_frontend_date != "" {
		if t.Set_frontend_date != "" {
			h.Log(this+h.Set_frontend_date+":", t.Set_frontend_date)
			h.SetDate(t.Set_frontend_date)
		}
		result("")
		return
	}

	h.milliseconds += 100
	var err string

	if t.Wait != "" {
		wait, e := strconv.Atoi(t.Wait)
		if e == nil {
			h.milliseconds += wait
		}
	}

	var object_use string
	if t.Name_object_use != "" {
		object_use = t.Name_object_use
	}

	if t.Form_complete != "" {
		object_use = t.Form_complete
	}

	if object_use != "" {
		h.obj, err = h.MainHandlerGetObjectByName(object_use)
		if err != "" {
			result(err)
			return
		}
	}

	if len(t.Clear_all_table_data) != 0 {
		h.Log(this+"Clear_all_table_data:", t.Clear_all_table_data)
		err = h.ClearAllTableDataInDB(t.Clear_all_table_data...)
		result(err)
		return
	}

	if h.obj == nil {
		result("error objeto para realizar test no definido ej: t.Name_object_use: ObjectName")
		return
	}
	// test asynchronous actions:
	h.WaitFor(h.milliseconds, func() {
		var err string

		if t.Click_menu_module != "" {
			h.Log(this+h.Click_menu_module+":", t.Click_menu_module)
			err = h.ElementClicking(h.QuerySelectorMenuModule(t.Click_menu_module))

		} else if t.Clicking_ID != "" {
			h.Log(this+h.Clicking_ID+":", t.Clicking_ID)
			err = h.obj.ClickingID(t.Clicking_ID)

		} else if t.Click_object_element != "" {
			h.Log(this+h.Click_object_element+":", t.Click_object_element)
			err = h.obj.ClickObjectElement(t.Click_object_element)

		} else if t.Form_complete != "" {

			h.Log(this+h.Form_complete+":", t.Form_complete)

			if len(t.Data) == 0 {
				err = this + "no llego data para completar formulario"
			} else {

				err = h.FormComplete(true, false)
			}

		} else if t.Count != nil {
			h.Count_Elements(this, t, result)
			return
		}

		result(err)
	})

}
