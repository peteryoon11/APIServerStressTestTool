package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"../pkg/structModule"
)

/*
test 시나리오 작성

유효한 유저 정보를 20개 정도 준비하고
 1000여번 정도 반복하고 로그를 확인한다.

*/
func main() {

	// test 용
	url := "http://localhost:8083"

	path := "/service/book"
	var (
		respondUser structModule.AuthRequestObject
		userId      string
		AuthKey     string
	)
	nowInit := time.Now()

	userId = "dd54092"
	AuthKey = "testAuth1232312"
	respondUser = structModule.AuthRequestObject{userId, AuthKey}
	var num int
	fmt.Print(url + path)
	fmt.Println(respondUser)

	for num < 100 {
		num++
		fmt.Println(num)
		time.Sleep(time.Millisecond * 250)

		// 최초에 들어가는
		go HttpClientTester(url+path, respondUser, num)
	}

	/*
		1000ms / 1sec 에 30개의 요청을 날리고 1000ms 에서 시간이 남아 있으면 sleep 을 걸어 둘까.
		그걸 1분 정도 동안 반복 하는 걸로

	*/
	finishTime := time.Now().Sub(nowInit)

	fmt.Println("=============")
	fmt.Println(finishTime.Seconds())
}
func PrintTest(num int) {
	fmt.Println(num, " test")
	time.Sleep(time.Millisecond * 1000)

}
func HttpClientTester(urlPath string, respondUser structModule.AuthRequestObject, num int) {

	pbytes, _ := json.Marshal(respondUser)
	buff := bytes.NewBuffer(pbytes)

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
	fmt.Println(num)

}
