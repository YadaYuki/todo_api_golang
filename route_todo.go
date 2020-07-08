package main
import (
	"fmt"
	"net/http"
	"encoding/json"
	"todo_api_golang/data"
	"log"
)
type TodoItemList struct{
	TodoItems [] data.TodoItem 
}
func Post(writer http.ResponseWriter, request *http.Request){
	
	// fmt.Fprintf(writer,"%s!",request.URL.Path[1:])
	return
}
func GetAll(writer http.ResponseWriter,request *http.Request){
	todoItems,err := data.GetAllTodo()
	if err !=nil{
		log.Fatal(err)
		return
	}
	log.Println(todoItems)
	todoItemList := TodoItemList{TodoItems:todoItems}
	output,err:=json.MarshalIndent(&todoItemList,"","")
	if err !=nil{
		log.Fatal(err)
		return
	}
	log.Println(output)
	writer.Header().Set("Content-Type","application/json")
	writer.Write(output)
}
func getImportantItem(writer http.ResponseWriter,request *http.Request){
	todoItems,err := data.GetImportantTodo()
	log.Println(todoItems)
	if err != nil{
		log.Fatal(err)
		return
	}
	todoItemList := TodoItemList{TodoItems:todoItems}
	output,err := json.MarshalIndent(&todoItemList,"","")
	if err != nil{
		log.Fatal(err)
		return
	}
	log.Println(output)
	writer.Header().Set("Content-Type","application/json")
	writer.Write(output)
}
func updateIsDone(writer http.ResponseWriter,request *http.Request){
	fmt.Fprintf(writer, "UPDATE IS DONE")
}
func updateImportant(writer http.ResponseWriter,request *http.Request){
	fmt.Fprintf(writer, "UPDATE IMPORTANT")
}
func main(){
	server := http.Server{
		Addr:"localhost:8000",
	}
	http.HandleFunc("/post",Post);
	http.HandleFunc("/get/all",GetAll);
	http.HandleFunc("/get/important",getImportantItem);
	http.HandleFunc("/update/is_done",updateIsDone);
	http.HandleFunc("/update/important",updateImportant);
	server.ListenAndServe()
}