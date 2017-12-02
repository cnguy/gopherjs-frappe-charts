package main

import (
	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	// Prepare data
	chartData := charts.NewChartData()
	chartData.Labels = []string{
		"12am-3am", "3am-6pm", "6am-9am", "9am-12am",
		"12pm-3pm", "3pm-6pm", "6pm-9pm", "9am-12am",
	}
	chartData.Datasets = []*charts.Dataset{
		charts.NewDataset(
			"Some Data",
			[]interface{}{25, 40, 30, 35, 8, 52, 17, -4},
		),
		charts.NewDataset(
			"Another Set",
			[]interface{}{25, 50, -10, 15, 18, 32, 27, 14},
		),
	}

	// Prepare constructor arguments via the NewChartArgs helper
	chartArgs := charts.NewScatterChartArgs("#chart", "My Awesome Chart", chartData, 250)
	chartArgs.Colors = []string{"light-blue", "violet"}
	chartArgs.FormatTooltipX = func(d string) string { return d + "...." }
	chartArgs.FormatTooltipY = func(d string) string { return "..." + d }
	// chartArgs.Parent = "#chart"
	// chartArgs.Title = "My Awesome Chart"
	// chartArgs.IsNavigable = ...
	// chartArgs.Heatline = ...

	chartArgs.Render()
}
