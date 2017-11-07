package charts

import (
	"github.com/cnguy/gopherjs-frappe-charts/utils"
	"github.com/gopherjs/gopherjs/js"
)

type ScatterChartArgs struct {
	*js.Object
	Parent    string     `js:"parent"`
	Title     string     `js:"title"`
	Data      *ChartData `js:"data"`
	Type      string     `js:"type"`
	Height    int        `js:"height"`
	XAxisMode string     `js:"x_axis_mode"`
	YAxisMode string     `js:"y_axis_mode"`
	IsSeries  int        `js:"is_series"`
}

func NewScatterChartArgs(parent string, title string, data *ChartData, height int) *ScatterChartArgs {
	new := &ScatterChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Title = title
	new.Type = "line"
	new.Data = data
	new.Height = height
	return new
}

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

func (chart *ScatterChart) PopDataPoint() {
	chart.Call("remove_data_point")
}

func (chart *ScatterChart) RemoveDataPoint(index int) {
	chart.Call("remove_data_point", index)
}

func (chart *ScatterChart) ShowAverages() {
	chart.Call("show_averages")
}

func (chart *ScatterChart) HideAverages() {
	chart.Call("hide_averages")
}

func (chartArgs *ScatterChartArgs) Render() *ScatterChart {
	new := ScatterChart{js.Global.Get("Chart").New(&chartArgs)}
	return &new
}
