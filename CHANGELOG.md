# 0.0.2

This version works with frappe/charts v0.0.2.

This library exposes a cleaner API for creating charts now.

Here are the changes that frappe made since frappe/charts@0.0.4:

```markdown
##### v0.0.7
- [Custom color values](https://github.com/frappe/charts/pull/71) for charts as hex codes. The API now takes an array of colors for all charts instead of a color for each dataset.
- [@iamkdev's](https://github.com/iamkdev) blog on [usage with Angular](https://medium.com/@iamkdev/frappé-charts-with-angular-c9c5dd075d9f).

##### v0.0.5
- More [flexible Y values](https://github.com/frappe/charts/commit/3de049c451194dcd8e61ff91ceeb998ce131c709): independent from exponent, minimum Y axis point for line graphs.
- Customisable [Heatmap colors](https://github.com/frappe/charts/pull/53); check out the Halloween demo on the [website](https://frappe.github.io/charts) :D
- Tooltip values can be [formatted](https://github.com/frappe/charts/commit/e3d9ed0eae14b65044dca0542cdd4d12af3f2b44).
```

This library now:
1. takes in colors in the way described above.
2. allows heatmap legend colors through HeatmapArgs#LegendColors.
3. allows users to format the tooltips (hover over data to see them) by providing callbacks to WhateverChart#FormatTooltipX and WhateverChart#FormatTooltipY.
    * example:
    ```go
    chartArgs.FormatTooltipX(
        func (d string)string {
            return d + "."
        }
    )
    ```

# 0.0.1

This version works with frappe/charts v0.0.4.