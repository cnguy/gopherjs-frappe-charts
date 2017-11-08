package charts

import (
	"github.com/gopherjs/gopherjs/js"
)

// PieChartArgs represents the args needed to instantiate a pie chart.
// PieChart does not have as many configuration arguments as other charts in frappe-charts.
type PieChartArgs struct {
	*js.Object
	Parent string     `js:"parent"`
	Title  string     `js:"title"`
	Data   *ChartData `js:"data"`
	Type   string     `js:"type"`
	Height int        `js:"height"`
}

// NewPieChartsArgs is a helper to instantiate a pie chart.
func NewPieChartArgs(parent string, title string, data *ChartData, height int) *PieChartArgs {
	new := &PieChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Title = title
	new.Data = data
	new.Height = height
	new.Type = "pie"
	return new
}

// Render simply renders a pie chart without returning anything, because pie charts don't seem to have any methods yet (in 0.0.3 frappe-charts).
func (chartArgs *PieChartArgs) Render() {
	js.Global.Get("Chart").New(&chartArgs)
}
