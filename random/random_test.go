package random

import (
	"log"
	"testing"
)

func TestNewRandom(t *testing.T) {
	hosts := []string{"192.168.0.1","127.0.0.1","1.88.2.88", "1.2.3.4"}

	req := 100

	lb := New(hosts...)
	loads := make(map[string]int)

	for i := 0; i <req * len(hosts); i++{
		host, _ := lb.Balance()
		loads[host]++
	}

	log.Println(loads)
}
