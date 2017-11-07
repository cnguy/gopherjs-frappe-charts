package charts

import (
	"github.com/gopherjs/gopherjs/js"
)

type PercentageChartArgs struct {
	*js.Object
	Parent string     `js:"parent"`
	Title  string     `js:"title"`
	Data   *ChartData `js:"data"`
	Type   string     `js:"type"`
	Height int        `js:"height"`
}

func NewPercentageChartArgs(parent string, title string, data *ChartData, height int) *PercentageChartArgs {
	new := &PercentageChartArgs{Object: js.Global.Get("Object").New()}
	new.Parent = parent
	new.Title = title
	new.Data = data
	new.Height = height
	new.Type = "percentage"
	return new
}

func (chartArgs *PercentageChartArgs) Render() {
	js.Global.Get("Chart").New(&chartArgs)
}
