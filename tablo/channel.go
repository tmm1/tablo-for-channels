package tablo

import (
	"encoding/json"
	"fmt"
)

type Channel struct {
	ID   json.Number `json:"object_id"`
	Path string      `json:"path"`
	Info struct {
		CallSign    string `json:"call_sign"`
		CallSignSrc string `json:"call_sign_src"`
		Major       int    `json:"major"`
		Minor       int    `json:"minor"`
		Network     string `json:"network"`
		Resolution  string `json:"resolution"`
	} `json:"channel"`
}

func (c *Channel) Number() string {
	return fmt.Sprintf("%v.%v", c.Info.Major, c.Info.Minor)
}
