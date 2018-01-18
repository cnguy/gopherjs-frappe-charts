package charts

import (
	"github.com/cnguy/gopherjs-frappe-charts/utils"
	"github.com/gopherjs/gopherjs/js"
)

// UpdateValuesArgs is a wrapper that contains the two parameters
// that chart.update_values(values, labels) requires.
// https://frappe.github.io/charts/
// Ctrl + F "update_values"
type UpdateDatasetsArgs struct {
	Values [][]*UpdateDataset
	Labels []string
}

// UpdateDataSet represents the JS object frappe-chart uses
// as parameters for chart.update_values(values, labels).
// e.g. { values: [1, 2, 3, 4] }
type UpdateDataSet struct {
	*js.Object
	Values []interface{} `js:"values"`
}

// NewUpdateValueSet is a helper to conveniently create an UpdateDataSet.
func NewUpdateDataSet(values []float64) *UpdateDataSet {
	new := &UpdateDataSet{Object: js.Global.Get("Object").New()}
	new.Values = utils.FloatSliceToInterface(values)
	return new
}
