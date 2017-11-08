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
	dataset := charts.NewDataset("", "blue", values)
	chartData.Datasets = []*charts.Dataset{dataset}

	// Prepare constructor arguments
	chartTitle := "Mean Total Sunspot Count - Yearly"
	chartArgs := charts.NewLineChartArgs("#chart", chartTitle, chartData, 180)
	chartArgs.SetShowDots(false)
	chartArgs.SetHeatline(true)
	chartArgs.SetRegionFill(true)
	chartArgs.XAxisMode = "tick"
	chartArgs.YAxisMode = "span"
	chartArgs.SetIsSeries(true)

	chart := chartArgs.Render()
	println(chart)
}
