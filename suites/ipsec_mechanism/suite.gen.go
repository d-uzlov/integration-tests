// Code generated by gotestmd DO NOT EDIT.
package ipsec_mechanism

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/spire/single_cluster"
)

type Suite struct {
	base.Suite
	single_clusterSuite single_cluster.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.single_clusterSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/ipsec_mechanism")
	s.T().Cleanup(func() {
		r.Run(`WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl delete mutatingwebhookconfiguration ${WH}` + "\n" + `kubectl delete ns nsm-system`)
	})
	r.Run(`kubectl create ns nsm-system`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/ipsec_mechanism?ref=a9dc3ffc936b0e904d60df34c3d3c799482e92c3`)
	r.Run(`WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl wait --for=condition=ready --timeout=1m pod ${WH} -n nsm-system`)
}
func (s *Suite) TestKernel2IP2Kernel() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2IP2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2ip2kernel`)
	})
	r.Run(`kubectl create ns ns-kernel2ip2kernel`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Kernel2IP2Kernel?ref=a9dc3ffc936b0e904d60df34c3d3c799482e92c3`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2ip2kernel`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-kernel2ip2kernel`)
	r.Run(`NSC=$(kubectl get pods -l app=alpine -n ns-kernel2ip2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ns-kernel2ip2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl exec ${NSC} -n ns-kernel2ip2kernel -- ping -c 4 172.16.1.100`)
	r.Run(`kubectl exec ${NSE} -n ns-kernel2ip2kernel -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestKernel2IP2Memif() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Kernel2IP2Memif")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-kernel2ip2memif`)
	})
	r.Run(`kubectl create ns ns-kernel2ip2memif`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Kernel2IP2Memif?ref=a9dc3ffc936b0e904d60df34c3d3c799482e92c3`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-kernel2ip2memif`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-memif -n ns-kernel2ip2memif`)
	r.Run(`NSC=$(kubectl get pods -l app=alpine -n ns-kernel2ip2memif --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-memif -n ns-kernel2ip2memif --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`kubectl exec ${NSC} -n ns-kernel2ip2memif -- ping -c 4 172.16.1.100`)
	r.Run(`result=$(kubectl exec "${NSE}" -n "ns-kernel2ip2memif" -- vppctl ping 172.16.1.101 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
}
func (s *Suite) TestMemif2IP2Kernel() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Memif2IP2Kernel")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-memif2ip2kernel`)
	})
	r.Run(`kubectl create ns ns-memif2ip2kernel`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Memif2IP2Kernel?ref=a9dc3ffc936b0e904d60df34c3d3c799482e92c3`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-memif -n ns-memif2ip2kernel`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-memif2ip2kernel`)
	r.Run(`NSC=$(kubectl get pods -l app=nsc-memif -n ns-memif2ip2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ns-memif2ip2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`result=$(kubectl exec "${NSC}" -n "ns-memif2ip2kernel" -- vppctl ping 172.16.1.100 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
	r.Run(`kubectl exec ${NSE} -n ns-memif2ip2kernel -- ping -c 4 172.16.1.101`)
}
func (s *Suite) TestMemif2IP2Memif() {
	r := s.Runner("../deployments-k8s/examples/use-cases/Memif2IP2Memif")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete ns ns-memif2ip2memif`)
	})
	r.Run(`kubectl create ns ns-memif2ip2memif`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/use-cases/Memif2IP2Memif?ref=a9dc3ffc936b0e904d60df34c3d3c799482e92c3`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nsc-memif -n ns-memif2ip2memif`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-memif -n ns-memif2ip2memif`)
	r.Run(`NSC=$(kubectl get pods -l app=nsc-memif -n ns-memif2ip2memif --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-memif -n ns-memif2ip2memif --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`result=$(kubectl exec "${NSC}" -n "ns-memif2ip2memif" -- vppctl ping 172.16.1.100 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
	r.Run(`result=$(kubectl exec "${NSE}" -n "ns-memif2ip2memif" -- vppctl ping 172.16.1.101 repeat 4)` + "\n" + `echo ${result}` + "\n" + `! echo ${result} | grep -E -q "(100% packet loss)|(0 sent)|(no egress interface)"`)
}