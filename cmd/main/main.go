package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli"
	C "go-hentai/constant"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	disableBwm           bool
	disableLogging       bool
	flushlogs            bool
	maxconnections       int
	port                 int
	rescancache          int
	skipFreeSpaceCheck   bool
	verifyCache          bool
	disableIpOriginCheck bool
	disableFloodControl  bool
)

func init() {

}
func main() {
	app := newHentaiApp()
	_ = app.Run(os.Args)
}

func newHentaiApp() *cli.App {
	app := cli.NewApp()
	app.Name = "Hentai@Home"
	app.Version = C.ClientVersion + " (Build " + strconv.Itoa(C.ClientBuild) + ")"
	app.Copyright = "Copyright (c) 2008-" + strconv.Itoa(time.Now().Year()) + ", E-Hentai.org - all rights reserved."
	app.Usage = "This software comes with ABSOLUTELY NO WARRANTY. This is free software, and you are welcome to modify and redistribute it under the GPL v3 license."

	app.Flags = flags
	app.Before = func(c *cli.Context) error {
		_, cancel := context.WithCancel(context.Background())
		sigChan := make(chan os.Signal, 2)
		signal.Notify(sigChan, syscall.SIGILL, syscall.SIGTERM)
		go func() {
			<-sigChan
			cancel()

		}()
		return nil
	}
	app.Action = func(c *cli.Context) error {
		if c.NumFlags() == 0 {
			return cli.ShowAppHelp(c)
		}
		fmt.Println("xxa")
		return func(c *cli.Context) error {
			return nil
		}(c)
	}
	return app
}
