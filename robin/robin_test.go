package robin

import (
	"log"
	"testing"
)

func TestNewRobin(t *testing.T)  {
	hosts := []string{"192.168.0.1","127.0.0.1","1.88.2.88"}

	req := 100

	lb := New(hosts...)
	loads := make(map[string]int)

	for i := 0; i <req * len(hosts); i++{
		host, _ := lb.Balance()
		loads[host]++
	}

	for h, load := range loads {
		if load > req {
			t.Fatalf("host(%s) got overloaded %d > %d\n", h, load, req)
		}
		if load < req {
			t.Fatalf("host(%s) got underloaded %d < %d\n", h, load, req)
		}
	}
	log.Println(loads)

}
