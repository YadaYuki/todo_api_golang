package main
import (
	"net/http"
)
func main(){
	
	// mux := http.NewServeMux()
	http.HandleFunc("/post",post);
	http.HandleFunc("/get/all",getAll);
	http.HandleFunc("/get/important",getImportantItem);
	http.HandleFunc("/delete",delete);
	http.HandleFunc("/update/important",updateImportant);
	server := http.Server{
		Addr:"localhost:8000",
	}
	server.ListenAndServe()
}