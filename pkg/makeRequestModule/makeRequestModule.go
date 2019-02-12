package makeRequestModule

import (
	"../CustomHttpClient"
	"../structModule"
)

func MakeCoinProvsionRequestObject(urlPath string) {
	/* 	No      int    `json:"No"`
	   	Name    string `json:"Name"`
	   	ISBN    string `json:"ISBN"`
	   	Forsale string `json:"Forsale"`
	   	Price   int    `json:"Price"` */
	respondUser := structModule.TestStruct{1, "name", "isbn", "n", 1000}

	CustomHttpClient.HttpClientProto(urlPath, respondUser)
}
