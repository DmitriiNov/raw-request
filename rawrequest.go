package rawrequest

import (
	"fmt"
	"os"

	"github.com/DmitriiNov/raw-request/http"
)

func DoRequest() {
	if len(os.Args) != 2 {
		fmt.Println("Incorrect number of arguments")
		return
	}
	url := os.Args[1]
	http.MakeHttpRequest(url)
}
