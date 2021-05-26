package tablo

import (
	"encoding/json"
	"net/http"
)

type Device struct {
	IP string
}

func (d *Device) RequestAPI(method, path string, obj interface{}) error {
	req, err := http.NewRequest(method, "http://"+d.IP+":8885"+path, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(obj)
}
