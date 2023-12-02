package wetest

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

var wt *WeTest

func AddE2ETestHandler(h *model.Handlers) (t *WeTest, err string) {

	if wt == nil {

		m := &model.Module{ModuleName: "wetest"}

		wt = &WeTest{}

		object, err := object.BuildObjectFromStruct(wt)
		if err != "" {
			return nil, "AddE2ETestHandler " + err
		}
		wt.Object.Module = m
		m.Handlers = h
		wt.Object = object

	}

	return wt, ""
}
