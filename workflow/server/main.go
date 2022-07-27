package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}

func main() {
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})
	http.HandleFunc("/upgrade", func(writer http.ResponseWriter, request *http.Request) {
		mutex.Lock()
		INFO("接收到更新请求")
		defer func() {
			mutex.Unlock()
		}()
		if suc := execCmd("chmod", "a+x", "startup.sh"); !suc {
			return
		}
		if suc := execCmd("sh", "./startup.sh"); !suc {
			return
		}
		INFO("更新结束，开始休眠10s，等待下次更新")
		time.Sleep(time.Second * 10)
		return
	})
	err := http.ListenAndServe(":10012", nil)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}

func execCmd(name string, args ...string) bool {
	cmd := exec.Command(name, args...)
	v, err := cmd.Output()
	if err != nil {
		ERR(fmt.Sprintf("执行[%v]失败，原因：%v", name, err.Error()))
		return false
	}
	INFO("bash-out: " + string(v))
	return true
}

func INFO(val string) {
	cur := time.Now()
	loggerV := fmt.Sprintf("[%v-%v-%v %v:%v:%v]=>INFO: %v", cur.Year(), cur.Month(), cur.Day(), cur.Hour(), cur.Minute(), cur.Second(), val)
	fmt.Println(loggerV)
	writeToLog(loggerV)
}

func ERR(val string) {
	cur := time.Now()
	loggerV := fmt.Sprintf("[%v-%v-%v %v:%v:%v]=>ERROR: %v", cur.Year(), cur.Month(), cur.Day(), cur.Hour(), cur.Minute(), cur.Second(), val)
	fmt.Println(loggerV)
	writeToLog(loggerV)
}

func writeToLog(val string) {
	cur := time.Now()
	prefix := fmt.Sprintf("%v-%v-%v", cur.Year(), cur.Month(), cur.Day())
	file, err := os.OpenFile("upgrade-"+prefix+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("写入log失败！：" + err.Error())
		return
	}
	_, err = file.WriteString(val + "\n")
	if err != nil {
		fmt.Println("写入log失败！：" + err.Error())
		return
	}
}
