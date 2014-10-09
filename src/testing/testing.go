package main

import (
	//"filefetch"
	"fmt"
	//"logger"
)

func init() {

	//logger.AppStart()
}

func main() {

	// logger.FilesView("E:\\Dockerstack\\")
	//filefetch.ReadDir("./")

	//filefetch.CheckFolder("E:\\Dockerstak\\")

	i := 1

	for i < 10 {

		if i%2 == 0 {
			fmt.Println(i, "Even")
		} else {

			fmt.Println(i, "odd")

		}
		i++
	}

}
