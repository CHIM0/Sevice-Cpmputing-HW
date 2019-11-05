package main

import (
	"os"
	"strconv"
	flag "github.com/spf13/pflag"
	"net/http"
    "github.com/codegangsta/negroni"
    "github.com/gorilla/mux"
    "github.com/unrolled/render"
)

const (
    PORT string = "8080" //默认端口
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	// 初始化JSON格式化工具
    formatter := render.New(render.Options{
        IndentJSON: true,
    })
	// 初始化服务器对象
	server := negroni.Classic()
	
	// 初始化路由
    router := mux.NewRouter()

	// 设置路由对象
    initRoutes(router, formatter)

    server.UseHandler(router)
    return server
}

//设置路由对象的handle函数,用于处理http请求的url路由信息
func initRoutes(mx *mux.Router, formatter *render.Render) {
	//第一个参数为要处理的路由,第二个参数为用于处理的函数
    mx.HandleFunc("/{num1}/{num2}", requestHandler(formatter)).Methods("GET")
}

// 返回处理函数
func requestHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
		//解析http请求的url
		vars := mux.Vars(req)
		//获取url中的参数,并转化为数字相加
		num1 := vars["num1"]
		num2 := vars["num2"]
		tnum1,error1 := strconv.Atoi(num1)
		tnum2,error2 := strconv.Atoi(num2)
		if error1 != nil || error2 != nil{
			formatter.JSON(w, http.StatusOK, struct{ Error string }{"One of the input is not num"})
		} else {
			// 参数一为用于写response的对象 参数2为请求状态码,参数三为返回的数据
			formatter.JSON(w, http.StatusOK, struct{ Sum string }{strconv.Itoa(tnum1+tnum2)})
		}
    }
}

func main() {
	//若没有设置go环境变量PORT的数值,则使用默认的8080端口
    port := os.Getenv("PORT")
    if len(port) == 0 {
        port = PORT
    }

	//接收命令行程序运行时输入的参数-p / port= 以指定端口
    pPort := flag.StringP("port", "p", PORT, "PORT to listen")
    if len(*pPort) != 0 {
        port = *pPort
    }

	// 初始化服务器
	server := NewServer()

	// 运行服务器在指定端口
    server.Run(":" + port)
}