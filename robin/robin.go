package robin

import (
	"lb"
	"sync"
)

type Robin struct {
	sync.Mutex

	next  int
	hosts []string
}

func New(hosts ...string) *Robin {
	return &Robin{next: 0, hosts: hosts}
}

func (rb *Robin) Add(host string) {
	rb.Lock()
	defer rb.Unlock()

	for _, h := range rb.hosts {
		if h == host {
			return
		}
	}
	rb.hosts = append(rb.hosts, host)
}

func (rb *Robin) Remove(host string) {
	rb.Lock()
	defer rb.Unlock()

	for i, h := range rb.hosts {
		if h == host {
			rb.hosts = append(rb.hosts[:i], rb.hosts[i+1:]...)
		}
	}
}

func (rb *Robin) Balance() (string, error) {
	rb.Lock()
	defer rb.Unlock()

	if len(rb.hosts) == 0 {
		return "", lb.ErrNoHost
	}

	host := rb.hosts[rb.next%len(rb.hosts)]
	rb.next++

	return host, nil
}
