package charts

import "github.com/gopherjs/gopherjs/js"

// NewDataset allows us to quickly create Dataset objects with the required values.
func NewDataset(title string, values []interface{}) (dataset *Dataset) {
	dataset = &Dataset{Object: js.Global.Get("Object").New()}
	dataset.Title = title
	dataset.Values = values
	return dataset
}

// Dataset represents a single dataset of the JS object 'data'.
// https://frappe.github.io/charts/
// Ctrl + F "datasets".
// Wrap this inside a dataset array.
type Dataset struct {
	*js.Object
	Title  string        `js:"title"`
	Values []interface{} `js:"values"`
}
