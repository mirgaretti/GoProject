package repository

import (
	"database/sql"

	"tasks/src/model"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() { 
	var err error
	db, err = sql.Open("mysql", "root:ML$P-%J%+2-9z$Y&@/tasksdb")
	if err != nil {
	panic(err)
	} 
	defer db.Close()
}

func GetTasksByTopicLabel(topicLabel string) ([]model.Task, error){
	rows, err := db.Query("select (id, label, text, rating) from tasksdb.Tasks where topicLabel = ?", topicLabel)
	var tasks []model.Task
	for rows.Next(){
		task := model.Task{}
		err := rows.Scan(&task.Id, &task.Label, &task.Text, &task.Rating)
		if err != nil{
			continue
		}
		tasks = append(tasks, task)
	}
	if err != nil{
		panic(err)
	}
	return tasks, err
}

func GetTopicsByClassLabel(classLabel int16) ([]model.Topic, error){
	rows, err := db.Query("select * from tasksdb.Topics where classLabel = ?", classLabel)
	var topics []model.Topic
	for rows.Next(){
		topic := model.Topic{}
		err := rows.Scan(&topic.Label, &topic.Description)
		if err != nil{
			continue
		}
		topic.Tasks, err = GetTasksByTopicLabel(topic.Label)
		topics = append(topics, topic)
	}
	if err != nil{
		panic(err)
	}
	return topics, err
}

func GetAllClasses() ([]model.Class, error){
	rows, err := db.Query("select * from tasksdb.Classes")
	var classes []model.Class
	for rows.Next(){
		class := model.Class{}
		err := rows.Scan(&class.Label, &class.Description)
		if err != nil{
			continue
		}
		class.Topics, err = GetTopicsByClassLabel(class.Label)
		classes = append(classes, class)
	}
	if err != nil{
		panic(err)
	}
	return classes, err
}
