package main

import (
	"net"
	"fmt"
	"os"
	"io"
)

func main() {
	if(len(os.Args) != 3){
		fmt.Println()
	}
	localAddr := os.Args[1]
	remoteAddr := os.Args[2]

	local,err := net.Listen("tcp",localAddr)

	if local == nil{
		fmt.Println(err)
	}
	for{
		conn,err := local.Accept()
		if conn == nil{
			fmt.Println(err)
		}
		go forward(conn,remoteAddr)
	}

}

func forward(local net.Conn,remoteAddr string){
	remote,err := net.Dial("tcp",remoteAddr)
	if remote == nil{
		fmt.Println(err)
		return
	}
	go io.Copy(local,remote)
	go io.Copy(remote,local)
}
