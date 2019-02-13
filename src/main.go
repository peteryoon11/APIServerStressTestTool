package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"../pkg/makeRequestModule"
	"../pkg/structModule"
)

/*
test 시나리오 작성

유효한 유저 정보를 20개 정도 준비하고
 1000여번 정도 반복하고 로그를 확인한다.

*/
func main() {

	// test 용
	//url := "http://localhost:8083"
	url := "http://localhost:8090"

	path := "/service/book"
	var (
		respondUser structModule.ValAPIKey
		userId      string
		AuthKey     string
	)
	nowInit := time.Now()

	userId = "dd54092"
	AuthKey = "testAuth1232312"
	//respondUser = structModule.AuthRequestObject{userId, AuthKey}
	respondUser = structModule.ValAPIKey{userId, AuthKey}
	var num int
	//fmt.Print(url + path)
	//	fmt.Println(respondUser)

	for num < 10 {
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

func ThreadFlowControl() {
	url := "http://localhost:8080/test/server"

	path := "/test/create" //
	makeRequestModule.MakeCoinProvsionRequestObject(url + path)

}
func ProcessCore() {

}
func initFunc(startWord []string) {
	var (
		webpageAddress string
		filepath       string
		identify       string
		loggerLocate   string
		/* 	workerRecorder *log.Logger
		fpLog          *os.File
		*/)
	//	startTime := time.Now() // 처음부터 끝까지 걸린 시간을 측정 하기 위한 시작시간 체크

	for _, item := range os.Args[1:] {
		if temp := strings.Split(item, "=")[0]; strings.EqualFold(temp, "site") {
			webpageAddress = strings.Split(item, "=")[1]
		}
		if temp := strings.Split(item, "=")[0]; strings.EqualFold(temp, "path") {
			filepath = strings.Split(item, "=")[1]
		}
		if temp := strings.Split(item, "=")[0]; strings.EqualFold(temp, "identi") {
			identify = strings.Split(item, "=")[1]
		}
		if temp := strings.Split(item, "=")[0]; strings.EqualFold(temp, "logger") {
			// logger 파일의 위치
			// 기본은 현재 실행 파일과 같은 경로  ./
			// 다른 옵션을 주면 src 위에 logs 로 가게 하기 위해서 ../ 를 하게 하는 방법?
			// 현재는 ../ 로 되어 있음 테스트 하기 위해서 그런데 이게 윈도우에서 똑같이 동작할까?
			loggerLocate = strings.Split(item, "=")[1]
		}
	}

	if len(webpageAddress) == 0 {
		webpageAddress = "https://kissme2145.tistory.com/1418?category=634440"
		//webpage = "https://comic.naver.com/webtoon/detail.nhn?titleId=675554&no=683"
		// 나중에 네이버 웹툰 페이지도 추가해 보자.
	}
	if len(filepath) == 0 {
		filepath = "temp"
	}
	if len(identify) == 0 {
		identify = "image"
	}
	if len(loggerLocate) == 0 {
		loggerLocate = "logs"
	}
	ProcessCore()
}

func PrintTest(num int) {
	fmt.Println(num, " test")
	time.Sleep(time.Millisecond * 1000)

}

//func HttpClientTester(urlPath string, respondUser structModule.AuthRequestObject, num int) {
func HttpClientTester(urlPath string, respondUser structModule.ValAPIKey, num int) {
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
