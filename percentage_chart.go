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
}

// NewPercentageChartsArgs is a helper to instantiate a percentage chart.
func NewPercentageChartArgs(parent string, title string, data *ChartData, height int) *PercentageChartArgs {
	new := &PercentageChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Title = title
	new.Data = data
	new.Height = height
	new.Type = "percentage"
	return new
}

// Render simply renders a percentage chart without returning anything, because percentage charts don't seem to have any methods yet (in 0.0.3 frappe-charts).
func (chartArgs *PercentageChartArgs) Render() {
	js.Global.Get("Chart").New(&chartArgs)
}
