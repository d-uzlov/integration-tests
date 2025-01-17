// Code generated by gotestmd DO NOT EDIT.
package external_nsc

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/k8s_monolith/configuration/loadbalancer"
	"github.com/networkservicemesh/integration-tests/suites/k8s_monolith/external_nsc/dns"
	"github.com/networkservicemesh/integration-tests/suites/k8s_monolith/external_nsc/docker"
	"github.com/networkservicemesh/integration-tests/suites/k8s_monolith/external_nsc/spire"
)

type Suite struct {
	base.Suite
	loadbalancerSuite loadbalancer.Suite
	dockerSuite       docker.Suite
	dnsSuite          dns.Suite
	spireSuite        spire.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.loadbalancerSuite, &s.dockerSuite, &s.dnsSuite, &s.spireSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/k8s_monolith/external_nsc")
	s.T().Cleanup(func() {
		r.Run(`WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl delete mutatingwebhookconfiguration ${WH}` + "\n" + `kubectl delete ns nsm-system`)
	})
	r.Run(`kubectl create ns nsm-system`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/k8s_monolith/configuration/cluster?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl get services registry -n nsm-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "ip"}}'`)
}
func (s *Suite) TestKernel2Wireguard2Kernel() {
	r := s.Runner("../deployments-k8s/examples/k8s_monolith/external_nsc/usecases/Kernel2Wireguard2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2wireguard2kernel-monolith-nsc`)
	})
	r.Run(`kubectl create ns ns-kernel2wireguard2kernel-monolith-nsc`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/k8s_monolith/external_nsc/usecases/Kernel2Wireguard2Kernel?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2wireguard2kernel-monolith-nsc`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ns-kernel2wireguard2kernel-monolith-nsc --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`docker exec nsc-simple-docker ping -c4 172.16.1.100`)
	r.Run(`kubectl exec ${NSE} -n ns-kernel2wireguard2kernel-monolith-nsc -- ping -c 4 172.16.1.101`)
}
