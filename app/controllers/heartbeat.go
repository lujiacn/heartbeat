package controllers

import (
	"github.com/revel/revel"
)

type HeartBeat struct {
	*revel.Controller
}

func (c *HeartBeat) Index() revel.Result {
	msg := map[string]string{"status": "ok"}
	return c.RenderJSON(msg)
}
