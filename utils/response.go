package utils

import "encoding/json"

type output struct {
	Status bool
	Data   any
}

func JsonResponse(status bool, data any) string {
	responseFormat := output{}
	responseFormat.Status = status
	responseFormat.Data = data
	responseJSON, err := json.MarshalIndent(responseFormat, "", "")
	response := string(responseJSON)
	if err != nil {
		Log("JsonResponse", err.Error())
	}

	return response
}
