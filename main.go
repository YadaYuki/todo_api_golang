package main
import (
	"golang.org/x/net/http2"
	"net/http"
)
func main(){
	// mux := http.NewServeMux()
	http.HandleFunc("/post",postTodo);
	http.HandleFunc("/get/all",getAllTodo);
	http.HandleFunc("/get/important",getImportantTodo);
	http.HandleFunc("/delete",deleteTodo);
	http.HandleFunc("/update/important",updateImportantTodo);
	server := http.Server{
		Addr:"localhost:8000",
	}
	http2.ConfigureServer(&server,&http2.Server{})
	server.ListenAndServe()
}