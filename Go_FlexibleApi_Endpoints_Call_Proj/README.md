# Go_FlexibleApi_Endpoints_Call_Project
Go_Api_call_project ,a flexible Api call project that allows for different endpoints from different web services in a single project

## Calling multple endpoints sequentially

package main
import (
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)
var apis map[int]string
func fetchData(API int) {
    url := apis[API]
    if resp, err := http.Get(url); err == nil {
        defer resp.Body.Close()
        if body, err := io.ReadAll(resp.Body);
            err == nil {
            var result map[string]interface{}
            json.Unmarshal([]byte(body), &result)
            switch API {
            case 1:   // for the data.fixer.io/api/API
                if result["success"] == true {
                    fmt.Println(result["rates"].(
                        map[string]interface{})["USD"])
                } else {
                    fmt.Println(result["error"].(
                        map[string]interface{})["info"])
}
            case 2:  // for the openweathermap.org API
                if result["main"] != nil {
                    fmt.Println(result["main"].(
                        map[string]interface{})["temp"])
                } else {
                    fmt.Println(result["message"])
}
}
} else {
            log.Fatal(err)
        }
    } else {
        log.Fatal(err)


		}
	

	}
func main() {
    apis = make(map[int]string)
apis[1] = "http://data.fixer.io/api/latest?access_key=" + "<access_key>"
apis[2] = "http://api.openweathermap.org/data/2.5/weather?" + "q=SINGAPORE&appid=<access-key>"
    fetchData(1)
    fetchData(2)
}



## Fetching from multiple web services
at the same time
One of the strengths of Go is its support for concurrency. Instead of calling the two web services sequentially (one after the other), why not call them concurrently (at the same time)?
To do so, you just need to prefix each call to the fetchData() function with the go keyword:

func main() {
    apis = make(map[int]string)
    apis[1] =
"http://data.fixer.io/api/latest?access_key=" + "<access_key>" 

apis[2] =
"http://api.openweathermap.org/data/2.5/weather?" + "q=SINGAPORE&appid=<api_key>"

go fetchData(1) 
go fetchData(2)
    fmt.Scanln()
}

## he output of the above codes when 
go run main.go  

command was run is: 

298.08
1.073555