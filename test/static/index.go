package main

import (
	"time"

	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	// Prepare data
	chartData := charts.NewChartData()
	chartData.Labels = []interface{}{
		"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat",
	}
	chartData.Datasets = []*charts.Dataset{
		charts.NewDataset("", "purple", []interface{}{25, 40, 30, 35, 8, 52, 17}),
		charts.NewDataset("", "orange", []interface{}{25, 50, -10, 15, 18, 32, 27}),
	}

	// Prepare constructor arguments
	chartArgs := charts.NewChartArgs("#chart", "", "bar", chartData, 250)

	chart := chartArgs.Render()

	time.Sleep(1 * time.Second)
	chart.ShowSums()

	time.Sleep(1 * time.Second)
	chart.HideSums()

	time.Sleep(1 * time.Second)
	chart.ShowAverages()
	println(chart)
	time.Sleep(1 * time.Second)
	chart.HideAverages()

	println(chart)
}
