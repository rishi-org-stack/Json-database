package main

import (
	// "fmt"/
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	mp "github.com/rishi-org-stack/proj-track/mapping"
	mo "github.com/rishi-org-stack/proj-track/model/json"
	ut "github.com/rishi-org-stack/proj-track/util"
)

func greet(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, "h")
}

func check(w http.ResponseWriter, r *http.Request) {
	f := mp.Function(1).(func(int) (bool, error))
	id, _ := strconv.Atoi(r.FormValue("id"))
	pres, _ := f(id)
	if pres {
		t, _ := template.ParseFiles("templates/home.html")
		data := mo.ReadUserId(id)
		user := mo.Rawuser(data)
		// projects:= mo.Rawprojectarray(user.Projs)
		t.Execute(w, user)
	} else {
		t, _ := template.ParseFiles("templates/signup.html")
		t.Execute(w, "heloo")
	}
}

func signup(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	age, _ := (strconv.Atoi(r.FormValue("age")))
	Age := float64(age)
	id, _ := strconv.Atoi(r.FormValue("id"))
	Id := float64(id)
	pass := r.FormValue("password")
	f := mp.Function(3).(func(string, float64, string, float64))
	f(name, Age, pass, Id)
	// fmt.Fprintf(w,"<a href=>home</a>")
}
func addprojtemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/addproject.html")
	t.Execute(w, "")
}

func addproj(w http.ResponseWriter, r *http.Request) {
	f := mp.Function(4).(func(int, float64, bool, string, string, string))
	name := r.FormValue("name")
	pid, _ := (strconv.Atoi(r.FormValue("pid")))
	id, _ := (strconv.Atoi(r.FormValue("id")))
	ID := int(id)
	PID := float64(pid)
	ldate := r.FormValue("endday")
	lmonth := string(r.FormValue("endmonth"))
	lyear := r.FormValue("endyear")
	last := ut.Dateconstruct(ldate, lmonth, lyear)
	sday := r.FormValue("startday")
	smonth := r.FormValue("startmonth")
	syear := r.FormValue("startyear")
	start := ut.Dateconstruct(sday, smonth, syear)
	fmt.Println(ut.Parse(last))
	f(ID, PID, false, name, last, start)
}

func detilsprojtemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/detailsofproject.html")
	t.Execute(w, "")
}

func getdetilsofproject(w http.ResponseWriter, r *http.Request) {
	f := mp.Function(6).(func(int, float64) map[string]interface{})
	pid, _ := strconv.Atoi(r.FormValue("ID"))
	PID := float64(pid)
	id, _ := strconv.Atoi(r.FormValue("UID"))
	ID := int(id)
	Rawproject := f(ID, PID)
	project :=mo.Rawproject(Rawproject)
	fmt.Fprint(w,project.Completed)
}
func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/check", check)
	http.HandleFunc("/check/signup", signup)
	http.HandleFunc("/home/add", addprojtemplate)
	http.HandleFunc("/home/add/proj", addproj)
	http.HandleFunc("/home/dproj", detilsprojtemplate)
	http.HandleFunc("/home/dproj/get", getdetilsofproject)
	http.ListenAndServe(":8080", nil)
}
