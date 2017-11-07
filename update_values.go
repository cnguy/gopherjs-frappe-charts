package charts

import (
	"github.com/gopherjs/gopherjs/js"
)

type UpdateValuesArgs struct {
	Values []*UpdateValueSet
	Labels []interface{}
}

type UpdateValueSet struct {
	*js.Object
	Values []interface{} `js:"values"`
}

func NewUpdateValueSet(values []interface{}) *UpdateValueSet {
	new := &UpdateValueSet{Object: js.Global.Get("Object").New()}
	new.Values = values
	return new
}
