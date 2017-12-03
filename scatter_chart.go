package charts

import (
	"github.com/cnguy/gopherjs-frappe-charts/utils"
	"github.com/gopherjs/gopherjs/js"
)

// ScatterChartArgs represents the relevant arguments that are used
// to instantiate a scatter chart.
type ScatterChartArgs struct {
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
	FormatTooltipX func(string) string `js:"format_tooltip_x"`
	FormatTooltipY func(string) string `js:"format_tooltip_y"`
}

func NewScatterChart(parent string, data *ChartData) *ScatterChartArgs {
	new := &ScatterChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Type = "scatter"
	new.Data = data
	return new
}

func (chartArgs *ScatterChartArgs) WithTitle(title string) *ScatterChartArgs {
	chartArgs.Title = title
	return chartArgs
}

func (chartArgs *ScatterChartArgs) WithHeight(height int) *ScatterChartArgs {
	chartArgs.Height = height
	return chartArgs
}

func (chartArgs *ScatterChartArgs) WithColors(colors []string) *ScatterChartArgs {
	chartArgs.Colors = colors
	return chartArgs
}

func (chartArgs *ScatterChartArgs) SetXAxisMode(mode string) *ScatterChartArgs {
	chartArgs.XAxisMode = mode
	return chartArgs
}

func (chartArgs *ScatterChartArgs) SetYAxisMode(mode string) *ScatterChartArgs {
	chartArgs.YAxisMode = mode
	return chartArgs
}

// SetIsSeries allows us to set is_series using a boolean instead of 1/0's.
func (chartArgs *ScatterChartArgs) SetIsSeries(val bool) *ScatterChartArgs {
	chartArgs.SetBton("is_series", val)
	return chartArgs
}

func (chartArgs *ScatterChartArgs) WithFormatTooltipX(callback func(string) string) *ScatterChartArgs {
	chartArgs.FormatTooltipX = callback
	return chartArgs
}

func (chartArgs *ScatterChartArgs) WithFormatTooltipY(callback func(string) string) *ScatterChartArgs {
	chartArgs.FormatTooltipY = callback
	return chartArgs
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
	// Set defaults that aren't handled by frappe.
	if chartArgs.Height == 0 {
		chartArgs.Height = 150
	}
	// Create actual chart.
	new := ScatterChart{js.Global.Get("Chart").New(&chartArgs)}
	return &new
}
