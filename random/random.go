package random

import (
	"lb"
	"math/rand"
	"sync"
)

type Random struct {
	sync.Mutex
	hosts []string
}

func New(hosts ...string) *Random {
	return &Random{hosts: hosts}
}

func (rd *Random) Add(host string) {
	rd.Lock()
	defer rd.Unlock()

	for _, h := range rd.hosts {
		if h == host {
			return
		}
	}
	rd.hosts = append(rd.hosts, host)
}

func (rd *Random) Remove(host string) {
	rd.Lock()
	defer rd.Unlock()

	for i, h := range rd.hosts {
		if h == host {
			rd.hosts = append(rd.hosts[:i], rd.hosts[i+1:]...)
		}
	}
}

func (rd *Random) Balance() (string, error) {
	rd.Lock()
	rd.Unlock()

	if len(rd.hosts) == 0 {
		return "", lb.ErrNoHost

	}
	h := rd.hosts[rand.Intn(len(rd.hosts))]

	return h, nil
}
