package options

import (

	"github.com/spf13/pflag"
)

// APIServer runs a kubernetes api server.
type AgentServer struct {
	Port                          int
	ServerIP                      string
	LocalIP                       string
	//Log-dir                     string
}

// NewAPIServer creates a new APIServer object with default parameters
func NewAgentServer() *AgentServer {
	s := AgentServer{
		Port:                8088,
		ServerIP:            "127.0.0.1",
		LocalIP:             "127.0.0.2",
	}
	return &s
}

// AddFlags adds flags for a specific APIServer to the specified FlagSet
func (s *AgentServer) AddFlags(fs *pflag.FlagSet) {
	pflag.IntVar(&s.Port, "port", s.Port, "server port")
	pflag.StringVar(&s.ServerIP, "server_ip", s.ServerIP, "Server IP")
	pflag.StringVar(&s.LocalIP, "local_ip", s.LocalIP, "Local IP")
	//pflag.StringVar(&s.log-dir, "Port", "./", "log path")
	pflag.Parse()
}
