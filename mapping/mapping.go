package mapping

import (
	// "fmt"

	mp "github.com/rishi-org-stack/proj-track/cli/mapping"
)

func Function(what int)interface{} {
	var functions = mp.Util()
	var f interface{}

		switch what {
		case 1:
			f = functions["search"]
		case 2:
			f = functions["readuser"]
		case 3:
			f = functions["adduser"]
		case 4:
			
			f = functions["addproj"]
		case 5:
			f = functions["listprojects"]
			
		case 6:
			f = functions["getspecificproject"]
		case 7:
			
			f = functions["removeuser"]
		case 8:

			f = functions["dropproject"]
		case 9:

			f = functions["changestatus"]
		case 10:
			f=functions["timedate"]
		}

	return f

}
