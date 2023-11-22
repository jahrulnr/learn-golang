package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Log(title string, msg string) error {
	const path string = "./logs"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)

		if err != nil {
			log.Fatal(err)
			return err
		}
	}

	time := time.Now().Format("2006-01-02")
	file, err := os.OpenFile(path+"/"+time+".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		return err
	}

	log.SetOutput(file)
	errMsg := "[" + title + "] " + msg
	log.Println(errMsg)
	defer file.Close()
	return nil
}
