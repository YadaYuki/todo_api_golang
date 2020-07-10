package main
import (
	"net/http"
	"encoding/json"
	"todo_api_golang/data"
	"log"
)
type TodoItemList struct{
	TodoItems [] data.TodoItem 
}

func postTodo(writer http.ResponseWriter, request *http.Request){
	len := request.ContentLength
	body := make([]byte,len)
	request.Body.Read(body)
	var todoItem data.TodoItem
	json.Unmarshal(body,&todoItem)
	log.Println(todoItem)
	err := todoItem.Create()
	if err != nil{
		log.Fatal(err)
		return
	}
	writer.WriteHeader(200)
	return
}
func getAllTodo(writer http.ResponseWriter,request *http.Request){
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
func getImportantTodo(writer http.ResponseWriter,request *http.Request){
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
func deleteTodo(writer http.ResponseWriter,request *http.Request){
	len := request.ContentLength
	body := make([]byte,len)
	request.Body.Read(body)
	var todoItem data.TodoItem
	json.Unmarshal(body,&todoItem)
	log.Println(todoItem)
	err := todoItem.DeleteTodoItem()
	if err !=nil{
		log.Fatal(err)
		return
	}
	writer.WriteHeader(200)
	return
}
func updateImportantTodo(writer http.ResponseWriter,request *http.Request){
	len := request.ContentLength
	body := make([]byte,len)
	request.Body.Read(body)
	var todoItem data.TodoItem
	json.Unmarshal(body,&todoItem)
	log.Println(todoItem)
	err := todoItem.UpdateIsImportant()
	if err !=nil{
		log.Fatal(err)
		return
	}
	writer.WriteHeader(200)
	return
}