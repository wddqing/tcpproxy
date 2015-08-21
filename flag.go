package tcpproxy

import "flag"

type Flag struct{
	Input_port int
	Input_addr string
	Input_multi	int
	Input_pid  	string
	Input_log  	bool
	Input_cmd  string
}

func (this *Flag)InitFlag()  {
		this.Input_port = flag.Int("p",8080,"proxy server port")
		this.Input_addr = flag.String("d","0.0.0.0","proxy server bind address")
		this.Input_multi = flag.Int("m",1,"cpu core will be used(default 1)")
		this.Input_pid  = flag.String("f","/var/run/tcpproxy.pid","pid file to store server's pid")
		this.Input_log  = flag.Bool("l",false,"open log(default:false)")
		this.Input_cmd  = flag.String("c","test","operation like start and stop")
}