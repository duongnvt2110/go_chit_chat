package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gochitchat/model/psql"
	"net/http"
	"time"
)

type User struct {
	Id int64 `json:"id"`
    Uuid string `json:"uuid"`
    Name string `json:"name"`
    Email string `json:"email"`
    Password string `json:"password"`
    CreatedAt time.Time `json:"created_at"`
}

const (
	GetUserQuery = "SELECT * FROM users";
	CreateUserQuery = "INSERT INTO users (uuid,name,email,password,created_at) values($1,$2,$3,$4,$5)"
	ValidationUuidQuery = "SELECT uuid FROM users where uuid = $1"
)

func Index(w http.ResponseWriter, r *http.Request){
	
	rows, err := psql.Psql.Query(GetUserQuery)
	if err != nil {
		http.Error(w, http.StatusText(500),http.StatusInternalServerError) 
		return 
	}

	us := make([]User,0)

	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
		if err != nil {
			http.Error(w, http.StatusText(500),http.StatusInternalServerError) 
			return 
		}
		us = append(us, user)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500),http.StatusInternalServerError) 
		return 
	}

	var users []byte;
	users, err = json.Marshal(us)
	if err != nil {
		http.Error(w, http.StatusText(500),http.StatusInternalServerError) 
		return 
	}
	w.Write([]byte(users))
}

func Create(w http.ResponseWriter, r *http.Request){
	body := json.NewDecoder(r.Body)
	var user User
	err := body.Decode(&user)
	fmt.Println(user.Id)
	if err != nil {
		http.Error(w, http.StatusText(500),http.StatusInternalServerError) 
		return 
	}
	// set time now
	user.CreatedAt = time.Now()
	uuid := user.Uuid 

	row := psql.Psql.QueryRow(ValidationUuidQuery,uuid)
	err = row.Scan(&user.Uuid)
	
	// check duplicate uuid
	if err != sql.ErrNoRows  {
		http.Error(w, http.StatusText(500),http.StatusInternalServerError)
		return
	}

	// execute update func
	_, err = psql.Psql.Exec(CreateUserQuery, user.Uuid, user.Name, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		http.Error(w, http.StatusText(500),http.StatusInternalServerError)
		return
	}
	w.Write([]byte("success to create"))
}

func Show(){

}

func Edit(){

}

func Update(){
	
}

func Delete(){
	
}
