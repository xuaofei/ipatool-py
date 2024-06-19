package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"net/http"
	"os"
)

func runServer(c *cli.Context) error {
	http.HandleFunc("/requestTask", requestTaskHandler)
	http.HandleFunc("/request2FA", request2FAHandler)
	http.HandleFunc("/uploadVersionsInfo", uploadVersionsInfoHandler)
	http.HandleFunc("/requestDownloadVersionsInfo", uploadVersionsInfoHandler)

	_ = http.ListenAndServe(":80", nil)

	return nil
}

func main() {
	log.Printf("args:%v", os.Args)

	app := &cli.App{
		Name:   "AppStoreHistoryIpaTool",
		Usage:  "run as client",
		Action: runServer,
	}

	_ = app.Run(os.Args)
}
