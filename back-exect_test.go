package wetest_test

import (
	"fmt"
	"testing"

	"github.com/cdvelop/cutkey"
	"github.com/cdvelop/logserver"
	"github.com/cdvelop/model"
	"github.com/cdvelop/timeserver"
	"github.com/cdvelop/wetest"
)

func TestServerActions(t *testing.T) {
	const date_required = "2023-08-29"

	h := &model.Handlers{
		TimeAdapter: timeserver.Add(),
		Logger:      logserver.Add(),
	}

	cutkey.AddDataConverter(h)

	err := wetest.AddE2ETestAdapter(h)
	if err != "" {
		t.Fatal(err)
		return
	}

	b, f := h.TestActions()

	usesCase := []model.Response{
		{
			Action:  b.Backend_action,
			Object:  "",
			Message: "cambiamos a fecha en el servidor",
			Data: []map[string]string{
				{b.Set_server_date: date_required},
			},
		},
		{
			Action:  f.Frontend_action,
			Object:  "",
			Message: "acciones comunes en el front",
			Data: []map[string]string{
				{f.Click_menu_module: "module_name"},
				{f.Clicking_ID: "123"},
			},
		},
	}

	h.AddUsesCaseTest(usesCase...)

	e2eTests := h.LoadE2Etests([]string{"test"})
	if e2eTests == "" {
		t.Fatalf("se esperaba que e2eTests no estuviere vació")
		return
	}

	today := h.ToDay("")
	if today != date_required {
		t.Fatalf("\nse esperaba que la fecha\n=>fuera:\n%v\n\n=>pero se obtuvo:\n%v\n", date_required, today)
		return
	}

	fmt.Println("FECHA ACTUAL:", today)

	// data_encode, err := h.EncodeResponses(usesCase...)
	// if err != "" {
	// 	t.Fatal(err)
	// 	return
	// }

	// // fmt.Printf("|||-%s-|||\n", data_encode)

	// responses, err := h.DecodeResponses(data_encode)
	// if err != "" {
	// 	t.Fatal(err)
	// 	return
	// }

	// if !reflect.DeepEqual(responses, usesCase) {
	// 	t.Fatalf("Unexpected result:\n\n=>response:\n%v\n\n=>expected:\n%v\n", responses, usesCase)
	// 	return
	// }

}
