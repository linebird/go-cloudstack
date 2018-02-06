package cloudstack

import (
	"log"
	"testing"
)

func TestListMpxLoadBalancers(t *testing.T) {
	println("TestListMpxLoadBalancers start")
	cs := NewClient("https://api.ucloudbiz.olleh.com/loadbalancer/v1/client/api", "P_0ChsGqxI6CWsKrL3_mAg83GxSaAJFczoxDO4jE1JmPhiYTEvF5f1q3tgf1jXCPViQgYtnfeBgZw5TDRJdSSA", "rLzgCnTQnAw3uHsvzzmWvxg1gKdtj7WSzdNwzynzla0obVkGU0ks1fW3ocNSvD9DwkP7dEf8urvNHrti2WfVkg", false)
	// p := cs.MpxLoadBalancer.NewListMpxLoadBalancerParams("", "", "", "", "")
	p := cs.MpxLoadBalancer.NewListMpxLoadBalancerParams()

	r, err := cs.MpxLoadBalancer.ListMpxLoadBalancers(p)
	println(r.Count)
	println(r.MpxLoadBalancers[0].Zonename)

	if err != nil {
		log.Fatalf("로드밸런서 조회 호춣시에, 에러발생 : %s", err)
	}
}
