package main                                                                                                                                                         

import (
    "log"
    "net"
)

func main() {

    laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:4040")
    if err != nil {
        log.Fatal(err)
    }

    ln, err := net.ListenUDP("udp", laddr)
    if err != nil {
        log.Fatal(err)
    }

    for {
        b := make([]byte, 1024)
        n, raddr, err := ln.ReadFromUDP(b)
        if err != nil {
            log.Println("read from err:", err)
        }

        log.Printf("mess:%s, %d, %s\n", b, n, raddr.String())

        //buf := append(b, []byte("OK")...)
        buf := []byte(string(b[:n]) + " OK!")
        log.Println("svr send:", string(buf))
        n, err = ln.WriteToUDP(buf, raddr)
        if err != nil {
            log.Println("svr send err: ", err)
        } else {
            log.Println("svr send ok: ", n)
        }
    }
}
