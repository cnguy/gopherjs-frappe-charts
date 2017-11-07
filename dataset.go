package charts

import "github.com/gopherjs/gopherjs/js"

func NewDataset(title string, color string, values []interface{}) (dataset *Dataset) {
	dataset = &Dataset{Object: js.Global.Get("Object").New()}
	dataset.Title = title
	dataset.Color = color
	dataset.Values = values
	return dataset
}

type Dataset struct {
	*js.Object
	Title  string        `js:"title"`
	Color  string        `js:"color"`
	Values []interface{} `js:"values"`
}
