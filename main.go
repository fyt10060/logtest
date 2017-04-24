// main
package main

import (
	"fmt"
	//	"html/template"
	//	"log"
	"net/http"
	"os"
	//	"strings"
	"encoding/json"
	"io/ioutil"
	"time"

	"logtest/controllers"

	"github.com/astaxie/beego"
)

const fileName string = "requestLog.txt"

type ApiVersion struct {
	Host         string `json:"host"`
	Version      string `json:"version"`
	AddonHost    string `json:"addon_host"`
	AddonVersion string `json:"addon_version"`
}

type ApiVersionSlice struct {
	Apis []ApiVersion `json:"data"`
}

func main() {
	//	fmt.Println("Hello World!")
	//	createFile()
	beego.Router("/", &controllers.RouterController{})
	//	beego.Router("/:func", &mainController{})
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
	case "api_version":
		getApiAddress(w)
	default:
		httpNotFound(w, r)
		go saveRequestToLog(w, r)
	}

}

func getApiAddress(w http.ResponseWriter) {
	var s ApiVersionSlice
	s.Apis = append(s.Apis, ApiVersion{Host: "apidev.weixinhost.com", Version: "/4/", AddonHost: "api.weixinhost.com", AddonVersion: "/3/"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("no json")
		return
	}
	fmt.Println(string(b))
	fmt.Fprintf(w, string(b))
}

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

func httpNotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your page is collected by ET, please ask him first")
}
