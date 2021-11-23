package condition

import (
	"net"
)

// сравнение сетей и IP адреса
type NetworkCondition struct {
	ip *net.IPNet
}

func (condition *NetworkCondition) ParseArgs(args []string) ([]string, error) {
	if len(args) == 0 {
		return nil, ErrExpectedValue("network CIDR or IP")
	}
	var err error
	_, condition.ip, err = net.ParseCIDR(args[0])
	if err != nil {
		_, condition.ip, err = net.ParseCIDR(args[0] + "/32")
		if err != nil {
			return nil, err
		}
	}
	return args[1:], nil
}

func (condition *NetworkCondition) Check(value interface{}) bool {
	var i net.IP
	switch v := value.(type) {
	case [4]byte:
		i = net.IP(v[:])
	case []byte:
		i = net.IP(v)
	case net.IP:
		i = v
	default:
		return false
	}
	return condition.ip.Contains(i)
}
