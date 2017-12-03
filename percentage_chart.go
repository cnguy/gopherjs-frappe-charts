package charts

import (
	"github.com/gopherjs/gopherjs/js"
)

// PercentageChartArgs represents the args needed to instantiate a percentage chart.
// PercentageChart does not have as many configuration arguments as other charts in frappe-charts.
type PercentageChartArgs struct {
	*js.Object
	Parent string     `js:"parent"`
	Title  string     `js:"title"`
	Data   *ChartData `js:"data"`
	Type   string     `js:"type"`
	Height int        `js:"height"`
	Colors []string   `js:"colors"`
}

func NewPercentageChart(parent string, data *ChartData) *PercentageChartArgs {
	new := &PercentageChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Type = "percentage"
	new.Data = data
	return new
}

func (chartArgs *PercentageChartArgs) WithTitle(title string) *PercentageChartArgs {
	chartArgs.Title = title
	return chartArgs
}

func (chartArgs *PercentageChartArgs) WithHeight(height int) *PercentageChartArgs {
	chartArgs.Height = height
	return chartArgs
}

func (chartArgs *PercentageChartArgs) WithColors(colors []string) *PercentageChartArgs {
	chartArgs.Colors = colors
	return chartArgs
}

// Render simply renders a percentage chart without returning anything, because percentage charts don't seem to have any methods yet (in 0.0.3 frappe-charts).
func (chartArgs *PercentageChartArgs) Render() {
	// Set defaults that aren't handled by frappe here.
	if chartArgs.Height == 0 {
		chartArgs.Height = 150
	}
	// Create actual chart.
	js.Global.Get("Chart").New(&chartArgs)
}
