package main

import (
	"encoding/json"
	"fmt"
)

var myNetworkDevices = []byte(`
[
  {
    "name": "My AP",
    "lat": 37.4180951010362,
    "lng": -122.098531723022,
    "serial": "Q234-ABCD-5678",
    "mac": "00:11:22:33:44:55",
    "model": "MR34",
    "address": "1600 Pennsylvania Ave",
    "notes": "My AP's note",
    "lanIp": "1.2.3.4",
    "tags": " recently-added ",
    "networkId": "N_24329156",
    "beaconIdParams": {
      "uuid": "00000000-0000-0000-0000-000000000000",
      "major": 5,
      "minor": 3
    }
  }
]
`)

func main() {
	m := []map[string]interface{}{}
	err := json.Unmarshal(myNetworkDevices, &m)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(m)

	type Devices []struct {
		Address        string `json:"address"`
		BeaconIDParams struct {
			Major int64  `json:"major"`
			Minor int64  `json:"minor"`
			UUID  string `json:"uuid"`
		} `json:"beaconIdParams"`
		LanIP     string  `json:"lanIp"`
		Lat       float64 `json:"lat"`
		Lng       float64 `json:"lng"`
		Mac       string  `json:"mac"`
		Model     string  `json:"model"`
		Name      string  `json:"name"`
		NetworkID string  `json:"networkId"`
		Notes     string  `json:"notes"`
		Serial    string  `json:"serial"`
		Tags      string  `json:"tags"`
	}

	devices := Devices{}

	err = json.Unmarshal(myNetworkDevices, &devices)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, device := range devices {
		fmt.Println(device.Name)
		fmt.Println(device.Address)
		fmt.Println(device.Lat)
		fmt.Println(device.Lng)
		fmt.Println(device.Serial)
	}
}
