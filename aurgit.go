package main

import (
	"flag"
	"fmt"
	"net/http"
	// "strings"
	// "sync"
)

var chttp = http.NewServeMux()

func main() {

	configFile := flag.String("config", "config.json", "the configuration file. Defaults to same aurgo_config in current dir")

	flag.Parse()

	app := ConfigureApplication(configFile)

	// cmd := Command{app.Log}

	// wg := new(sync.WaitGroup)
	// commands := []string{"git status"}
	// for _, str := range commands {
	// 	wg.Add(1)
	// 	go cmd.ExeCmd(str, wg)
	// }
	// wg.Wait()

	app.Log.Infoln("Starting Aurgit Server...")

	http.HandleFunc("/", StatusHandler)   // status page
	http.HandleFunc("/sync", SyncHandler) // sync page
	app.Log.Infoln("Listening on " + app.Port)
	http.ListenAndServe(":"+app.Port, nil)

}

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Server is running")
}

func SyncHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Performing sync")
}
