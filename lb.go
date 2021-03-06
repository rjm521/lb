package lb

import "errors"

type Balancer interface {
	New(hosts ...string)
	Add(host string)
	Remove(host string)
	Balance() (string, error)
}

var ErrNoHost = errors.New("host not found")
