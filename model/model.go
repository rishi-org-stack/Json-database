package model

import (
	"log"
)

/*User hold type user which is basically abs class*/
type User struct {
	ID       float64   `json:"id"`
	Name     string    `json:"name"`
	Age      float64   `json:"age"`
	Password string    `json:"password"`
	Projs    []Project `json:"projs"`
}

// /*Project base proj class*/
type Project struct {
	ID        float64
	Name      string
	Completed bool
	Lastdate  string
	Startdate string
	//todo specification and tag type
}

// type Project map[string]interface{}
//todo
type database []User

//Newuser create a new user
func Newuser(name string, age float64, pass string) *User {
	//for now i dont know how to allocate ids
	var user = new(User)
	if user == nil {
		log.Fatal("stop memory error in creating user")
	}
	user.Name = name
	user.Age = age
	user.Password = pass
	return user
}

//Newproj creates a new projrct instance
func Newproj(name string, date1 string, date2 string) map[string]interface{} {
	var p =make(map[string]interface{})
	// if p == nil {
	// 	log.Fatal("stop memory error in creating proj")
	// }
	// p =make(map[string]interface{})
	p["name"] = name
	p["completed"] = false
	p["startdate"]=date1
	p["lastdate"] = date2
	return p
}

//Addproj adds a proj in list of projects of given user
//need improvement it delets previous projs
// func (u *User) Addproj(p map[string]interface{}) {
// 	if len(u.Projs) > 0 {
// 		for _, proj := range u.Projs {
// 			if proj["id"]!= p["id"] && proj["name"] != p["name"] {
// 				u.Projs = append(u.Projs, p)
// 			}
// 		}
// 	} else {
// 		u.Projs = append(u.Projs, p)
// 	}
// }
