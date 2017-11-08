package charts

import "github.com/gopherjs/gopherjs/js"

// SpecificValue represents a frappe-chart specific value object.
// https://frappe.github.io/charts/
// Ctrl + F "specific_values"
type SpecificValue struct {
	*js.Object
	Title    string  `js:"title"`
	LineType string  `js:"line_type"`
	Value    float64 `js:"value"`
}

// NewSpecificValue is a helper to create a Specific Value JS object.
func NewSpecificValue(title string, lineType string, value float64) *SpecificValue {
	new := &SpecificValue{Object: js.Global.Get("Object").New()}
	new.Title = title
	new.LineType = lineType
	new.Value = value
	return new
}
