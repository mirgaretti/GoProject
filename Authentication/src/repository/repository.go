package repository

import (
	"database/sql"
	"fmt"
	"authentication/src/model"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() { 
	var err error
	db, err = sql.Open("mysql", "root:ML$P-%J%+2-9z$Y&@/securitydb")
	if err != nil {
	panic(err)
	} 
	defer db.Close()
}

func SaveUser(user model.User) (error){
	result, err := db.Exec("insert into securitydb.Users (id, login, encryptedPassword) values (?, ?, ?)", user.Id, user.Login, user.EncryptedPassword)
	fmt.Printf("result: %v\n", result)
	return err;
}

func GetUserByLogin(login string) (model.User, error){
	row := db.QueryRow("select * from securitydb.Users where login = ?", login)
	user := model.User{}
	err := row.Scan(&user.Id, &user.Login, &user.EncryptedPassword)
	if err != nil{
		panic(err)
	}
	return user, err
}

func SaveSession(session model.UserSession) error {
	result, err := db.Exec("insert into securitydb.UserSession (userId, accessToken, startTime, endTime) values (?, ?, ?, ?) on duplicate key update",
		session.UserId, session.AccessToken, session.StartTime, session.ExpireTime)
	fmt.Printf("result: %v\n", result)
	return err;
}

func GetSessionByToken(accessToken string) (model.UserSession, error) {
	row := db.QueryRow("select * from securitydb.UserSession where accessToken = ?", accessToken)
	session := model.UserSession{}
	err := row.Scan(&session.UserId, &session.AccessToken, &session.StartTime, &session.ExpireTime)
	if err != nil{
		panic(err)
	}
	return session, err
}