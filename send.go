package main

import "fmt"

func (conn Connections) Send(msg *IPFIX, s string) {
	switch s {
	case "QOS":
		// Send only QOS stats with meaningful values
		if msg.Data.QOS.IncMos > 0 && msg.Data.QOS.OutMos > 0 {
			if *baddr != "" {
				conn.SendBanshee(msg, "QOS")
			}
			if *haddr != "" {
				conn.SendHep(msg, "QOS33")
				conn.SendHep(msg, "incQOS34")
				conn.SendHep(msg, "outQOS34")
				conn.SendHep(msg, "incQOS35")
				conn.SendHep(msg, "outQOS35")
			}
			if *saddr != "" {
				conn.SendStatsD(msg, "QOS")
			}
		}
		if *gaddr != "" {
			conn.SendLog(msg, "QOS")
		}
	default:
		if *haddr != "" {
			conn.SendHep(msg, "SIP")
		}
		if *gaddr != "" {
			conn.SendLog(msg, "SIP")
		}
		if *debug {
			fmt.Println("SIP output:")
			fmt.Printf("%s\n", msg.Data.SIP.SipMsg)
		}
	}
}

/*
// Template for a sync.Pool buffer
var buffers = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 65536)
	},
}
packet := buffers.Get().([]byte)
buffers.Put(packet)
*/
