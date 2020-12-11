package json

//we cant update id 3 and 5

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	mo "github.com/rishi-org-stack/proj-track/model"
	ut "github.com/rishi-org-stack/proj-track/util"
)

const (
	add = "C:/Users/Documents/Langauges/rest api/Go/proj/proj3/data/"
	ext = ".json"
)

func Find(id int) (bool, error) {
	present := false
	name := strconv.Itoa(id) + ext
	files, err := ioutil.ReadDir(add)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if file.Name() == name {
			present = true
		}
	}
	return present, nil
}

func ReadUserId(id int) map[string]interface{} {
	title := strconv.Itoa(id) + ext
	var res interface{}
	pres, _ := Find(id)
	if pres {
		data, _ := ioutil.ReadFile(add + title)
		err := json.Unmarshal(data, &res)
		if err != nil {
			log.Fatal(err)

		}
	}
	return res.(map[string]interface{})
}

func Insert(name string, age float64, pass string, id float64) {
	var u = mo.Newuser(name, age, pass)
	u.ID = id
	title := strconv.Itoa(int(u.ID))
	data, err := json.Marshal(u)
	pres, _ := Find(int(u.ID))
	if err != nil {
		log.Fatal(err)
	}
	if pres {
		fmt.Println("already present")
	} else {
		_, er := os.Create(add + title + ext)
		if er != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(add+title+ext, data, 7)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("done !")
		}
	}

}

func convert(ar []interface{}) []map[string]interface{} {
	res := make([]map[string]interface{}, 0)
	for i := 0; i < len(ar); i++ {
		res = append(res, ar[i].(map[string]interface{}))
	}
	return res
}


func AddProj(id int, pid float64, comp bool, name string, ldate string, sdate string) {
	if pre, _ := Find(id); pre == true {
		if ut.Isvalid(ldate) && ut.Isvalid(sdate) {
			title := strconv.Itoa(id)
			oldinfo := ReadUserId(id)
			proj := mo.Newproj(name, sdate, ldate)
			proj["pid"] = pid
			if oldinfo["projs"] == nil {
				oldinfo["projs"] = make([]interface{},0)
				oldinfo["projs"] = append(oldinfo["projs"].([]interface{}), proj)
				data, _ := json.MarshalIndent(oldinfo, "", " ")
				fmt.Println(oldinfo)
				os.Remove(add + title + ext)
				os.Create(add + title + ext)
				ioutil.WriteFile(add+title+ext, data, 7)
			} else {
				//if oldinfo["projs"]!=nil{
				oldinfo["projs"] = append(oldinfo["projs"].([]interface{}), proj)
				data, _ := json.MarshalIndent(oldinfo, "", " ")
				fmt.Println(oldinfo)
				os.Remove(add + title + ext)
				os.Create(add + title + ext)
				ioutil.WriteFile(add+title+ext, data, 7)
			}

		}
	}
}

//GetUserallProjs requires authentication its incomplete till now
func GetUserallProjs(id int) []map[string]interface{} {
	var list []map[string]interface{}
	pres, _ := Find(id)
	if pres {
		userinfo := ReadUserId(id)
		project := convert(userinfo["projs"].([]interface{}))
		for _, projs := range project {
			list = append(list, projs)
		}
	}
	return list
}

//GetUseroneProjs requires authentication as well
func GetUseroneProjs(id int, pid float64) map[string]interface{} {
	var proj map[string]interface{}
	pres, _ := Find(id)
	if pres {
		userinfo := ReadUserId(id)
		if userinfo["projs"] != nil {
			project := convert(userinfo["projs"].([]interface{}))
			for _, projs := range project {
				if projs["pid"].(float64) == pid {
					proj = projs
				}
			}
		} else {

			fmt.Println("no project to display addone using connamd addproj first")
		}

	}
	if proj == nil {
		fmt.Println("given project with id is not presernt")
	} else {
		return proj
	}
	return nil
}

func Deleteuser(id int) {
	pres, _ := Find(id)

	if pres {
		title := strconv.Itoa(id)
		err := os.Remove(add + title + ext)
		if err != nil {
			fmt.Println("having error in removing the gven file :->", err)
		}
	} else {
		fmt.Println("can you destroy what you haven't created")
	}
}

func DropOneProject(id int, pid float64) {
	title := strconv.Itoa(id)
	UserInfo := ReadUserId(id)
	project := convert(UserInfo["projs"].([]interface{}))
	temp := make([]map[string]interface{}, 0)
	for i, projs := range project {
		if projs["pid"].(float64) == pid {
			temp = append(temp, project[:i]...)
			temp = append(temp, project[i+1:]...)
		}
	}
	UserInfo["projs"] = temp
	data, _ := json.MarshalIndent(UserInfo, "", " ")
	os.Remove(add + title + ext)
	os.Create(add + title + ext)
	ioutil.WriteFile(add+title+ext, data, 7)
}

func ChangeProjectStatus(id int, pid float64, status string) {
	title := strconv.Itoa(id)
	UserInfo := ReadUserId(id)
	project := convert(UserInfo["projs"].([]interface{}))
	for _, projs := range project {
		if projs["pid"].(float64) == pid {
			projs["completed"] = status
		}
	}
	UserInfo["projs"] = project
	data, _ := json.MarshalIndent(UserInfo, "", " ")
	os.Remove(add + title + ext)
	os.Create(add + title + ext)
	ioutil.WriteFile(add+title+ext, data, 7)
}

func GetUserProjectTimeDate(id int, pid float64) (int, int, int) {
	Project := GetUseroneProjs(id, pid)
	sdate := Project["startdate"].(string)
	ldate := Project["lastdate"].(string)
	sd, sm, sy := ut.Parse(sdate)
	ld, lm, ly := ut.Parse(ldate)
	return ld - sd, lm - sm, ly - sy
}

func Rawuser(m map[string]interface{}) mo.User {
	var u mo.User
	u.Name = m["name"].(string)
	u.Age = m["age"].(float64)
	u.ID = m["id"].(float64)
	u.Password = m["password"].(string)
	if m["projs"] != nil {
		u.Projs = Rawprojectarray(convert(m["projs"].([]interface{})))
	} else {
		// u.Projs = make([]map[string]interface{}, 0)
	}

	return u
}

func Rawprojectarray(ar []map[string]interface{}) []mo.Project {
	var pro mo.Project
	var res = make([]mo.Project, 0)
	for _, p := range ar {
		pro.Name = p["name"].(string)
		pro.ID = p["pid"].(float64)
		pro.Completed = p["completed"].(bool)
		pro.Startdate = p["startdate"].(string)
		pro.Lastdate = p["lastdate"].(string)
		res = append(res, pro)
	}
	return res
}

func Rawproject(p map[string]interface{})mo.Project{
	var pro mo.Project
	pro.Name = p["name"].(string)
	pro.ID = p["pid"].(float64)
	pro.Completed = p["completed"].(bool)
	pro.Startdate = p["startdate"].(string)
	pro.Lastdate = p["lastdate"].(string)
		return pro
}