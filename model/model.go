package model

import (
	"encoding/json"
)

type Data struct {
	Measurement string                 `json: "measurement"`
	Timestamp   int64                  `json: "timestamp"`
	Fields      map[string]interface{} `json: "fields"`
	Tags        map[string]interface{} `json: "tags"`
}

func (d *Data) ToJson() []byte {
	b, err := json.Marshal(d)
	if err != nil {
		panic(err)
	}
	return b
}
