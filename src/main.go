package main
import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"encoding/json"
)
//写一个json 数据，作为http返回的内容
type ServerTime struct {
	MName string
	MTime string

}
type ServerMomter struct {
	Servers []ServerTime
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	var s ServerMomter
	str := `{"servers":[{"mName":"wuhan","mTime":"2016/08/02"},{"mName":"Beijing","mTime":"2017/08/02"}]}`
	json.Unmarshal([]byte(str), &s)
	fmt.Fprintf(w,str) //这个写入到w的是输出到客户端的


}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
