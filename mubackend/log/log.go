package log

import (
	"fmt"
	"io"
	"log"
	"mubackend/config"
	"os"
)

var Log *log.Logger

//日志初始化，同时打印到控制台和文件
func InitLog(){
	_, err := os.Stat(config.Cfg.Log)
	if err !=nil{
		create, _ := os.Create(config.Cfg.Log)
		create.Close()
	}
	file, err := os.OpenFile(config.Cfg.Log, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err !=nil {
		fmt.Println(err)
		fmt.Println(11)
	}
	mulWriter := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mulWriter)
	log.SetFlags(log.Ldate|log.Ltime|log.Lshortfile)
}
