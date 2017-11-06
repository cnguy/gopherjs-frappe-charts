package charts

import "honnef.co/go/js/dom"

type eventListenerCallback func(event dom.Event)

func (chart *Chart) AddEventListener(eventName string, useCapture bool, fn eventListenerCallback) {
	parent := chart.Get("parent").Get("id").String()
	d := dom.GetWindow().Document()
	h := d.GetElementByID(parent)
	h.AddEventListener(eventName, useCapture, fn)
}
