package data
import (
	"time"
	_ "github.com/go-sql-driver/mysql"
)
type TodoItem struct {
	Id          int
	Title       string
	CreatedOn   time.Time
	IsDone      bool
	IsImportant bool
}
func (todoItem *TodoItem) Create() (err error) {
	insertTodoQuery := "insert into todo(title) values (?)"
	stmt, err := Db.Prepare(insertTodoQuery)
	if err != nil {
		return
	}
	defer stmt.Close()
	 stmt.Exec(todoItem.Title)
	return
}
func (todoItem *TodoItem) UpdateIsImportant() (err error){
	updateTodoQuery := "update todo set is_important=1-is_important where id=?"
	stmt,err := Db.Prepare(updateTodoQuery)
	if err != nil{
		return
	}
	defer stmt.Close()
	_,err = stmt.Exec(todoItem.Id)
	return
}
func (todoItem *TodoItem) DeleteTodoItem() (err error){
	deleteTodoQuery := "delete from todo where id=?"
	stmt,err := Db.Prepare(deleteTodoQuery)
	if err != nil{
		return
	}
	defer stmt.Close()
	_,err = stmt.Exec(todoItem.Id)
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
		if err = rows.Scan(&todoItem.Id,&todoItem.Title,&todoItem.CreatedOn,&todoItem.IsDone,&todoItem.IsImportant) ;err != nil{
			return
		}
		todoItems = append(todoItems,todoItem)
	}
	rows.Close()
	return
}
func GetImportantTodo() (todoItems []TodoItem, err error){
	selectImportantTodoQuery := "select * from todo where is_important=true";
	rows,err := Db.Query(selectImportantTodoQuery)
	if err != nil{
		return
	}
	for rows.Next(){
		todoItem := TodoItem{}
		if err=rows.Scan(&todoItem.Id,&todoItem.Title,&todoItem.CreatedOn,&todoItem.IsDone,&todoItem.IsImportant) ; err != nil{
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