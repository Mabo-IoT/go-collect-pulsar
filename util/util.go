package util

import (
	// "fmt"
	"github.com/mabo-iot/go-collect/model"
	"math/rand"
	"time"
)

func Generate() []byte {

	measurement := "Device"
	timestamp := time.Now().Unix()
	// seed := time.Now().UnixNano()
	temp := rand.Float32()*(30.0-20.0) + 20.0
	humidity := rand.Float32() * 100.0
	data := &model.Data{
		Measurement: measurement,
		Timestamp:   timestamp,
		Fields: map[string]interface{}{
			"temperature": temp,
			"humidity":    humidity,
		},
		Tags: map[string]interface{}{
			"eqpt_no": "1900-01",
		},
	}

	return data.ToJson()

}
