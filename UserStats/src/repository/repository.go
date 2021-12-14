package repository

import (
	"database/sql"
	"fmt"

	"userStats/src/model"

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

func GetUserById(id string) (model.User, error) { 
	row := db.QueryRow("select (id, rating) from userStatsdb.Users where id = ?", id)
	user := model.User{}
	err := row.Scan(&user.Id, &user.Rating)
	user.TasksSolvedIds, err = GetTasksIdSolved(user.Id)
	return user, err
}

func GetTopUserIds() ([]model.User, error) { 
	rows, err := db.Query("SELECT (id, rating) FROM userStatsdb.Users ORDER BY rating desc")
	var users []model.User
	for rows.Next(){
		var user model.User 
		err := rows.Scan(&user.Id, user.Rating)
		if err != nil{
			continue
		}
		users = append(users, user)
	}
	return users, err
}

func GetTasksIdSolved(userId string) ([]string, error) {
	rows, err := db.Query("select taskId from userStatsdb.UserTasks where userId = ?", userId)
	var taskIds []string
	for rows.Next(){
		var taskId string 
		err := rows.Scan(&taskId)
		if err != nil{
			continue
		}
		taskIds = append(taskIds, taskId)
	}
	return taskIds, err
}

func SaveTaskSolvedByUser(userId string, taskId string) (error) {
	result, err := db.Exec("insert into userStatsdb.UserTasks (userId, taskId) values (?, ?)", userId, taskId)
	fmt.Println(result.RowsAffected())
	return err
}

func SaveUser(user model.User) (error) {
	result, err := db.Exec("insert into userStatsdb.Users (id, rating) values (?, ?) on duplicate key update", user.Id, user.Rating)
	fmt.Println(result.RowsAffected())
	return err
}