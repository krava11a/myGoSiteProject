package model

import "myGoSiteProject/app/server"

type User struct {
	Id      int
	Name    string
	Surname string
}

func NewUser(name string, surname string) *User {
	return &User{
		Name:    name,
		Surname: surname,
	}
}

func (u *User) AddUser() (err error) {
	query := `INSERT INTO dbo.users (name, surname) VALUES (@p1,@p2)`
	_, err = server.Db.Exec(query, u.Name, u.Surname)
	return
}

func GetUserById(id string) (u User, err error) {
	query := `SELECT * FROM dbo.users WHERE id = @p1`
	err = server.Db.Get(&u, query, id)
	return
}

func (u *User) DeleteUser() (err error) {
	query := `DELETE FROM dbo.users WHERE id = @p1`
	_, err = server.Db.Exec(query, u.Id)
	return
}

func (u *User) UpdateUser() (err error) {
	query := `UPDATE users SET name = @p1, surname = @p2 WHERE id = @p3`
	_, err = server.Db.Exec(query, u.Name, u.Surname, u.Id)
	return
}

func GetAllUsers() (users []User, err error) {
	//variant sqlx
	query := `SELECT * FROM test.dbo.users`
	err = server.Db.Select(&users, query)
	return

	//variant sql
	//query := `SELECT * FROM test.dbo.users`
	//rows, err := server.Db.Queryx(query)
	//if err != nil {
	//	return users, err
	//}
	//defer rows.Close()
	//user := User{}
	//for rows.Next() {
	//	err = rows.StructScan(&user)
	//	if err != nil {
	//		return users, err
	//	}
	//	users = append(users, user)
	//}
	//return users, nil
	//

	//users = []User{
	//	{1,"Джон","До"},
	//	{2,"Говард","Рорк"},
	//	{3,"Джек","Доусон"},
	//	{4,"Лизель","Мемингер"},
	//	{5,"Джейн","Эйр"},
	//	{6,"Мартин","Иден"},
	//	{7,"Джон","Голт"},
	//	{8,"Сэмвелл","Тарли"},
	//	{9,"Гермиона","Грейнджер"},
	//}
	//return
}
