package main

import (
	ftp "github.com/hduplooy/ftpsync"
	"log"
)

func main() {
	err := ftp.SyncFolders(&ftp.Config{
		Host:       "ftphost",
		Port:       "21",
		Username:   "username",
		Password:   "password",
		LocalPath:  "D:/tmp",
		RemotePath: "/public",
	})
	if err != nil {
		log.Fatal(err)
	}

}
