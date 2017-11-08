package charts

import "github.com/gopherjs/gopherjs/js"

// ChartData represents the "data" object that is passed into any of the charts
// except Heatmap.
// https://frappe.github.io/charts/
// Ctrl + F "let data ="
type ChartData struct {
	*js.Object
	Labels         []interface{}    `js:"labels"`
	Datasets       []*Dataset       `js:"datasets"`
	SpecificValues []*SpecificValue `js:"specific_values"`
}

// NewChartData just instantiates a ChartData JS object.
func NewChartData() *ChartData {
	return &ChartData{Object: js.Global.Get("Object").New()}
}
