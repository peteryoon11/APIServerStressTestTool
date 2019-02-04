package structModule

type AuthRequestObject struct {
	UserId  string `json:"userId"`
	AuthKey string `json:"AuthKey"`
}

type ResultObjct struct {
	Result string `json:"Results"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ValAPIKey struct {
	ID     string `json:"ID"`
	APIKey string `json:"APIKey"`
}
type EBookInfo struct {
	No      int    `json:"No"`
	Name    string `json:"Name"`
	ISBN    string `json:"ISBN"`
	Forsale string `json:"Forsale"`
	Price   int    `json:"Price"`
}
