package TemporyLogger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func FileAppend(respondUser interface{}, str string, urlPath string) {

	file, err := os.OpenFile("apilog.log", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	out, err := json.Marshal(respondUser)
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	len, err := file.WriteString("path " + urlPath + " [" + time.Now().Format("Mon Jan _2 15:04:05 2006") + "]" + string(out) + "\n" + str + "\n")
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	fmt.Printf("\nLength: %d bytes", len)
	fmt.Printf("\nFile Name: %s", file.Name())
}
