package charts

import "github.com/gopherjs/gopherjs/js"

type ChartData struct {
	*js.Object
	Labels         []interface{}    `js:"labels"`
	Datasets       []*Dataset       `js:"datasets"`
	SpecificValues []*SpecificValue `js:"specific_values"`
}

func NewChartData() *ChartData {
	return &ChartData{Object: js.Global.Get("Object").New()}
}
