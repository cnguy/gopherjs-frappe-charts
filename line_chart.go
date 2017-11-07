package charts

import (
	"github.com/cnguy/gopherjs-frappe-charts/utils"
	"github.com/gopherjs/gopherjs/js"
)

type LineChartArgs struct {
	*js.Object
	Parent     string     `js:"parent"`
	Title      string     `js:"title"`
	Data       *ChartData `js:"data"`
	Type       string     `js:"type"`
	Height     int        `js:"height"`
	XAxisMode  string     `js:"x_axis_mode"`
	YAxisMode  string     `js:"y_axis_mode"`
	IsSeries   int        `js:"is_series"`
	ShowDots   int        `js:"show_dots"`
	Heatline   int        `js:"heatline"`
	RegionFill int        `js:"region_fill"`
}

func NewLineChartArgs(parent string, title string, data *ChartData, height int) *LineChartArgs {
	new := &LineChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Title = title
	new.Type = "line"
	new.Data = data
	new.Height = height
	return new
}

func (chartArgs *LineChartArgs) SetIsSeries(val bool) {
	chartArgs.SetBton("is_series", val)
}

func (chartArgs *LineChartArgs) SetShowDots(val bool) {
	chartArgs.SetBton("show_dots", val)
}

func (chartArgs *LineChartArgs) SetHeatline(val bool) {
	chartArgs.SetBton("heatline", val)
}

func (chartArgs *LineChartArgs) SetRegionFill(val bool) {
	chartArgs.SetBton("region_fill", val)
}

func (chartArgs *LineChartArgs) SetBton(key string, val bool) {
	chartArgs.Set(key, utils.Btoi(val))
}

type LineChart struct {
	*js.Object
}

func (chart *LineChart) UpdateValues(updateValues *UpdateValuesArgs) {
	chart.Call("update_values", updateValues.Values, updateValues.Labels)
}

func (chart *LineChart) AppendDataPoint(values []interface{}, label string) {
	chart.Call("add_data_point", values, label)
}

func (chart *LineChart) AddDataPoint(values []interface{}, label string, index int) {
	chart.Call("add_data_point", values, label, index)
}

func (chart *LineChart) PopDataPoint() {
	chart.Call("remove_data_point")
}

func (chart *LineChart) RemoveDataPoint(index int) {
	chart.Call("remove_data_point", index)
}

func (chart *LineChart) ShowAverages() {
	chart.Call("show_averages")
}

func (chart *LineChart) HideAverages() {
	chart.Call("hide_averages")
}

func (chartArgs *LineChartArgs) Render() *LineChart {
	new := LineChart{js.Global.Get("Chart").New(&chartArgs)}
	return &new
}
