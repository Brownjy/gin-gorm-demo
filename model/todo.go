package model

// Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

// 创建一个todo
func CreateTodo(todo *Todo) (err error) {
	err = DB.Create(&todo).Error
	return
}

// 获取所有todo
func GetAllTodo() (todoList []*Todo, err error) {
	err = DB.Find(&todoList).Error
	return
}

// 获取一个todo
func GetATodo(id string) (todoList *Todo, err error) {
	err = DB.Where("id=?", id).First(&todoList).Error
	return
}

// 更新一个todo
func UpdateATodo(todo *Todo) (err error) {
	err = DB.Save(&todo).Error
	return
}

// 删除一个todo
func DeleteTodo(id string) (err error) {
	err = DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
