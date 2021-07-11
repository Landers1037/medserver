package main
// webserver of med
// provide simple static file server

// here are two instance of server 1.gin 2.go http
// only 100 lines to build a good web service
import "fmt"

func main() {
	e := InitCMD()
	if e != nil {
		fmt.Printf("medserver exit with %s\n", e.Error())
	}
}
