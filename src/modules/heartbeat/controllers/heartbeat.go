package heartbeat_controller

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct{
	Status string `json:"Status"`
	Data string `json:"Data"`
}

type ErrorResponseData struct{
	Status string `json:"Status"`
	ErrorMessage string `json:"ErrorMessage"`
}

func Handle(res http.ResponseWriter, req *http.Request){
	// Defining content type
	res.Header().Set("Content-Type", "application/json")
	
	// Creating response data
	var dataDetails = ResponseData{Status: "success", Data: "GoTickets API Server"}

	// Formatting data to return
	var data, err = json.Marshal(dataDetails)
	if(err != nil){
		var dataErrorDetails = ErrorResponseData{Status: "error", ErrorMessage: "Failed to format data"}
		data, _ = json.Marshal(dataErrorDetails)
		res.WriteHeader(http.StatusInternalServerError)
	}else{
		res.WriteHeader(http.StatusOK)
	}
	
	res.Write(data)
}