# GopherJS bindings for [frappe/charts](https://github.com/frappe/charts) 

[![Go Report Card](https://goreportcard.com/badge/github.com/cnguy/gopherjs-frappe-charts)](https://goreportcard.com/report/github.com/cnguy/gopherjs-frappe-charts)
[![GoDoc](https://godoc.org/github.com/cnguy/gopherjs-frappe-charts?status.svg)](https://godoc.org/github.com/cnguy/gopherjs-frappe-charts)

"`Simple, responsive, modern SVG Charts with zero dependencies`"

Event listener with navigable bar chart:

![Alt text](/_pictures/event_listener.gif?raw=true "Event Listener")

Various chart operations (append, remove, pop) with randomized data and actions:

![Alt text](/_pictures/realtime.gif?raw=true "Realtime")

## Contents
* [Installation](#installation)
* [Breaking Changes](#breaking-changes)
* [The Bare Minimum](#the-bare-minimum)
* [Usage](#usage)
* [Disclaimer](#disclaimer)
* [Contributions](#contributions)
* [Changelog](#changelog)

### Installation

*Built on go 1.9*

`go get -u github.com/cnguy/gopherjs-frappe-charts`

Make sure you also have the [necessary JS files](https://github.com/frappe/charts#installation).

Currently, this API supports frappe/charts@0.0.7 (maybe 0.0.8 too, since it seems that they didn't change the API in that version).

Dependencies will be managed via [dep](https://github.com/golang/dep) in the future.

Unfortunately, `gopherjs` is not [vendorable](https://github.com/gopherjs/gopherjs/issues/415) at the moment.

### Breaking Changes

This libary seeks to be updated to always work with the latest version of frappe charts. It also seeks to have a clean API so that users can easily bring in charts into their GopherJS-based website. One can also expect that this library will improve as I improve at Go.

Frappe hasn't changed their API dramatically so far, so one can expect that this library won't change dramatically as well. If there are breaking changes, it should be easy to update one's code by simply following the type errors.

### The Bare Minimum

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
	<div id="chart-2"></div>
	<div id="heatmap"></div>
	<script src="https://unpkg.com/frappe-charts@0.0.7/dist/frappe-charts.min.iife.js"></script>
	<script src="static.js" data-cover></script>
</body>
</html>
```

where `static.js` is the name of your bundled JS file when your folder is named `static` and you run `gopherjs build`.

```go
package main

import (
	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	chartData := charts.NewChartData()
	chartData.Labels = []string{
		"1", "2", "3",
	}
	chartData.Datasets = []*charts.Dataset{
		charts.NewDataset(
			"Some Data",
			[]interface{}{25, 40, 30},
		),
		charts.NewDataset(
			"Another Set",
			[]interface{}{25, 50, -10},
		),
	}

	_ = charts.NewBarChart("#chart", chartData).Render()
}
```
![Alt text](/_pictures/the_bare_minimum.png?raw=true "The Bare Minimum")

This is the bare minimum needed to create a chart. Try swapping `NewBarChart` with other functions such as `NewPercentageChart` or `NewPieChart`.

frappe already sets defaults for a lot of properties. All `New${Type}Chart` calls handles the two necessary properties (the title and the chart data).

Everything else is optional (the title, colors, etc).

To learn how to provide more arguments, scroll down to [Usage](#usage).

### Usage

The basic development flow of this library is:

1) Chart data: Declare your data via NewDataset() and labels (an exception is Heatmap which simply takes a map[string]{interface})
2) Chart args: Create the arguments builder via NewXXXChart(), and then set extra config if needed (heatline, is_navigable, etc), where `XXX` is Bar, Scatter, etc.. Note that it's easy to swap NewXXXChart to a different chart just by changing the name due to the fact that the API for each chart works the same.

Tip: Use WithXXX/SetXXX to chain arguments.

3) Call Render() on the chart args (which may also return a chart object which you can call certain functions on such as ShowSum, ShowAverages). 

Basically, every type of frappe-chart (Bar, Scatter, etc) are divided into separate classes so that there's more type-safety. For example, frappe_chart.show_sums() only works on bar charts, but it's still callable on the line chart (even though it doesn't work). I simply just don't include this method in the struct LineChart. I'm not doing any inheritance stuffs atm since I hacked this together really fast. Will clean up code as I improve though!

* [Hello World](#hello-world)
* [Navigable](#navigable)
* [Line Properties](#line-properties)
* [Simple Aggregations](#simple-aggregations)
* [Event Listener](#event-listener)
* [Heatmaps](#heatmaps)
* [Realtime](#realtime)

#### Hello World:

```go
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

	chart := charts.NewBarChart("#chart", chartData).
		WithHeight(250). // `150` is default if this call is not applied.
		WithColors([]string{"light-blue", "violet"}).
		Render() // not calling Render returns a XXXChartArgs instance instead

	// TIP: Try swapping NewBarChart with NewPercentageChart or NewScatterChart
	// to see how easy it is to swap to a different type of chart.
	// chartArgs.Parent = "#chart"
	// chartArgs.Title = "My Awesome Chart"
	// chartArgs.IsNavigable = ...
	// chartArgs.Heatline = ...
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
	chartData.Labels = []string{
		"2007", "2008", "2009", "2010", "2011", "2012", "2013", "2014", "2015", "2016", "2017",
	}
	values := []interface{}{17, 40, 33, 44, 126, 156, 324, 333, 478, 495, 373}
	dataset := charts.NewDataset("Events", values)
	chartData.Datasets = []*charts.Dataset{dataset}

	chartTitle := "Fireball/Bolide Events - Yearly (more than 5 reports)"
	chart := charts.NewBarChart("#chart", chartData).
		WithTitle(chartTitle).
		WithColors([]string{"orange"}).
		SetIsNavigable(true). // ChartArgs#IsNavigable = 1
		SetIsSeries(true).    // ChartArgs#IsSeries = 1
		Render()

	println(chart)
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
	chartsUtils "github.com/cnguy/gopherjs-frappe-charts/utils"
)

func main() {
	// Prepare data
	chartData := charts.NewChartData()
	chartData.Labels = chartsUtils.NumberLabelsToStr([]int{
		1967, 1968, 1969, 1970, 1971, 1972, 1973, 1974, 1975, 1976,
		1977, 1978, 1979, 1980, 1981, 1982, 1983, 1984, 1985, 1986,
		1987, 1988, 1989, 1990, 1991, 1992, 1993, 1994, 1995, 1996,
		1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006,
		2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015, 2016,
	})
	values := []interface{}{
		132.9, 150.0, 149.4, 148.0, 94.4, 97.6, 54.1, 49.2, 22.5, 18.4,
		39.3, 131.0, 220.1, 218.9, 198.9, 162.4, 91.0, 60.5, 20.6, 14.8,
		33.9, 123.0, 211.1, 191.8, 203.3, 133.0, 76.1, 44.9, 25.1, 11.6,
		28.9, 88.3, 136.3, 173.9, 170.4, 163.6, 99.3, 65.3, 45.8, 24.7,
		12.6, 4.2, 4.8, 24.9, 80.8, 84.5, 94.0, 113.3, 69.8, 39.8,
	}
	dataset := charts.NewDataset("", values)
	chartData.Datasets = []*charts.Dataset{dataset}

	chartTitle := "Mean Total Sunspot Count - Yearly"
	chart := charts.NewLineChart("#chart", chartData).
		WithTitle(chartTitle).
		WithColors([]string{"blue"}).
		SetShowDots(false).
		SetHeatline(true).
		SetRegionFill(true).
		SetXAxisMode("tick").
		SetYAxisMode("span").
		SetIsSeries(true).
		Render()

	println(chart)
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
	chartData := charts.NewChartData()
	chartData.Labels = []string{
		"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat",
	}
	chartData.Datasets = []*charts.Dataset{
		charts.NewDataset("", []interface{}{25, 40, 30, 35, 8, 52, 17}),
		charts.NewDataset("", []interface{}{25, 50, -10, 15, 18, 32, 27}),
	}

	chart := charts.NewBarChart("#chart", chartData).
		WithColors([]string{"purple", "orange"}).
		Render()

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
	barChartData.Labels = []string{
		"2007", "2008", "2009", "2010", "2011", "2012",
		"2013", "2014", "2015", "2016", "2017",
	}
	barChartData.Datasets = []*charts.Dataset{
		charts.NewDataset("Events", reportCountList),
	}
	barChartTitle := "Fireball/Bolide Events - Yearly (more than 5 reports"
	barChart := charts.NewBarChart("#chart", barChartData).
		WithTitle(barChartTitle).
		WithHeight(180).
		WithColors([]string{"orange"}).
		SetIsNavigable(true).
		SetIsSeries(true).
		Render()

	// Make line chart
	lineChartValues := []interface{}{36, 46, 45, 32, 27, 31, 30, 36, 39, 49, 0, 0}
	lineChartData := charts.NewChartData()
	lineChartData.Labels = []string{
		"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec",
	}
	lineChartData.Datasets = []*charts.Dataset{
		charts.NewDataset("", lineChartValues),
	}
	lineChart := charts.NewLineChart("#chart-2", lineChartData).
		WithHeight(180).
		WithColors([]string{"green"}).
		SetIsSeries(true).
		Render()

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
```

Output 5:

![Alt text](/_pictures/event_listener.gif?raw=true "Event Listener")

#### Heatmaps

```go
package main

import (
	"math/rand"
	"strconv"

	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	data := make(map[string]interface{})
	currentTimestamp := 1477699200
	for i := 0; i < 375; i++ {
		data[strconv.Itoa(currentTimestamp)] = rand.Intn(10)
		currentTimestamp += 86400
	}
	charts.NewHeatmapChart("#chart", data).
		WithHeight(115).
		/* Halloween colors */
		// WithLegendColors([]string{"#ebedf0", "#fdf436", "#ffc700", "#ff9100", "#06001c"}).
		Render()
}
```

##### With discrete domains set to default

![Alt text](/_pictures/heatmap_0.gif?raw=true "Heatmap 0")

##### With discrete domains set to 1 via chartArgs.SetDiscreteDomain(true)

![Alt text](/_pictures/heatmap_1.gif?raw=true "Heatmap 1")

##### With different colors

![Alt text](/_pictures/heatmap_2.png?raw=true "Heatmap 2")

#### Realtime

```go
package main

import (
	"math/rand"
	"strconv"
	"time"

	charts "github.com/cnguy/gopherjs-frappe-charts"
)

func main() {
	chartData := charts.NewChartData()
	chartData.Labels = []string{
		"1", "2", "3",
	}
	chartData.Datasets = []*charts.Dataset{
		charts.NewDataset(
			"Some Data",
			[]interface{}{25, 40, 30},
		),
		charts.NewDataset(
			"Another Set",
			[]interface{}{25, 50, -10},
		),
	}

	chart := charts.NewScatterChart("#chart", chartData).
		WithHeight(250).
		WithColors([]string{"red", "green"}).
		Render()
	temp := 4
	for i := 0; i < 100; i++ {
		go func(i interface{}) {
			val := rand.Intn(3) + 2*i.(int)
			println(i, "sleeping for", val)
			time.Sleep(time.Duration(val) * time.Second)
			newLabel := strconv.Itoa(temp)
			actionCond := rand.Intn(4)
			switch actionCond {
			case 0:
				chart.RemoveDataPoint(rand.Int())
			case 1:
				chart.PopDataPoint()
			case 2:
				chart.AppendDataPoint([]interface{}{rand.Intn(1000), rand.Intn(1000)}, newLabel)
			case 3:
				chart.AddDataPoint([]interface{}{rand.Intn(200), rand.Intn(200)}, newLabel, rand.Intn(200))
			}
			temp++
			println(temp)
		}(i)
	}

	println("is this async?")
}
```

Output 6:

![Alt text](/_pictures/realtime.gif?raw=true "Realtime")

### Disclaimer

I'm still a big Go noob, and so I don't know the proper way to manage releases and dependencies yet.
I'll try to figure it out soon!

### Contributions

Any contributions are welcome. :)

### Changelog

[here](https://github.com/cnguy/gopherjs-frappe-charts/blob/master/CHANGELOG.md) when available