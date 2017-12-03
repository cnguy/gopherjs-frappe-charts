package charts

import (
	"github.com/cnguy/gopherjs-frappe-charts/utils"
	"github.com/gopherjs/gopherjs/js"
)

type BarChartArgs struct {
	*js.Object
	Parent         string              `js:"parent"`
	Title          string              `js:"title"`
	Data           *ChartData          `js:"data"`
	Type           string              `js:"type"`
	Height         int                 `js:"height"`
	Colors         []string            `js:"colors"`
	XAxisMode      string              `js:"x_axis_mode"`
	YAxisMode      string              `js:"y_axis_mode"`
	IsSeries       int                 `js:"is_series"`
	IsNavigable    int                 `js:"is_navigable"`
	FormatTooltipX func(string) string `js:"format_tooltip_x"`
	FormatTooltipY func(string) string `js:"format_tooltip_y"`
}

func NewBarChart(parent string, data *ChartData) *BarChartArgs {
	new := &BarChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Type = "bar"
	new.Data = data
	return new
}

func (chartArgs *BarChartArgs) WithTitle(title string) *BarChartArgs {
	chartArgs.Title = title
	return chartArgs
}

func (chartArgs *BarChartArgs) WithHeight(height int) *BarChartArgs {
	chartArgs.Height = height
	return chartArgs
}

func (chartArgs *BarChartArgs) WithColors(colors []string) *BarChartArgs {
	chartArgs.Colors = colors
	return chartArgs
}

func (chartArgs *BarChartArgs) SetXAxisMode(mode string) *BarChartArgs {
	chartArgs.XAxisMode = mode
	return chartArgs
}

func (chartArgs *BarChartArgs) SetYAxisMode(mode string) *BarChartArgs {
	chartArgs.YAxisMode = mode
	return chartArgs
}

func (chartArgs *BarChartArgs) SetIsSeries(val bool) *BarChartArgs {
	chartArgs.SetBton("is_series", val)
	return chartArgs
}

func (chartArgs *BarChartArgs) SetIsNavigable(val bool) *BarChartArgs {
	chartArgs.SetBton("is_navigable", val)
	return chartArgs
}

func (chartArgs *BarChartArgs) WithFormatTooltipX(callback func(string) string) *BarChartArgs {
	chartArgs.FormatTooltipX = callback
	return chartArgs
}

func (chartArgs *BarChartArgs) WithFormatTooltipY(callback func(string) string) *BarChartArgs {
	chartArgs.FormatTooltipY = callback
	return chartArgs
}

func (chartArgs *BarChartArgs) SetBton(key string, val bool) {
	chartArgs.Set(key, utils.Btoi(val))
}

type BarChart struct {
	*js.Object
}

func (chart *BarChart) UpdateValues(updateValues *UpdateValuesArgs) {
	chart.Call("update_values", updateValues.Values, updateValues.Labels)
}

func (chart *BarChart) AppendDataPoint(values []interface{}, label string) {
	chart.Call("add_data_point", values, label)
}

func (chart *BarChart) AddDataPoint(values []interface{}, label string, index int) {
	chart.Call("add_data_point", values, label, index)
}

func (chart *BarChart) PopDataPoint() {
	chart.Call("remove_data_point")
}

func (chart *BarChart) RemoveDataPoint(index int) {
	chart.Call("remove_data_point", index)
}

func (chart *BarChart) ShowSums() {
	chart.Call("show_sums")
}

func (chart *BarChart) HideSums() {
	chart.Call("hide_sums")
}

func (chart *BarChart) ShowAverages() {
	chart.Call("show_averages")
}

func (chart *BarChart) HideAverages() {
	chart.Call("hide_averages")
}

func (chartArgs *BarChartArgs) Render() *BarChart {
	// Set defaults that aren't handled by frappe here.
	if chartArgs.Height == 0 {
		chartArgs.Height = 150
	}
	// Create actual chart.
	new := BarChart{js.Global.Get("Chart").New(&chartArgs)}
	return &new
}

type DataSelectEvent struct {
	*js.Object
	Value interface{} `js:"value"`
	Label interface{} `js:"label"`
	Index int         `js:"index"`
}

type eventListenerCallback func(event *DataSelectEvent)

func (chart *BarChart) AddEventListener(eventName string, fn eventListenerCallback) {
	parent := chart.Get("parent").Get("id").String()
	element := js.Global.Get("document").Call("getElementById", parent)
	element.Call("addEventListener", eventName, fn)
}

// AddDataSelectListener is a helper function to allow users to quickly set a data-select event.
// This allows users to access the index, value, and label of any particular data point!
func (chart *BarChart) AddDataSelectListener(fn eventListenerCallback) {
	chart.AddEventListener("data-select", fn)
}
