package main

import (
	"flag"
	"fmt"
	"log"
	"logwatcher/syslogd"
	"net/http"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/jeromer/syslogparser"
)

const (
	VERSION    = "0.1"
	EVENTS_LOG = "events.log"
	ERROR_LOG  = "error.log"
)

var (
	errLog *log.Logger
	logger *log.Logger
	conf   *Config
)

var (
	flag_v, flag_debug, flag_help bool
)

type Config struct {
	Debug      bool
	GOMAXPROCS int
	Authen     bool
	Http_port  string
	Udp_port   string
	User_token []string
	Allow_ip   []string
}

func init() {
	fmt.Println("init.....")

	flag.BoolVar(&flag_v, "v", false, "Version")
	flag.BoolVar(&flag_debug, "debug", false, "Debug")
	flag.BoolVar(&flag_help, "h", false, "Help")
	flag.Parse()

	if flag_v {
		fmt.Println("Version:", VERSION)
		os.Exit(0)
	}

	if flag_help {
		flag.Usage()
		os.Exit(0)
	}

	//create error logger
	f0, err := os.OpenFile(ERROR_LOG, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		errLog.Fatal("[error]opening error file: %v", err)
	}
	errLog = log.New(f0, "", log.Lshortfile|log.LstdFlags)

	//create events logger
	f1, err := os.OpenFile(EVENTS_LOG, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		errLog.Fatal("[error]opening error file: %v", err)
	}
	logger = log.New(f1, "", log.LstdFlags)

	//Decode config.toml
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		errLog.Fatal(err)
	}

	if conf.Debug {
		fmt.Println(conf)
	}

	//Single instance
	Singleton()

	//CPU Max procs
	if conf.GOMAXPROCS > 0 {
		runtime.GOMAXPROCS(conf.GOMAXPROCS)
	}

}

func main() {

	finish := make(chan bool)

	//http listen server
	hs := httpsvr()
	go func() {
		http.ListenAndServe(":"+conf.Http_port, hs)
	}()
	fmt.Println("小艾同学们开工啰:我在听:http://" + conf.Http_port)

	//udp listen server
	channel := make(chan syslogparser.LogParts, 1)

	svr := syslogd.NewServer()
	svr.ListenUDP(":" + conf.Udp_port)
	svr.Start(channel)

	go func() {
		for {
			//logparts := <-channel
			//Logswitcher(logparts)
			Logswitcher(<-channel)
		}
	}()

	<-finish
}
