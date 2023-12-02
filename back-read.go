package wetest

import "github.com/cdvelop/model"

func (t WeTest) Read(u model.User, data ...map[string]string) (out []map[string]string, err string) {

	t.Log("TESS REQUERIDO:", data)

	for _, r := range t.required_tests {
		data[0][r] = "ok"
	}

	return
}
