package dns

import (
	"context"
	"net"
	"strings"
	"time"
)

const DefaultDNSTimeout = 10 * time.Second

/*
 * Performs DNS resolution and returns single list of IP addresses,
 * where IPv6 addresses appear first.
 */
func LookupHost(domain string, timeout time.Duration) ([]string, error) {
	addrsIPv4, addrsIPv6, err := LookupHostSeparate(domain, timeout)
	if err != nil {
		return nil, err
	}

	for _, addrIPv4 := range addrsIPv4 {
		addrsIPv6 = append(addrsIPv6, addrIPv4)
	}

	return addrsIPv6, nil
}

/*
 * Performs DNS resolution and returns two separate lists of IPv4 and IPv6 addresses.
 */
func LookupHostSeparate(domain string, timeout time.Duration) ([]string, []string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	ipAddrs, err := net.DefaultResolver.LookupHost(ctx, domain)
	if err != nil {
		return nil, nil, err
	}

	var addrsIPv4, addrsIPv6 []string

	for _, ipAddr := range ipAddrs {
		if strings.Contains(ipAddr, ":") {
			addrsIPv6 = append(addrsIPv6, ipAddr)
		} else {
			addrsIPv4 = append(addrsIPv4, ipAddr)
		}
	}

	return addrsIPv4, addrsIPv6, nil
}
