package charts

import (
	"github.com/cnguy/gopherjs-frappe-charts/utils"
	"github.com/gopherjs/gopherjs/js"
)

// ScatterChartArgs represents the relevant arguments that are used
// to instantiate a scatter chart.
type ScatterChartArgs struct {
	*js.Object
	Parent    string     `js:"parent"`
	Title     string     `js:"title"`
	Data      *ChartData `js:"data"`
	Type      string     `js:"type"`
	Height    int        `js:"height"`
	Colors    []string   `js:"colors"`
	XAxisMode string     `js:"x_axis_mode"`
	YAxisMode string     `js:"y_axis_mode"`
	IsSeries  int        `js:"is_series"`
}

// NewScatterChartArgs is a helper that creates a new ScatterChatArgs with important
// fields filled out.
func NewScatterChartArgs(parent string, title string, data *ChartData, height int) *ScatterChartArgs {
	new := &ScatterChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Title = title
	new.Type = "scatter"
	new.Data = data
	new.Height = height
	return new
}

// SetIsSeries is a small helper that just allows users to work with booleans
// instead of 0's and 1's. In this case, users can set the JS property "is_series" to 1
// by calling chartArgs.SetIsSeries(true).
func (chartArgs *ScatterChartArgs) SetIsSeries(val bool) {
	chartArgs.SetBton("is_series", val)
}

func (chartArgs *ScatterChartArgs) SetBton(key string, val bool) {
	chartArgs.Set(key, utils.Btoi(val))
}

type ScatterChart struct {
	*js.Object
}

func (chart *ScatterChart) UpdateValues(updateValues *UpdateValuesArgs) {
	chart.Call("update_values", updateValues.Values, updateValues.Labels)
}

func (chart *ScatterChart) AppendDataPoint(values []interface{}, label string) {
	chart.Call("add_data_point", values, label)
}

func (chart *ScatterChart) AddDataPoint(values interface{}, label string, index int) {
	chart.Call("add_data_point", values, label, index)
}

// PopDataPoint calls remove_data_point without arguments, which defaults to removing the element at index 0.
func (chart *ScatterChart) PopDataPoint() {
	chart.Call("remove_data_point")
}

// RemoveDataPoint allows users to remove a data point anywhere in the array.
func (chart *ScatterChart) RemoveDataPoint(index int) {
	chart.Call("remove_data_point", index)
}

// ShowAverages calls show_averages which makes a new horizontal line representing the average for every dataset.
func (chart *ScatterChart) ShowAverages() {
	chart.Call("show_averages")
}

// HideAverages simply hides the averages displayed using ShowAverages.
func (chart *ScatterChart) HideAverages() {
	chart.Call("hide_averages")
}

// Render returns a ScatterChart that allows users to call the above functions.
func (chartArgs *ScatterChartArgs) Render() *ScatterChart {
	new := ScatterChart{js.Global.Get("Chart").New(&chartArgs)}
	return &new
}
