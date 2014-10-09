/*
*
*This is the Logging Package for the Dockerstack Julie Engine!!!
*
*
 */

package logger

import (
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("logger")

var format = "%{color}%{time:15:04:05.000000} %{level:.10s} %{id:03x}%{color:reset} %{message}"

//This is used when the application starts

func AppStart(port int64) {

	logging.SetFormatter(logging.MustStringFormatter(format))
	log.Debug("Application Started at Port %d", port)

}

// this will be used to present some of the Critical Log messages in the application

func Critical(message string) error {

	logging.SetFormatter(logging.MustStringFormatter(format))

	log.Critical(message)

	return nil
}

// This message is to represent the Files views inside a particular folder

func FilesView(folder string) {

	logging.SetFormatter(logging.MustStringFormatter(format))

	log.Notice("Files View Started for %q", folder)
}

// This message is to represent whether the folder is existing or not

func FolderExist(folder string) {

	logging.SetFormatter(logging.MustStringFormatter(format))

	log.Debug("Folder %q exist!!!!", folder)

}
