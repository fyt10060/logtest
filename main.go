// main
package main

import (
	"fmt"
	//	"html/template"
	//	"log"
	"net/http"
	"os"
	//	"strings"
	"io/ioutil"
	"time"

	"github.com/astaxie/beego"
)

const fileName string = "requestLog.txt"

//type MyMux struct{}

//func (p *MyMux) ServeHttp(w http.ResponseWriter, r *http.Request) {
//	switch r.URL.Path {
//	case "/log":
//		fmt.Fprintf(w, getLog(w, r))
//	default:
//		saveRequestToLog(w, r)
//		fmt.Fprintf(w, "{\"err_code\":0, \"data\": \"success\"")
//	}
//}

func getLog(w http.ResponseWriter, r *http.Request) {
	fl, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer fl.Close()
	buf, err := ioutil.ReadAll(fl)
	if err != nil {
		fmt.Println("file not exist")
		return
	}
	if len(buf) == 0 {
		fmt.Println("buf is empty")
		return
	} else {
		fmt.Fprintf(w, string(buf))
		return
	}
}

func saveRequestToLog(w http.ResponseWriter, r *http.Request) {
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error")
		return
	}
	r.ParseForm()
	t := time.Now()

	file.WriteString(t.Format("2006-01-02 15:04:06"))

	file.WriteString(" ip:" + r.RemoteAddr)
	file.WriteString("\n")
	file.Close()
}

func main() {
	//	fmt.Println("Hello World!")
	createFile()
	beego.Router("/:func", &mainController{})
	beego.Run()
}

func createFile() {
	fl, err := os.Create(fileName)
	if err != nil {
		return
	}
	defer fl.Close()
	fl.WriteString("file create time: ")
	fl.WriteString(time.Now().Format("2006-01-02 15:04:05"))
	fl.WriteString("\n")
}

type mainController struct {
	beego.Controller
}

func (c *mainController) Get() {
	operation := c.Ctx.Input.Param(":func")
	w := c.Ctx.ResponseWriter
	r := c.Ctx.Request
	switch operation {
	case "log":
		getLog(w, r)
	case "push":
		break
	default:
		httpNotFound(w, r)
		go saveRequestToLog(w, r)
	}

}

func httpNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your page is collected by ET, please ask him first")
}
