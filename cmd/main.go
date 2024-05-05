package main

import (
	"fmt"

	"go-tomcat/internal/startup"
	"go-tomcat/servlet"
	"go-tomcat/webapps"
	"go-tomcat/webapps/test"
)

func main() {
	bs := startup.NewBootStrap(":8008")
	bs.SetMaxConnections(200)
	bs.RegisterServlets([]servlet.Servlet{&webapps.HelloServlet{}, &webapps.StaticResourceProcessor{}})
	bs.RegisterFilters([]servlet.Filter{&test.TestFilter{}})
	if err := bs.Start(); err != nil {
		fmt.Println("Start Fail ,err ", err)
	}
}