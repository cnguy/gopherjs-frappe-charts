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
	Colors []string   `js:"colors"`
}

func NewPieChart(parent string, data *ChartData) *PieChartArgs {
	new := &PieChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Type = "pie"
	new.Data = data
	return new
}

func (chartArgs *PieChartArgs) WithTitle(title string) *PieChartArgs {
	chartArgs.Title = title
	return chartArgs
}

func (chartArgs *PieChartArgs) WithHeight(height int) *PieChartArgs {
	chartArgs.Height = height
	return chartArgs
}

func (chartArgs *PieChartArgs) WithColors(colors []string) *PieChartArgs {
	chartArgs.Colors = colors
	return chartArgs
}

// Render simply renders a pie chart without returning anything, because pie charts don't seem to have any methods yet (in 0.0.3 frappe-charts).
func (chartArgs *PieChartArgs) Render() {
	// Set defaults that aren't handled by frappe here.
	if chartArgs.Height == 0 {
		chartArgs.Height = 150
	}
	// Create actual chart.
	js.Global.Get("Chart").New(&chartArgs)
}
