package main

import (
	"context"
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
	app := cli.NewApp()
	app.Name = "Hentai@Home"
	app.Version = C.ClientVersion + " (Build " + strconv.Itoa(C.ClientBuild) + ")"
	app.Copyright = "Copyright (c) 2008-" + strconv.Itoa(time.Now().Year()) + ", E-Hentai.org - all rights reserved."
	app.Usage = "This software comes with ABSOLUTELY NO WARRANTY. This is free software, and you are welcome to modify and redistribute it under the GPL v3 license."

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "disable_bwm,db",
			Usage:       "(Can also be set from the client's settings page)\nDisables the bandwidth monitor, which prevents the client from using more upload speed than the value of the \"Maximum Burst Speed\" parameter. The H@H dispatcher will still respect the \"Maximum Burst Speed\" parameter so that the mean upload speed of the client won't exceed that value. However, the client might use as much upload speed as its connection provides in order to send files, which generally results in upload peaks.\nIn some cases, the bandwidth monitor is unable to use the whole maximum upload speed that has been set, which can result in an under-utilization of the client's upload speed (and might trigger false overload notifications). In such cases, it is better to disable the bandwidth monitor in order for the client to give better performances.",
			Destination: &disableBwm,
		},
		cli.BoolFlag{
			Name:        "disable_logging,dl",
			Usage:       "(Can also be set from the client's settings page)\nDisables the opening, creation, and writing to the log_out file. This will significantly reduce I/O for HDD but make network troubleshooting harder. Java errors will still be logged to log_err.",
			Destination: &disableLogging,
		},
		cli.BoolFlag{
			Name:        "flush-logs,fl",
			Usage:       "Flushes the log to disk for every written line. This will increase I/O; mostly recommended for when the log is written to a ramdisk.",
			Destination: &flushlogs,
		},
		cli.IntFlag{
			Name:        "max_connections,mc",
			Usage:       "Sets the maximum number of connections the application can handle to conn. The default maximum number of connections depends on the max burst speed parameter, as detailed above. In some rare circumstances the default value is too low and the maximum number of connections needs to be increased slightly in order to avoid triggering false overload notifications. Setting this value too high might have severe consequences on a client's performance and impact the whole H@H network badly. DO NOT CHANGE THIS VALUE UNLESS YOU ABSOLUTELY KNOW WHAT YOU ARE DOING!",
			Value:       0,
			Destination: &maxconnections,
		},
		cli.IntFlag{
			Name:        "port,p",
			Usage:       "Overrides the port set in the client's settings. ",
			Value:       0,
			Destination: &port,
		},
		cli.IntFlag{
			Name:        "rescan-cache,rc",
			Usage:       "Checks the cache for errors.",
			Destination: &rescancache,
		},
		cli.BoolFlag{
			Name:        "skip_free_space_check,sc",
			Usage:       "Disables the free space check so that the application won't err if the space left on the partition that contains the cache folder is less than the configured parameter (or 100MB if that parameter is set to a value less than 100MB).",
			Destination: &skipFreeSpaceCheck,
		},
		cli.BoolFlag{
			Name:        "verify_cache,vc",
			Usage:       "Forces a check of the entire cache. This is the same as --rescan-cache with the addition that the SHA-1 hash of each file is also checked; this can take a long time.",
			Destination: &verifyCache,
		},
		cli.BoolFlag{
			Name:        "disable-ip-origin-check,dic",
			Usage:       "Disables the requirement that RPC server requests come from a whitelisted IP. Using this is discouraged as it reduces security, but may allow the client to work in some common non-transparent proxy configurations.",
			Destination: &disableIpOriginCheck,
		},
		cli.BoolFlag{
			Name:        "disable-flood-control,dfc",
			Usage:       "Disables the rate limit on connections per IP address. May be necessary to use with the above trigger.",
			Destination: &disableFloodControl,
		},
		cli.StringFlag{
			Name:        "cache-dir",
			Destination: &C.Path.CacheDir,
		},
		cli.StringFlag{
			Name:        "data-dir",
			Destination: &C.Path.DataDir,
		},
		cli.StringFlag{
			Name:        "download-dir",
			Destination: &C.Path.DownloadDir,
		},
		cli.StringFlag{
			Name:        "log-dir",
			Destination: &C.Path.LogDir,
		},
		cli.StringFlag{
			Name:        "temp-dir",
			Destination: &C.Path.TempDir,
		},
	}
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
}
