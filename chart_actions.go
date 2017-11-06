package charts

import "github.com/gopherjs/gopherjs/js"

type UpdateValuesArgs struct {
	Values []*UpdateValueSet
	Labels []string
}

type UpdateValueSet struct {
	*js.Object
	Values []interface{} `js:"values"`
}

func NewUpdateValueSet() *UpdateValueSet {
	new := &UpdateValueSet{Object: js.Global.Get("Object").New()}
	return new
}

func (chart *Chart) UpdateValues(updateValues *UpdateValuesArgs) {
	chart.Call("update_values", updateValues.Values, updateValues.Labels)
}

func (chart *Chart) AppendDataPoint(values []interface{}, label string) {
	chart.Call("add_data_point", values, label)
}

func (chart *Chart) AddDataPoint(values interface{}, label string, index int) {
	chart.Call("add_data_point", values, label, index)
}

func (chart *Chart) PopDataPoint() {
	chart.Call("remove_data_point")
}

func (chart *Chart) RemoveDataPoint(index int) {
	chart.Call("remove_data_point", index)
}

func (chart *Chart) ShowSums() {
	chart.Call("show_sums")
}

func (chart *Chart) HideSums() {
	chart.Call("hide_sums")
}

func (chart *Chart) ShowAverages() {
	chart.Call("show_averages")
}

func (chart *Chart) HideAverages() {
	chart.Call("hide_averages")
}
