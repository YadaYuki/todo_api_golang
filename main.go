package main
import (
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
	server.ListenAndServe()
}