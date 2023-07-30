package slowreader

import (
	"encoding/json"
	"fmt"

	"github.com/go-faker/faker/v4"
)

type Allarm struct {
	Sensor    string `json:"sensor" faker:"first_name"`
	Timestamp string `json:"timestamp" faker:"timestamp"`
	Severity  int    `json:"severity" faker:"boundary_start=1, boundary_end=5"`
}

func (a *Allarm) String() string {
	return "[" + a.Sensor + "," + a.Timestamp + "," + fmt.Sprint(a.Severity) + "+]"
}

func GiveAllarm() string {
	var alarm Allarm
	faker.FakeData(&alarm)
	b, _ := json.Marshal(alarm)
	return string(b)
}

func GiveAllarmArray(n int) string {
	allarms := make([]Allarm, n)
	faker.FakeData(&allarms)
	b, _ := json.Marshal(&allarms)
	return string(b)
}
