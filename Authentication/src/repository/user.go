package repository

import (
	"database/sql"

	"authentication/src/model"
)

func SaveUser(user model.User) (error){
	result, err := db.Exec("insert into securitydb.Users (id, login, encryptedPassword) values (?, ?, ?)", user.Id, user.Login, user.EncryptedPassword)
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
