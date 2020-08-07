package util

import (
	"encoding/json"

	"github.com/gfbatista/xy-inc/model"
)

//ObjToJSON convert function
func ObjToJSON(object []model.Poi) ([]byte, error) {
	json, err := json.Marshal(object)
	return json, err
}
