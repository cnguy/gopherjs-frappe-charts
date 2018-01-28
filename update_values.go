package charts

import (
	"github.com/cnguy/gopherjs-frappe-charts/utils"
	"github.com/gopherjs/gopherjs/js"
)

// UpdateValuesArgs is a wrapper that contains the two parameters
// that chart.update_values(values, labels) requires.
// https://frappe.github.io/charts/
// Ctrl + F "update_values"
type UpdateValuesArgs struct {
	Values []*UpdateValueSet
	Labels []string
}

// UpdateValueSet represents the JS object frappe-chart uses
// as parameters for chart.update_values(values, labels).
// e.g. { values: [1, 2, 3, 4] }
type UpdateValueSet struct {
	*js.Object
	Values []interface{} `js:"values"`
}

// NewUpdateValueSet is a helper to conveniently create an UpdateValueSet.
func NewUpdateValueSet(values []float64) *UpdateValueSet {
	new := &UpdateValueSet{Object: js.Global.Get("Object").New()}
	new.Values = utils.FloatSliceToInterface(values)
	return new
}

// NewUpdateValueSets is a helper to reduce the amount of boilercode when using NewUpdateValueSet with multiple datasets.
func NewUpdateValueSets(sets ...[]float64) []*UpdateValueSet {
	new := []*UpdateValueSet{}
	for i := 0; i < len(sets); i++ {
		newSet := NewUpdateValueSet(sets[i])
		new = append(new, newSet)
	}
	return new
}