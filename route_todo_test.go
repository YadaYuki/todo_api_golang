// package main

// import (
// 	"testing"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// )
// var mux *http.ServeMux
// var writer *httptest.ResponseRecorder
// func TestMain(m *testing.M){
// 	setUp()
// 	code := m.Run()
// 	os.Exit(code)
// }
// func setUp(){
// 	mux  = http.NewServeMux()
// 	http.HandleFunc("/post",postTodo);
// 	http.HandleFunc("/get/all",getAllTodo);
// 	http.HandleFunc("/get/important",getImportantTodo);
// 	http.HandleFunc("/delete",deleteTodo);
// 	http.HandleFunc("/update/important",updateImportantTodo);
// 	writer = httptest.NewRecorder()
// }
// func TestHandleGet(t *testing.T){
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/get/all",getAllTodo)
// 	writer := httptest.NewRecorder()
// 	request,_ := http.NewRequest("GET","/get/all",nil)
// 	mux.ServeHTTP(writer,request)
// 	if writer.Code != 200 {
// 		t.Errorf("Response Code Is %v",writer.Code)
// 	}
// 	var todoItemList TodoItemList
// 	json.Unmarshal(writer.Body.Bytes(),&todoItemList)
// 	t.Log(todoItemList)
// 	if todoItemList.TodoItems == nil {
// 		t.Errorf("Json %v",writer.Code)
// 	}
// }