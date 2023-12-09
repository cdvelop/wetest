package wetest

var test_name_run = ""

const repository = `github.com/cdvelop/wetest`

func AddFlags() map[string]string {

	return map[string]string{
		repository + ".test_name_run": "",
	}
}
