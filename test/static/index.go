package main

import (
	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	// Make bar chart
	reportCountList := []interface{}{17, 40, 33, 44, 126, 156,
		324, 333, 478, 495, 373}
	barChartData := charts.NewChartData()
	barChartData.Labels = []interface{}{
		"2007", "2008", "2009", "2010", "2011", "2012",
		"2013", "2014", "2015", "2016", "2017",
	}
	barChartData.Datasets = []*charts.Dataset{
		charts.NewDataset("Events", "orange", reportCountList),
	}
	barChartTitle := "Fireball/Bolide Events - Yearly (more than 5 reports"
	barChartArgs := charts.NewBarChartArgs("#chart", barChartTitle, barChartData, 180)
	barChartArgs.SetIsNavigable(true)
	barChartArgs.SetIsSeries(true)
	barChart := barChartArgs.Render()

	// Make line chart
	lineChartValues := []interface{}{36, 46, 45, 32, 27, 31, 30, 36, 39, 49, 0, 0}
	lineChartData := charts.NewChartData()
	lineChartData.Labels = []interface{}{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	lineChartData.Datasets = []*charts.Dataset{
		charts.NewDataset("", "green", lineChartValues),
	}
	lineChartArgs := charts.NewLineChartArgs("#chart-2", "", lineChartData, 180)
	lineChartArgs.SetIsSeries(true)
	lineChart := lineChartArgs.Render()

	// Prepare update values for event listener
	moreLineData := []*charts.UpdateValueSet{
		charts.NewUpdateValueSet([]interface{}{4, 0, 3, 1, 1, 2, 1, 2, 1, 0, 1, 1}),
		charts.NewUpdateValueSet([]interface{}{2, 3, 3, 2, 1, 4, 0, 1, 2, 7, 11, 4}),
		charts.NewUpdateValueSet([]interface{}{7, 7, 2, 4, 0, 1, 5, 3, 1, 2, 0, 1}),
		charts.NewUpdateValueSet([]interface{}{0, 2, 6, 2, 2, 1, 2, 3, 6, 3, 7, 10}),
		charts.NewUpdateValueSet([]interface{}{9, 10, 8, 10, 6, 5, 8, 8, 24, 15, 10, 13}),
		charts.NewUpdateValueSet([]interface{}{9, 13, 16, 9, 4, 5, 7, 10, 14, 22, 23, 24}),
		charts.NewUpdateValueSet([]interface{}{20, 22, 28, 19, 28, 19, 14, 19, 51, 37, 29, 38}),
		charts.NewUpdateValueSet([]interface{}{29, 20, 22, 16, 16, 19, 24, 26, 57, 31, 46, 2}),
		charts.NewUpdateValueSet([]interface{}{36, 24, 38, 27, 15, 22, 24, 38, 32, 57, 139, 26}),
		charts.NewUpdateValueSet([]interface{}{37, 36, 32, 33, 12, 34, 52, 45, 58, 57, 64, 35}),
		charts.NewUpdateValueSet([]interface{}{36, 46, 45, 32, 27, 31, 30, 36, 39, 49, 0, 0}),
	}

	barChart.AddDataSelectListener(func(event *charts.DataSelectEvent) {
		updateValues := &charts.UpdateValuesArgs{
			Values: []*charts.UpdateValueSet{moreLineData[event.Index]},
			// keep labels same as before
			Labels: lineChartData.Labels,
		}
		lineChart.UpdateValues(updateValues)
	})
}
