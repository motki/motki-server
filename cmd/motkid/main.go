// Command motkid starts a motki application server.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/motki/core/log"
	"github.com/motki/motki-server/app"
)

var confPath = flag.String("conf", "config.toml", "Path to configuration file.")
var version = flag.Bool("version", false, "Display the application version.")

var Version = "dev"

// fatalf creates a default logger, writes the given message, and exits.
func fatalf(format string, vals ...interface{}) {
	logger := log.New(log.Config{Level: "error"})
	logger.Fatalf(format, vals...)
}

func main() {
	flag.Parse()
	if *version {
		fmt.Printf("%s %s. %s\n", os.Args[0], Version, "https://github.com/motki/motki-server")
		os.Exit(0)
	}
	conf, err := app.NewConfigFromTOMLFile(*confPath)
	if err != nil {
		fatalf("motkid: unable to initialize config: %s", err.Error())
	}
	env, err := app.NewWebEnv(conf)
	if err != nil {
		fatalf("motkid: unable to initialize application environment: %s", err.Error())
	}

	err = env.Scheduler.ScheduleAt(
		env.Scheduler.RepeatFuncEvery(
			env.Model.UpdateCorporationDataFunc(env.Logger),
			1*time.Hour),
		time.Now().Add(10*time.Second))
	if err != nil {
		env.Logger.Warnf("motkid: unable to schedule corporation data update worker: %s", err.Error())
	}

	go func() {
		err = env.Web.ListenAndServe()
		if err != nil {
			env.Logger.Warnf("motkid: http server returned with error: %s", err.Error())
		}
	}()
	go func() {
		err = env.Web.ListenAndServeTLS()
		if err != nil {
			env.Logger.Warnf("motkid: https server returned with error: %s", err.Error())
		}
	}()

	env.BlockUntilSignal(make(chan os.Signal, 1))
}
