# GopherJS bindings for the recently popular [frappe/charts](https://github.com/frappe/charts) library

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

Dependencies are managed via [dep](https://github.com/golang/dep).

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

Hello world:

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
    
        // Prepare constructor arguments
        chartArgs := charts.NewChartArgs("#chart", "My Awesome Chart", "bar", chartData, 250)
        // You can also manually set the arguments. The helper above
        // is simply to get the most important details set.
        // chartArgs.Parent = ...
        // chartArgs.IsNavigable = ...
        // chartArgs.Heatline = ...

	chart := chartArgs.Render()
	println(chart) // to suppress the annoying error
}
```

Output 1:

![Alt text](/_pictures/hello_world.png?raw=true "Hello World")

Navigable:

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

Line Properties:

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

Simple Aggregations:

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

### Disclaimer

I'm still a big Go noob, and so I don't know the proper way to manage releases and dependencies yet.
I'll try to figure it out soon!

### Contributions

Any contributions are welcome. :)

### Changelog

[here](https://github.com/cnguy/gopherjs-frappe-charts/blob/master/CHANGELOG.md) when available