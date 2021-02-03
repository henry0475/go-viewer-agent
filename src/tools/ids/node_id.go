package ids

import (
	"log"
	"net"
	"strconv"
	"strings"
)

// GetNodeID will return a number range from 0 to 1023
// It can be temp used for identifying a node with an unique id
func GetNodeID() (nodeID int64) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		// Cannot run this application because the IP addresses cannot be found
		log.Fatalln(err)
		return
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipStr := strings.ReplaceAll(ipnet.IP.String(), ".", "")
				ipInt, _ := strconv.ParseInt(ipStr, 10, 64)
				nodeID = ipInt % 1024
				return
			}
		}
	}

	return
}
