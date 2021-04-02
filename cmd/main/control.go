package main

import (
	"github.com/urfave/cli"
	C "go-hentai/constant"
)

var (
	flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "disable_bwm,db",
			Usage:       "Can also be set from the client's settings page.Disables the bandwidth monitor.",
			Destination: &disableBwm,
		},
		cli.BoolFlag{
			Name:        "disable_logging,dl",
			Usage:       "Can also be set from the client's settings page.Disables the opening, creation, and writing to the log_out file.",
			Destination: &disableLogging,
		},
		cli.BoolFlag{
			Name:        "flush-logs,fl",
			Usage:       "Flushes the log to disk for every written line. This will increase I/O; mostly recommended for when the log is written to a ramdisk.",
			Destination: &flushlogs,
		},
		cli.IntFlag{
			Name:        "max_connections,mc",
			Usage:       "Sets the maximum number of connections the application can handle to conn.",
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
)
