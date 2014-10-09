package filefetch

import (
	"fmt"
	"github.com/kr/fs"
	"github.com/pelletier/go-toml"
	"logger"
	"net"
	"net/http"
	"os"
)

var (
	FolderExist int
)

/*func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, "+req.URL.Query().Get(":name")+"!\n")
}*/

//This will read for the files of the respository

func ReadDir(w http.ResponseWriter, req *http.Request) {

	// 192.168.2.128

	//ip := strings.Split(req.RemoteAddr, ":")[0]

	config, err := toml.LoadFile("../../config/julie.toml")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Print("Working")
	}

	security := config.Get("security.accept").(string)

	ip, _, _ := net.SplitHostPort(req.RemoteAddr)

	if security == ip {

		params := req.URL.Query()
		folder := params.Get(":folder")

		ip, _, _ := net.SplitHostPort(req.RemoteAddr)

		logger.FilesView(folder)

		//	fmt.Print(req.Host)
		fmt.Println(ip + "\n\n")

		//fmt.Fprintf(w, "X-Forwarded-For :"+req.Header.Get("X-FORWARDED-FOR"))

		walker := fs.Walk("E:\\Dockerstack\\" + folder)

		for walker.Step() {
			if err := walker.Err(); err != nil {
				w.Write([]byte("No Folder Found!!!"))
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			w.Write([]byte(walker.Path() + "\n"))

		}

	} else {

		w.Write([]byte("Wrong Ip"))
	}
}

// This will Check whether the folder Exist or not

func CheckFolder(w http.ResponseWriter, req *http.Request) {

	params := req.URL.Query()
	folder := params.Get(":folder")

	if _, err := os.Stat(folder); os.IsNotExist(err) {

		logger.Critical("no such file or directory:")

		FolderExist = 0
		return
	}

	FolderExist = 1

	logger.FolderExist(folder)

}

func HelloWorldHandler() {

	fmt.Print("Hello World")
}
