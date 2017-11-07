# GopherJS bindings for [frappe/charts](https://github.com/frappe/charts) 

[![Go Report Card](https://goreportcard.com/badge/github.com/cnguy/gopherjs-frappe-charts)](https://goreportcard.com/report/github.com/cnguy/gopherjs-frappe-charts)
[![GoDoc](https://godoc.org/github.com/cnguy/gopherjs-frappe-charts?status.svg)](https://godoc.org/github.com/cnguy/gopherjs-frappe-charts)

`Simple, responsive, modern SVG Charts with zero dependencies`

![Alt text](/_pictures/event_listener.gif?raw=true "Event Listener")

## Contents
* [Installation](#installation)
* [Usage](#usage)
* [Disclaimer](#disclaimer)
* [Contributions](#contributions)
* [Changelog](#changelog)

### Installation

*Built on go 1.9*

`go get -u github.com/cnguy/gopherjs-frappe-charts`

Make sure you also have the [necessary JS files](https://github.com/frappe/charts#installation).

Currently, this API supports frappe-chart version 0.0.3.

Dependencies will be managed via [dep](https://github.com/golang/dep) in the future.

Unfortunately, `gopherjs` is not [vendorable](https://github.com/gopherjs/gopherjs/issues/415) at the moment.

### Usage

Following examples assume an HTML file like so:

```html
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>GopherJS bindings for frappe-charts</title>
</head>
<body>
	<div id="chart"></div>
	<script src="https://unpkg.com/frappe-charts@0.0.3/dist/frappe-charts.min.iife.js"></script>
	<script src="static.js" data-cover></script>
</body>
</html>
```

where `static.js` is the name of your bundled JS file when your folder is named `static` and you run `gopherjs build`.

The basic development flow of this library is:

1) Chart data: Declare your data via NewDataset() and labels
2) Chart args: Declare the basic core arguments via NewChartArgs(), and then set extra config if needed (heatline, is_navigable)
3) Call Render() on the chart args, which returns a reference to a Chart object that you can call various methods on (ShowSums, ShowAverages)

* [Hello World](#hello-world)
* [Navigable](#navigable)
* [Line Properties](#line-properties)
* [Simple Aggregations](#simple-aggregations)
* [Event Listener](#event-listener)

#### Hello World:

```go
package main

import (
    charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	// Prepare data
	chartData := charts.NewChartData()
	chartData.Labels = []interface{}{
		"12am-3am", "3am-6pm", "6am-9am", "9am-12am",
		"12pm-3pm", "3pm-6pm", "6pm-9pm", "9am-12am",
	}
	chartData.Datasets = []*charts.Dataset{
		charts.NewDataset(
			"Some Data",
			"light-blue",
			[]interface{}{25, 40, 30, 35, 8, 52, 17, -4},
		),
		charts.NewDataset(
			"Another Set",
			"violet",
			[]interface{}{25, 50, -10, 15, 18, 32, 27, 14},
		),
	}
	
	// Prepare constructor arguments via the NewChartArgs helper
	chartArgs := charts.NewChartArgs("#chart", "My Awesome Chart", "bar", chartData, 250)
	// chartArgs.Parent = "#chart"
	// chartArgs.Title = "My Awesome Chart"
	// chartArgs.IsNavigable = ...
	// chartArgs.Heatline = ...

	chart := chartArgs.Render()
	println(chart) // to suppress the annoying error
}
```

Output 1:

![Alt text](/_pictures/hello_world.png?raw=true "Hello World")

#### Navigable:

```go
package main

import (
	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	// Prepare data
	chartData := charts.NewChartData()
	chartData.Labels = []interface{}{
		"2007", "2008", "2009", "2010", "2011", "2012", "2013", "2014", "2015", "2016", "2017",
	}
	values := []interface{}{17, 40, 33, 44, 126, 156, 324, 333, 478, 495, 373}
	dataset := charts.NewDataset("Events", "orange", values)
	chartData.Datasets = []*charts.Dataset{dataset}

	// Prepare constructor arguments
	chartTitle := "Fireball/Bolide Events - Yearly (more than 5 reports"
	chartArgs := charts.NewChartArgs("#chart", chartTitle, "bar", chartData, 180)
	chartArgs.SetIsNavigable(true) // chartArgs.IsNavigable = 1
	chartArgs.SetIsSeries(true) // chartArgs.IsSeries = 1

	chart := chartArgs.Render()
}
```

Output 2:

*Using left/right arrow keys to traverse the graph*

![Alt text](/_pictures/navigable.gif?raw=true "Navigable")

#### Line Properties:

```go
package main

import (
	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	// Prepare data
	chartData := charts.NewChartData()
	chartData.Labels = []interface{}{
		1967, 1968, 1969, 1970, 1971, 1972, 1973, 1974, 1975, 1976,
		1977, 1978, 1979, 1980, 1981, 1982, 1983, 1984, 1985, 1986,
		1987, 1988, 1989, 1990, 1991, 1992, 1993, 1994, 1995, 1996,
		1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006,
		2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016,
	}
	values := []interface{}{
		132.9, 150.0, 149.4, 148.0, 94.4, 97.6, 54.1, 49.2, 22.5, 18.4,
		39.3, 131.0, 220.1, 218.9, 198.9, 162.4, 91.0, 60.5, 20.6, 14.8,
		33.9, 123.0, 211.1, 191.8, 203.3, 133.0, 76.1, 44.9, 25.1, 11.6,
		28.9, 88.3, 136.3, 173.9, 170.4, 163.6, 99.3, 65.3, 45.8, 24.7,
		12.6, 4.2, 4.8, 24.9, 80.8, 84.5, 94.0, 113.3, 69.8, 39.8,
	}
	dataset := charts.NewDataset("", "blue", values)
	chartData.Datasets = []*charts.Dataset{dataset}

	// Prepare constructor arguments
	chartTitle := "Mean Total Sunspot Count - Yearly"
	chartArgs := charts.NewChartArgs("#chart", chartTitle, "line", chartData, 180)
	chartArgs.SetShowDots(false)
	chartArgs.SetHeatline(true)
	chartArgs.SetRegionFill(true)
	chartArgs.XAxisMode = "tick"
	chartArgs.YAxisMode = "span"
	chartArgs.SetIsSeries(true)

	chart := chartArgs.Render()
}
```

Output 3:

![Alt text](/_pictures/line_properties.gif?raw=true "Navigable")

#### Simple Aggregations:

```go
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

	time.Sleep(1 * time.Second)
	chart.HideAverages()
}
```

Output 4:

![Alt text](/_pictures/simple_aggregations.gif?raw=true "Simple Aggregations")

#### Event Listener:

```go
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
	barChartArgs := charts.NewChartArgs("#chart", barChartTitle, "bar", barChartData, 180)
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
	lineChartArgs := charts.NewChartArgs("#chart-2", "", "line", lineChartData, 180)
	lineChartArgs.SetIsSeries(true)
	lineChart := lineChartArgs.Render()

	// Prepare update values for event listener
	moreLineData := make(map[int]*charts.UpdateValueSet)
	moreLineData[0] = charts.NewUpdateValueSet([]interface{}{4, 0, 3, 1, 1, 2, 1, 2, 1, 0, 1, 1})
	moreLineData[1] = charts.NewUpdateValueSet([]interface{}{2, 3, 3, 2, 1, 4, 0, 1, 2, 7, 11, 4})
	moreLineData[2] = charts.NewUpdateValueSet([]interface{}{7, 7, 2, 4, 0, 1, 5, 3, 1, 2, 0, 1})
	moreLineData[3] = charts.NewUpdateValueSet([]interface{}{0, 2, 6, 2, 2, 1, 2, 3, 6, 3, 7, 10})
	moreLineData[4] = charts.NewUpdateValueSet([]interface{}{9, 10, 8, 10, 6, 5, 8, 8, 24, 15, 10, 13})
	moreLineData[5] = charts.NewUpdateValueSet([]interface{}{9, 13, 16, 9, 4, 5, 7, 10, 14, 22, 23, 24})
	moreLineData[6] = charts.NewUpdateValueSet([]interface{}{20, 22, 28, 19, 28, 19, 14, 19, 51, 37, 29, 38})
	moreLineData[7] = charts.NewUpdateValueSet([]interface{}{29, 20, 22, 16, 16, 19, 24, 26, 57, 31, 46, 2})
	moreLineData[8] = charts.NewUpdateValueSet([]interface{}{36, 24, 38, 27, 15, 22, 24, 38, 32, 57, 139, 26})
	moreLineData[9] = charts.NewUpdateValueSet([]interface{}{37, 36, 32, 33, 12, 34, 52, 45, 58, 57, 64, 35})
	moreLineData[10] = charts.NewUpdateValueSet([]interface{}{36, 46, 45, 32, 27, 31, 30, 36, 39, 49, 0, 0})

	barChart.AddDataSelectListener(func(event *charts.DataSelectEvent) {
		updateValues := &charts.UpdateValuesArgs{}
		updateValues.Values = []*charts.UpdateValueSet{moreLineData[event.Index]}
		// Labels must be set to the old labels or else we pass in nothing
		// and frappe-charts will explode.
		updateValues.Labels = lineChartData.Labels
		lineChart.UpdateValues(updateValues)
	})
}
```

Output 5:

![Alt text](/_pictures/event_listener.gif?raw=true "Event Listener")

### Disclaimer

I'm still a big Go noob, and so I don't know the proper way to manage releases and dependencies yet.
I'll try to figure it out soon!

### Contributions

Any contributions are welcome. :)

### Changelog

[here](https://github.com/cnguy/gopherjs-frappe-charts/blob/master/CHANGELOG.md) when available