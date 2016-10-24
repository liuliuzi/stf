package main

import (	
	"fmt"
	"os"
	"runtime"	

	"github.com/spf13/pflag"
	"github.com/golang/glog"
	"flag"

	"liuyaoting.io/stf/cmd/stf-apiserver/app"
	"liuyaoting.io/stf/cmd/stf-apiserver/app/options"
	"liuyaoting.io/stf/pkg/util"

)	

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	s := options.NewAPIServer()
	s.AddFlags(pflag.CommandLine)

	flag.Set("FLAGS_log_dir", "./log")
	flag.Set("FLAGS_log_level", "0")
	flag.Parse()
	util.InitLogs()
	defer util.FlushLogs()

	glog.Info("invalid ==================")
	
	if err := app.Run(s); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}