package charts

import (
	"github.com/gopherjs/gopherjs/js"
)

type DataSelectEvent struct {
	*js.Object
	Value interface{} `js:"value"`
	Label interface{} `js:"label"`
	Index int         `js:"index"`
}

type eventListenerCallback func(event *DataSelectEvent)

func (chart *Chart) AddEventListener(eventName string, fn eventListenerCallback) {
	parent := chart.Get("parent").Get("id").String()
	element := js.Global.Get("document").Call("getElementById", parent)
	element.Call("addEventListener", eventName, fn)
}

// AddDataSelectListener is a helper function to allow users to quickly set a data-select event.
// This allows users to access the index, value, and label of any particular data point!
func (chart *Chart) AddDataSelectListener(fn eventListenerCallback) {
	chart.AddEventListener("data-select", fn)
}
