package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	cur_path := r.URL.Path
	switch cur_path {
	case "/":
		fmt.Fprint(w, "<h1>Hello, 这里是 goblog</h1>")
	case "/about":
		fmt.Fprint(w, "此11博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
			"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
	default:
		fmt.Fprint(w, "<h1>请求页面未找到 :(</h1>"+
			"<p>如有疑惑，请联系我们。</p>")

	}
}
func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
