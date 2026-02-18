package acmedns

import (
	"encoding/json"
	"net"
)

// cidrslice is a list of allowed cidr ranges
type Cidrslice []string

func (c *Cidrslice) JSON() string {
	ret, _ := json.Marshal(c.ValidEntries())
	return string(ret)
}

func (c *Cidrslice) IsValid() error {
	for _, v := range *c {
		_, _, err := net.ParseCIDR(sanitizeIPv6addr(v))
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Cidrslice) ValidEntries() []string {
	valid := []string{}
	for _, v := range *c {
		_, _, err := net.ParseCIDR(sanitizeIPv6addr(v))
		if err == nil {
			valid = append(valid, sanitizeIPv6addr(v))
		}
	}
	return valid
}

// AllowedFrom Check if IP belongs to an allowed net
func (c *Cidrslice) AllowedFrom(ip string) bool {
	remoteIP := net.ParseIP(ip)
	// Range not limited
	if len(c.ValidEntries()) == 0 {
		return true
	}
	for _, v := range c.ValidEntries() {
		_, vnet, _ := net.ParseCIDR(v)
		if vnet.Contains(remoteIP) {
			return true
		}
	}
	return false
}

// AllowedFromList Go through list (most likely from headers) to check for the IP.
// Reason for this is that some setups use reverse proxy in front of acme-dns
func (c *Cidrslice) AllowedFromList(ips []string) bool {
	if len(ips) == 0 {
		// If no IP provided, check if no whitelist present (everyone has access)
		return c.AllowedFrom("")
	}
	for _, v := range ips {
		if c.AllowedFrom(v) {
			return true
		}
	}
	return false
}
