package host_controller

import (
    // "github.com/gin-gonic/gin"
    "github.com/tatsushid/go-fastping"
    "github.com/sayden/docker-commander/entities/host"
    "net"
    "fmt"
    "time"
)

// func HostList(a *gin.Context) {}
// func HostDetail(a *gin.Context) {}

func isHostAlive(h *host.Host) bool{
  p := fastping.NewPinger()

  ra,err :=   net.ResolveIPAddr("ip4:icmp", h.Ip)
  if(err != nil){
    fmt.Println(err)
    return false
  }

  p.AddIPAddr(ra)
  p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
    fmt.Println("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
  }

  p.OnIdle = func() {
    fmt.Println("finish")
  }

  err = p.Run()
  if err != nil {
    fmt.Println(err)
    return false
  } else {
    return true
  }
}
