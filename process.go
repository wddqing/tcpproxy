package tcpproxy


// Process 进程管理
type Process struct{
	Pid int
	PidFile string
}

func(p *Process) kill()  error{

	return nil
}

func(p *Process) storePidToFile() error{

	return nil
}

func(p *Process) getPidFromFile() int{

	return nil
}
