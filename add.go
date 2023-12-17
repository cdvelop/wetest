package wetest

import (
	"github.com/cdvelop/model"
	"github.com/cdvelop/object"
)

var wt *WeTest

func AddE2ETestHandler(log model.Logger) (t *WeTest, err string) {

	if wt == nil {

		m := &model.Module{
			ModuleName:  "wetest",
			MainHandler: &model.MainHandler{Logger: log},
		}

		wt = &WeTest{}

		object, err := object.BuildObjectFromStruct(wt)
		if err != "" {
			return nil, "AddE2ETestHandler " + err
		}

		wt.Object.Module = m
		wt.Object = object

	}

	return wt, ""
}

func (w *WeTest) AddHandlersToWetest(h *model.MainHandler) {
	*w.Module.MainHandler = *h
}
