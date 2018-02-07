package cloudstack

import (
	"log"
	"testing"
)

func TestListMpxLoadBalancers(t *testing.T) {
	println("TestListMpxLoadBalancers start")
	cs := NewClient("user api url", "user api-key", "user sec-key", false)
	// p := cs.MpxLoadBalancer.NewListMpxLoadBalancerParams("", "", "", "", "")
	p := cs.MpxLoadBalancer.NewListMpxLoadBalancerParams()

	r, err := cs.MpxLoadBalancer.ListMpxLoadBalancers(p)
	println(r.Count)
	println(r.MpxLoadBalancers[0].Zonename)

	if err != nil {
		log.Fatalf("로드밸런서 조회 호춣시에, 에러발생 : %s", err)
	}
}
