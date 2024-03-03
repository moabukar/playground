package accBlock

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type Config struct {
	Accounts map[string][]string `json:"accounts,omitempty"` // Account names mapped to IP ranges
}

func CreateConfig() *Config {
	return &Config{}
}

// AccBlock represents the middleware instance with a next handler and the configuration
type AccBlock struct {
	next     http.Handler
	name     string
	accounts map[string][]net.IPNet // Processed account names mapped to their parsed IP ranges
}

// New creates a new instance of the AccBlock middleware
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	accounts := make(map[string][]net.IPNet)
	for accountName, ips := range config.Accounts {
		var ipNets []net.IPNet
		for _, ipRange := range ips {
			_, ipNet, err := net.ParseCIDR(ipRange)
			if err != nil {
				return nil, fmt.Errorf("invalid IP range %s for account %s: %v", ipRange, accountName, err)
			}
			ipNets = append(ipNets, *ipNet)
		}
		accounts[accountName] = ipNets
	}

	return &AccBlock{
		next:     next,
		name:     name,
		accounts: accounts,
	}, nil
}

func (a *AccBlock) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	sourceIP := net.ParseIP(strings.Split(req.RemoteAddr, ":")[0])

	allowed := false
	for _, ranges := range a.accounts {
		for _, ipRange := range ranges {
			if ipRange.Contains(sourceIP) {
				allowed = true
				break
			}
		}
		if allowed {
			break
		}
	}

	if allowed {
		a.next.ServeHTTP(rw, req)
	} else {
		reject(rw)
	}
}

func reject(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusUnauthorized)
	_, err := rw.Write([]byte("Unauthorized"))
	if err != nil {
		fmt.Printf("Error sending unauthorized response: %v", err)
	}
}
