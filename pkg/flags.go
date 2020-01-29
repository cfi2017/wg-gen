package pkg

var (
	Network    = "192.168.61.0"
	Gateway    = "192.168.61.1"
	Broadcast  = "192.168.61.255"
	Mask       = 24
	ConfigFile = "/etc/wireguard/wg0.conf"

	DNS       = []string{"192.168.61.1", "8.8.8.8"}
	PublicKey = ""
	Networks  = []string{"192.168.60.0/24", "192.168.61.0/24"}
	Endpoint  = "office.fossilo.com:51820"
)
