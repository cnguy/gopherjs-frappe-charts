package charts

import (
	"github.com/cnguy/gopherjs-frappe-charts/utils"
	"github.com/gopherjs/gopherjs/js"
)

// LineChartArgs represents the configuration arguments that can be used to instantiate a line chart.
type LineChartArgs struct {
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
	ShowDots       int                 `js:"show_dots"`
	Heatline       int                 `js:"heatline"`
	RegionFill     int                 `js:"region_fill"`
	FormatTooltipX func(string) string `js:"format_tooltip_x"`
	FormatTooltipY func(string) string `js:"format_tooltip_y"`
}

func NewLineChart(parent string, data *ChartData) *LineChartArgs {
	new := &LineChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Type = "line"
	new.Data = data
	return new
}

func (chartArgs *LineChartArgs) WithTitle(title string) *LineChartArgs {
	chartArgs.Title = title
	return chartArgs
}

func (chartArgs *LineChartArgs) WithHeight(height int) *LineChartArgs {
	chartArgs.Height = height
	return chartArgs
}

func (chartArgs *LineChartArgs) WithColors(colors []string) *LineChartArgs {
	chartArgs.Colors = colors
	return chartArgs
}

func (chartArgs *LineChartArgs) SetXAxisMode(mode string) *LineChartArgs {
	chartArgs.XAxisMode = mode
	return chartArgs
}

func (chartArgs *LineChartArgs) SetYAxisMode(mode string) *LineChartArgs {
	chartArgs.YAxisMode = mode
	return chartArgs
}

// SetIsSeries allows us to set is_series using a boolean instead of 1/0's.
func (chartArgs *LineChartArgs) SetIsSeries(val bool) *LineChartArgs {
	chartArgs.SetBton("is_series", val)
	return chartArgs
}

// SetShowDots allows us to set show_dots using a boolean instead of 1/0's.
func (chartArgs *LineChartArgs) SetShowDots(val bool) *LineChartArgs {
	chartArgs.SetBton("show_dots", val)
	return chartArgs
}

// SetHeatline allows us to set heatline using a boolean instead of 1/0's.
func (chartArgs *LineChartArgs) SetHeatline(val bool) *LineChartArgs {
	chartArgs.SetBton("heatline", val)
	return chartArgs
}

// SetRegionFill allows us to set region_fill using a boolean instead of 1/0's.
func (chartArgs *LineChartArgs) SetRegionFill(val bool) *LineChartArgs {
	chartArgs.SetBton("region_fill", val)
	return chartArgs
}

func (chartArgs *LineChartArgs) SetBton(key string, val bool) {
	chartArgs.Set(key, utils.Btoi(val))
}

func (chartArgs *LineChartArgs) WithFormatTooltipX(callback func(string) string) *LineChartArgs {
	chartArgs.FormatTooltipX = callback
	return chartArgs
}

func (chartArgs *LineChartArgs) WithFormatTooltipY(callback func(string) string) *LineChartArgs {
	chartArgs.FormatTooltipY = callback
	return chartArgs
}

type LineChart struct {
	*js.Object
}

// UpdateValues allows us to batch update the values of the chart.
// Make sure to pass in the original labels if you want to keep the structure of your chart.
// This should be improved.
func (chart *LineChart) UpdateValues(updateValues *UpdateValuesArgs) {
	chart.Call("update_values", updateValues.Values, updateValues.Labels)
}

// AppendDataPoint calls add_data_point without an index (which defaults to the end).
func (chart *LineChart) AppendDataPoint(values []float64, label string) {
	chart.Call("add_data_point", utils.FloatSliceToInterface(values), label)
}

// AddDataPoint allows us to add a data point to anywhere in the chart.
func (chart *LineChart) AddDataPoint(values []float64, label string, index int) {
	chart.Call("add_data_point", utils.FloatSliceToInterface(values), label, index)
}

// PopDataPoint calls remove_data_point without arguments, which defaults to removing the element at index 0.
func (chart *LineChart) PopDataPoint() {
	chart.Call("remove_data_point")
}

// RemoveDataPoint allows users to remove a data point anywhere in the array.
func (chart *LineChart) RemoveDataPoint(index int) {
	chart.Call("remove_data_point", index)
}

// ShowAverages calls show_averages which makes a new horizontal line representing the average for every dataset.
func (chart *LineChart) ShowAverages() {
	chart.Call("show_averages")
}

// HideAverages simply hides the averages displayed using ShowAverages.
func (chart *LineChart) HideAverages() {
	chart.Call("hide_averages")
}

// Render returns a LineChart that allows users to call the above functions.
func (chartArgs *LineChartArgs) Render() *LineChart {
	// Set defaults that aren't handled by frappe here.
	if chartArgs.Height == 0 {
		chartArgs.Height = 150
	}
	// Create actual chart.
	new := LineChart{js.Global.Get("Chart").New(&chartArgs)}
	return &new
}
