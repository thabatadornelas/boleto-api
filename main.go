package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os/signal"
	"runtime"
	"strconv"
	"syscall"

	"os"

	"bitbucket.org/mundipagg/boletoapi/app"
	"bitbucket.org/mundipagg/boletoapi/config"
	"bitbucket.org/mundipagg/boletoapi/log"
)

var (
	processID  = os.Getpid()
	totalProcs = runtime.NumCPU()
	devMode    = flag.Bool("dev", false, "-dev To run in dev mode")
	mockMode   = flag.Bool("mock", false, "-mock To run mock requests")
	disableLog = flag.Bool("nolog", false, "-nolog disable seq log")
)

func init() {
	createPIDfile()
	// Configure signal handler
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go handleSignal(c)
}

func handleSignal(c chan os.Signal) {
	<-c
	config.Stop()
	log.Info("Quiting BoletoApi")
	log.Close()
	//db.GetDB().Close()
	fmt.Println("Done")
	os.Exit(1)
}

func createPIDfile() {
	p := strconv.Itoa(processID)
	ioutil.WriteFile("boletoapi.pid", []byte(p), 0644)
}

func main() {
	flag.Parse()
	logo1()
	app.Run(*devMode, *mockMode, *disableLog)
}

func logo1() {
	l := `
$$$$$$$\            $$\            $$\                $$$$$$\            $$\ 
$$  __$$\           $$ |           $$ |              $$  __$$\           \__|
$$ |  $$ | $$$$$$\  $$ | $$$$$$\ $$$$$$\    $$$$$$\  $$ /  $$ | $$$$$$\  $$\ 
$$$$$$$\ |$$  __$$\ $$ |$$  __$$\\_$$  _|  $$  __$$\ $$$$$$$$ |$$  __$$\ $$ |
$$  __$$\ $$ /  $$ |$$ |$$$$$$$$ | $$ |    $$ /  $$ |$$  __$$ |$$ /  $$ |$$ |
$$ |  $$ |$$ |  $$ |$$ |$$   ____| $$ |$$\ $$ |  $$ |$$ |  $$ |$$ |  $$ |$$ |
$$$$$$$  |\$$$$$$  |$$ |\$$$$$$$\  \$$$$  |\$$$$$$  |$$ |  $$ |$$$$$$$  |$$ |
\_______/  \______/ \__| \_______|  \____/  \______/ \__|  \__|$$  ____/ \__|
                                                               $$ |          
                                                               $$ |          
                                                               \__|          
	`
	fmt.Println(l)
	fmt.Println("Version: " + config.Get().Version)
}
