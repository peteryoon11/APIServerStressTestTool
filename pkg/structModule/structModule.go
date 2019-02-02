package structModule

type AuthRequestObject struct {
	UserId  string `json:"userId"`
	AuthKey string `json:"AuthKey"`
}

type ResultObjct struct{
	Result string `json:"Results"`
}
