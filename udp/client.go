package main                                                                                                                                                         

import (
    "log"
    "math/rand"
    "net"
    "strconv"
    "time"
)

func main() {
    taddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:4040")
    if err != nil {
        log.Fatal(err)
    }

    conn, err := net.DialUDP("udp", nil, taddr)
    if err != nil {
        log.Fatal(err)
    }

    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)

    for {
        msg := strconv.Itoa(r1.Intn(1000))
        buf := []byte(msg)

        log.Println("Sending > " + string(buf))
        n, err := conn.Write(buf)
        if err != nil {
            log.Println("write err", msg, err)
        } else {
            log.Println("send ok", msg, "len:", n)
        }

        buf = make([]byte, 1024)
        n, _, err = conn.ReadFromUDP(buf)
        if err != nil {
            log.Println("error on receiving:", err)
        } else {
            log.Println("received from server: ", string(buf), " len: ", n)
        }

        time.Sleep(2 * time.Second)
    }
}          
