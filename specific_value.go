package charts

import "github.com/gopherjs/gopherjs/js"

type SpecificValue struct {
	*js.Object
	Title    string      `js:"title"`
	LineType string      `js:"line_type"`
	Value    interface{} `js:"value"`
}

func NewSpecificValue(title string, lineType string, value interface{}) *SpecificValue {
	new := &SpecificValue{Object: js.Global.Get("Object").New()}
	new.Title = title
	new.LineType = lineType
	new.Value = value
	return new
}
