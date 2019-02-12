package CustomHttpClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../TemporyLogger"
)

func HttpClientProto(urlPath string, respondUser interface{}) {

	//respondUser = structModule.CoinProvsionRequestObject{userId, serviceProviderId}
	//person := Person{"Alex", 10}
	pbytes, _ := json.Marshal(respondUser)
	buff := bytes.NewBuffer(pbytes)
	//req, err := http.NewRequest("POST", url+path, buff)
	req, err := http.NewRequest("POST", urlPath, buff)
	req.Header.Add("User-Agent", "StressTestAgent")
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)
	fmt.Println(">>>>>>>>>>>>>>>>>>>")

	TemporyLogger.FileAppend(respondUser, str, urlPath)

}
