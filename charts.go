package charts

import (
	"github.com/cnguy/gopherjs-frappe-charts/utils"
	"github.com/gopherjs/gopherjs/js"
)

// ChartArgs is just a bag of arguments used to instantiate a frappe-chart.
type ChartArgs struct {
	*js.Object
	Parent      string     `js:"parent"`
	Title       string     `js:"title"`
	Data        *ChartData `js:"data"`
	Type        string     `js:"type"`
	Height      int        `js:"height"`
	XAxisMode   string     `js:"x_axis_mode"`
	YAxisMode   string     `js:"y_axis_mode"`
	IsNavigable int        `js:"is_navigable"`
	IsSeries    int        `js:"is_series"`
	ShowDots    int        `js:"show_dots"`
	Heatline    int        `js:"heatline"`
	RegionFill  int        `js:"region_fill"`
}

// ChartData is a set of labels and their datasets.
type ChartData struct {
	*js.Object
	Labels         []interface{}    `js:"labels"`
	Datasets       []*Dataset       `js:"datasets"`
	SpecificValues []*SpecificValue `js:"specific_values"`
}

// Dataset if a set of data to be passed into ChartData.
type Dataset struct {
	*js.Object
	Title  string        `js:"title"`
	Color  string        `js:"color"`
	Values []interface{} `js:"values"`
}

// Render is a ChartArgs function that returns a Chart object.
func (chartArgs *ChartArgs) Render() *Chart {
	new := Chart{js.Global.Get("Chart").New(&chartArgs)}
	return &new
}

// NewChartData is a helper function to return a reference to an instantiated ChartData.
func NewChartData() *ChartData {
	return &ChartData{Object: js.Global.Get("Object").New()}
}

// NewChartArgs is a helper function to return a reference to an instantiated ChartArgs that has all mandatory fields filled out.
func NewChartArgs(parent string, title string, chartType string, data *ChartData, height int) *ChartArgs {
	new := &ChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Title = title
	new.Type = chartType
	new.Data = data
	new.Height = height
	return new
}

// SetIsNavigable is a helper function that sets is_navigable using booleans instead of numbers.
func (chartArgs *ChartArgs) SetIsNavigable(val bool) {
	chartArgs.setBoolToNum("is_navigable", val)
}

// SetIsSeries is a helper function that sets is_series using booleans instead of numbers.
func (chartArgs *ChartArgs) SetIsSeries(val bool) {
	chartArgs.setBoolToNum("is_series", val)
}

// SetShowDots is a helper function that sets show_dots using booleans instead of numbers.
func (chartArgs *ChartArgs) SetShowDots(val bool) {
	chartArgs.setBoolToNum("show_dots", val)
}

// SetHeatline is a helper function that sets heatline using booleans instead of numbers.
func (chartArgs *ChartArgs) SetHeatline(val bool) {
	chartArgs.setBoolToNum("heatline", val)
}

// SetRegionFill is a helper function that sets region_fill using booleans instead of numbers.
func (chartArgs *ChartArgs) SetRegionFill(val bool) {
	chartArgs.setBoolToNum("region_fill", val)
}

// SetBoolToNum is a helper function that abstracts away the use of 1's/0's instead of true/false.
func (chartArgs *ChartArgs) setBoolToNum(key string, val bool) {
	chartArgs.Set(key, utils.Btoi(val))
}

// NewDataset returns a reference to a newly-instantiated dataset.
func NewDataset(title string, color string, values []interface{}) (dataset *Dataset) {
	dataset = &Dataset{Object: js.Global.Get("Object").New()}
	dataset.Title = title
	dataset.Color = color
	dataset.Values = values
	return dataset
}

// Chart is a reference to an object instantiated from new Chart(chart arguments).
// The reason that this struct is empty is that it seems that a lot of the properties
// are not really intended to be changed manually (through dot notation).
// Instead, the developers expect the endusers to rely on the public API's
// such as update_values(), show_sums(), remove_data_point(), etc.
type Chart struct {
	*js.Object
}

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
