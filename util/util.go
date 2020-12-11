package utils

import (
	"fmt"
	"strconv"
)

// Parse need improvement works only for dd/mm format we need dd/mm/yyy
func Parse(s string) (int, int,int) {
	//dd/mm format
	m := ""
	d := ""
	y := ""
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '/' {
			count++
		}
	}
	if count > 2  {
		fmt.Println("sorry wrong format '/' m2")
	}
	if count<2{
		fmt.Println("sorry wrong format '/' l2")
	}
	if count == 2 && len(s)==8{
		d =s[:2]
		m =s[3:5]
		y= s[6:len(s)]
	}

	if count ==2 && len(s)==6{
		d =s[0:1]
		m =s[2:3]
		y= s[4:len(s)]
	}
	day,_:=strconv.Atoi(d)
	month,_:=strconv.Atoi(m)
	year,_:=strconv.Atoi(y)
	return day, month,year
}

//Isvalid checks if the given date is valid or not
func Isvalid(s string)bool{
	ans :=false
	dm:=make(map[int]int)
	dm[1] =31
	dm[2] =28
	dm[3] =31
	dm[4] =30
	dm[5] =31
	dm[6] =30
	dm[7] =31
	dm[8] =31
	dm[9] =30
	dm[10]=31
	dm[11]=30
	dm[12]=31
	// year,m,d:=time.Now().Date()
	gd,gm,gy:=Parse(s)
	if gm>=1&&gm<13&&gy>=20{

		if gd<=dm[gm]&&gd>0{
			ans =true
		}else{
			ans = false
			fmt.Println("Nt a valid date for given month")
		} 
	}else{
		ans = false 
		fmt.Println(gm) 
		fmt.Println(gd)
		fmt.Println(gy)
		fmt.Println("invalid month")
	}
	return ans 
}

func Dateconstruct(dd string,mm string,yy string)string{
	var date = dd+"/"+mm+"/"+yy
	return date
}