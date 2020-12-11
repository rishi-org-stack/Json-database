package mapping

import (
	mo "github.com/rishi-org-stack/proj-track/model/json"
)

func Util()map[string]interface{}{
	var m = make(map[string]interface{})
	m["search"]=mo.Find
	m["readuser"]=mo.ReadUserId
	m["adduser"]= mo.Insert
	m["addproj"]= mo.AddProj
	m["listprojects"]=mo.GetUserallProjs
	m["getspecificproject"]= mo.GetUseroneProjs
	m["removeuser"]= mo.Deleteuser
	m["dropproject"]= mo.DropOneProject
	m["changestatus"]=mo.ChangeProjectStatus
	m["timedate"]= mo.GetUserProjectTimeDate
	return m
}
// m["search"] = mo.Find
