package options

import (

	"github.com/spf13/pflag"
)

// APIServer runs a kubernetes api server.
type APIServer struct {
	Port                        int
	//Log-dir                     string
}

// NewAPIServer creates a new APIServer object with default parameters
func NewAPIServer() *APIServer {
	s := APIServer{
		Port:                66,
		//Log-dir:             "./",
	}
	return &s
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet
func (s *APIServer) AddFlags(fs *pflag.FlagSet) {
	pflag.IntVar(&s.Port, "Port", 8088, "api port")
	//pflag.StringVar(&s.log-dir, "Port", "./", "log path")
	
	pflag.Parse()
}
