package wrobin

import "sync"

type Wrobin struct {
	sync.Mutex
	hosts []string
	next  int
}


func New(hosts ...string) *Wrobin {
	return &Wrobin{next: 0, hosts: hosts}
}

// TODO: implement weighted robin algorithm
