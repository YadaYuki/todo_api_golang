package data
import (
	"time"
	_ "github.com/go-sql-driver/mysql"
)
type TodoItem struct{
	Id int
	Title string
	CreatedOn time.Time
	IsDone bool
	IsImportant bool
}
func (todo *TodoItem) Create() (err error){
	insertTodoQuery := "insert into todo(title) values(?)"
	stmt, err := Db.Prepare(insertTodoQuery)
	if err != nil {
		return
	}
	defer stmt.Close()
	_,err = stmt.Exec(todo.Title);
	return 
}
func(todo *TodoItem) UpdateTitle(id int,title string)(err error){
	updateTodoQuery := "update todo set title=? where id=?"
	stmt,err := Db.Prepare(updateTodoQuery)
	if err != nil{
		return
	}
	defer stmt.Close()
	_,err = stmt.Exec(title,id);
	return
}
func(todo *TodoItem) UpdateIsDone(id int)(err error){
	updateTodoQuery := "update todo set title=? where id=?"
	stmt,err := Db.Prepare(updateTodoQuery)
	if err != nil{
		return
	}
	defer stmt.Close()
	_,err = stmt.Exec(id);
	return
}
func GetAllTodo() (todoItems []TodoItem,err error){
	selectTodoQuery := "select * from todo"
	rows,err := Db.Query(selectTodoQuery)
	if err != nil{
		return
	}
	for rows.Next(){
		todoItem := TodoItem{}
		if err = rows.Scan(&todoItem.Id,&todoItem.Title,&todoItem.CreatedOn,&todoItem.IsDone,&todoItem.IsImportant); err !=nil{
			return
		}
		todoItems = append(todoItems,todoItem)
	}
	rows.Close()
	return
}
func main(){
	GetAllTodo()
}