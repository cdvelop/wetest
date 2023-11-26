package wetest

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

func AddE2ETestAdapter(h *model.Handlers) (err string) {

	m := &model.Module{ModuleName: "wetest"}

	t := &weTest{}

	object, err := object.BuildObjectFromStruct(t)
	if err != "" {
		return "AddE2ETestAdapter " + err
	}
	t.Object.Module = m
	m.Handlers = h
	t.Object = object

	h.TestAdapter = t

	return
}
