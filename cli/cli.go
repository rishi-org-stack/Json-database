package main

import (
	"fmt"

	mp "github.com/rishi-org-stack/proj-track/cli/mapping"
)

func main() {
	var functions = mp.Util()
	fmt.Println("1)for searching a particulaur user ->search")
	fmt.Println("2)for reading a particulaur user ->readuser")
	fmt.Println("3)for adding a particulaur user ->adduser")
	fmt.Println("4)for searching a particulaur proj ->gsp")
	fmt.Println("5)for adding a particulaur proj ->addproj")
	fmt.Println("6)for listing a all proj ->list")
	fmt.Println("7)for removing a particulaur user ->remove")
	fmt.Println("8)for dropinig  a particulaur proj ->drop")
	fmt.Println("9)for changing status of a particulaur proj ->cstatus")
	fmt.Println("10)for getting time datee a particulaur proj ->td")
	fmt.Println("11)for exiting  ->stop")

	for {
		var what string
		fmt.Scan(&what)
		switch what {
		case "search":
			var id int
			fmt.Println("Id of user ->")
			fmt.Scan(&id)
			f := functions[what].(func(int)(bool,error))
			pres ,err:=f(id)
			if err!=nil{
				fmt.Println("error:",err)
			}else{
				fmt.Println(pres)
			}
		case "readuser":
			var id int
			fmt.Println("id you want to read")
			fmt.Scan(&id)
			f := functions[what].(func(int) map[string]interface{})
			fmt.Println(f(id))
		case "adduser":
			var name string
			var age float64
			var password string
			var id float64
			fmt.Println("name ->")
			fmt.Scan(&name)
			fmt.Println("age ->")
			fmt.Scan(&age)
			fmt.Println("strong password ->")
			fmt.Scan(&password)
			fmt.Println("unique Id->")
			fmt.Scan(&id)
			f := functions[what].(func(string, float64, string, float64))
			f(name, age, password, id)
		case "addproj":
			var id int
			var pid float64
			var name string
			var lastdate string
			var start string
			fmt.Println("your id")
			fmt.Scan(&id)
			fmt.Println("unique peoject id")
			fmt.Scan(&pid)
			fmt.Println("name ->")
			fmt.Scan(&name)
			fmt.Println("lastdate->")
			fmt.Scan(&lastdate)
			fmt.Println("startdate->")
			fmt.Scan(&start)
			f := functions[what].(func(int, float64, bool,string, string, string))
			f(id, pid, false,name, lastdate, start)
		case "list":
			var id int
			fmt.Println("Id of user")
			fmt.Scan(&id)
			f := functions["listprojects"].(func(int) []map[string]interface{})
			for _, val := range f(id) {
				fmt.Println(val)
			}
		case "gsp":
			var id int
			var pid float64
			fmt.Println("your id->")
			fmt.Scan(&id)
			fmt.Println("proj id->")
			fmt.Scan(&pid)
			f := functions["getspecificproject"].(func(int, float64) map[string]interface{})
			fmt.Println(f(id, pid))
		case "remove":
			var id int
			fmt.Println("id pls->")
			fmt.Scan(&id)
			f := functions["removeuser"].(func(int))
			f(id)
		case "drop":
			var id int
			fmt.Println("your id->")
			fmt.Scan(&id)
			var pid float64
			fmt.Println("project id")
			fmt.Scan(&pid)
			f := functions["dropproject"].(func(int, float64))
			f(id, pid)
		case "cstatus":
			var id int
			var pid float64
			var status string
			fmt.Println("your id")
			fmt.Scan(&id)
			fmt.Println("project id")
			fmt.Scan(&pid)
			fmt.Println("staus ->")
			fmt.Scan(&status)
			f := functions["changestatus"].(func(int, float64, string))
			f(id,pid,status)
		case "td":
			var id int
			var pid float64
			fmt.Println("your id->")
			fmt.Scan(&id)
			fmt.Println("your projrct id->")
			fmt.Scan(&pid)
			f:=functions["timedate"].(func(int,float64)(int,int,int))
			fmt.Println(f(id,pid))
		case "stop":
			fmt.Println("bye")
			return 
		}
	}

}
