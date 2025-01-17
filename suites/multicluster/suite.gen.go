// Code generated by gotestmd DO NOT EDIT.
package multicluster

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
	"github.com/networkservicemesh/integration-tests/suites/multicluster/dns"
	"github.com/networkservicemesh/integration-tests/suites/multicluster/loadbalancer"
	"github.com/networkservicemesh/integration-tests/suites/multicluster/spire"
)

type Suite struct {
	base.Suite
	loadbalancerSuite loadbalancer.Suite
	dnsSuite          dns.Suite
	spireSuite        spire.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite, &s.loadbalancerSuite, &s.dnsSuite, &s.spireSuite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/multicluster")
	s.T().Cleanup(func() {
		r.Run(`export KUBECONFIG=$KUBECONFIG1` + "\n" + `WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl delete mutatingwebhookconfiguration ${WH}` + "\n" + `kubectl delete ns nsm-system`)
		r.Run(`export KUBECONFIG=$KUBECONFIG2` + "\n" + `WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl delete mutatingwebhookconfiguration ${WH}` + "\n" + `kubectl delete ns nsm-system`)
		r.Run(`export KUBECONFIG=$KUBECONFIG3 && kubectl delete ns nsm-system`)
	})
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl create ns nsm-system`)
	r.Run(`kubectl apply -k ./clusters-configuration/cluster1`)
	r.Run(`kubectl get services nsmgr-proxy -n nsm-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "ip"}}'`)
	r.Run(`WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl wait --for=condition=ready --timeout=1m pod ${WH} -n nsm-system`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl create ns nsm-system`)
	r.Run(`kubectl apply -k ./clusters-configuration/cluster2`)
	r.Run(`kubectl get services nsmgr-proxy -n nsm-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "ip"}}'`)
	r.Run(`WH=$(kubectl get pods -l app=admission-webhook-k8s -n nsm-system --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `kubectl wait --for=condition=ready --timeout=1m pod ${WH} -n nsm-system`)
	r.Run(`export KUBECONFIG=$KUBECONFIG3`)
	r.Run(`kubectl create ns nsm-system`)
	r.Run(`kubectl apply -k ./clusters-configuration/cluster3`)
	r.Run(`kubectl get services registry -n nsm-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "ip"}}'`)
}
func (s *Suite) TestFloating_Kernel2Vxlan2Kernel() {
	r := s.Runner("../deployments-k8s/examples/multicluster/usecases/floating_Kernel2Vxlan2Kernel")
	s.T().Cleanup(func() {
		r.Run(`export KUBECONFIG=$KUBECONFIG1`)
		r.Run(`kubectl delete ns ns-floating-kernel2vxlan2kernel`)
		r.Run(`export KUBECONFIG=$KUBECONFIG2`)
		r.Run(`kubectl delete ns ns-floating-kernel2vxlan2kernel`)
	})
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl create ns ns-floating-kernel2vxlan2kernel`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_Kernel2Vxlan2Kernel/cluster2?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-floating-kernel2vxlan2kernel`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ns-floating-kernel2vxlan2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `[[ ! -z $NSE ]]`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl create ns ns-floating-kernel2vxlan2kernel`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_Kernel2Vxlan2Kernel/cluster1?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-floating-kernel2vxlan2kernel`)
	r.Run(`NSC=$(kubectl get pods -l app=alpine -n ns-floating-kernel2vxlan2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `[[ ! -z $NSC ]]`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl exec ${NSC} -n ns-floating-kernel2vxlan2kernel -- ping -c 4 172.16.1.2`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl exec ${NSE} -n ns-floating-kernel2vxlan2kernel -- ping -c 4 172.16.1.3`)
}
func (s *Suite) TestFloating_Kernel2Wireguard2Kernel() {
	r := s.Runner("../deployments-k8s/examples/multicluster/usecases/floating_Kernel2Wireguard2Kernel")
	s.T().Cleanup(func() {
		r.Run(`export KUBECONFIG=$KUBECONFIG1`)
		r.Run(`kubectl delete ns ns-floating-kernel2wireguard2kernel`)
		r.Run(`export KUBECONFIG=$KUBECONFIG2`)
		r.Run(`kubectl delete ns ns-floating-kernel2wireguard2kernel`)
	})
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl create ns ns-floating-kernel2wireguard2kernel`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_Kernel2Wireguard2Kernel/cluster2?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-floating-kernel2wireguard2kernel`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ns-floating-kernel2wireguard2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `[[ ! -z $NSE ]]`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl create ns ns-floating-kernel2wireguard2kernel`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_Kernel2Wireguard2Kernel/cluster1?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-floating-kernel2wireguard2kernel`)
	r.Run(`NSC=$(kubectl get pods -l app=alpine -n ns-floating-kernel2wireguard2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `[[ ! -z $NSC ]]`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl exec ${NSC} -n ns-floating-kernel2wireguard2kernel -- ping -c 4 172.16.1.2`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl exec ${NSE} -n ns-floating-kernel2wireguard2kernel -- ping -c 4 172.16.1.3`)
}
func (s *Suite) TestFloating_vl3_basic() {
	r := s.Runner("../deployments-k8s/examples/multicluster/usecases/floating_vl3-basic")
	s.T().Cleanup(func() {
		r.Run(`export KUBECONFIG=$KUBECONFIG3 && kubectl delete -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-basic/cluster3?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
		r.Run(`export KUBECONFIG=$KUBECONFIG2 && kubectl delete -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-basic/cluster2?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
		r.Run(`export KUBECONFIG=$KUBECONFIG1 && kubectl delete -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-basic/cluster1?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	})
	r.Run(`export KUBECONFIG=$KUBECONFIG3`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-basic/cluster3?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-basic/cluster1?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-basic/cluster2?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-floating-vl3-basic`)
	r.Run(`nsc2=$(kubectl get pods -l app=alpine -n ns-floating-vl3-basic --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-floating-vl3-basic`)
	r.Run(`nsc1=$(kubectl get pods -l app=alpine -n ns-floating-vl3-basic --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`ipAddr2=$(kubectl --kubeconfig=$KUBECONFIG2 exec -n ns-floating-vl3-basic $nsc2 -- ifconfig nsm-1)` + "\n" + `ipAddr2=$(echo $ipAddr2 | grep -Eo 'inet addr:[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'| cut -c 11-)` + "\n" + `kubectl exec $nsc1 -n ns-floating-vl3-basic -- ping -c 4 $ipAddr2`)
	r.Run(`kubectl exec $nsc1 -n ns-floating-vl3-basic -- ping -c 4 172.16.0.0` + "\n" + `kubectl exec $nsc1 -n ns-floating-vl3-basic -- ping -c 4 172.16.1.0`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`ipAddr1=$(kubectl --kubeconfig=$KUBECONFIG1 exec -n ns-floating-vl3-basic $nsc1 -- ifconfig nsm-1)` + "\n" + `ipAddr1=$(echo $ipAddr1 | grep -Eo 'inet addr:[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'| cut -c 11-)` + "\n" + `kubectl exec $nsc2 -n ns-floating-vl3-basic -- ping -c 4 $ipAddr1`)
	r.Run(`kubectl exec $nsc2 -n ns-floating-vl3-basic -- ping -c 4 172.16.0.0` + "\n" + `kubectl exec $nsc2 -n ns-floating-vl3-basic -- ping -c 4 172.16.1.0`)
}
func (s *Suite) TestFloating_vl3_scale_from_zero() {
	r := s.Runner("../deployments-k8s/examples/multicluster/usecases/floating_vl3-scale-from-zero")
	s.T().Cleanup(func() {
		r.Run(`export KUBECONFIG=$KUBECONFIG3 && kubectl delete -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-scale-from-zero/cluster3?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
		r.Run(`export KUBECONFIG=$KUBECONFIG2 && kubectl delete -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-scale-from-zero/cluster2?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
		r.Run(`export KUBECONFIG=$KUBECONFIG1 && kubectl delete -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-scale-from-zero/cluster1?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	})
	r.Run(`export KUBECONFIG=$KUBECONFIG3`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-scale-from-zero/cluster3?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-scale-from-zero/cluster1?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/floating_vl3-scale-from-zero/cluster2?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-floating-vl3-scale-from-zero`)
	r.Run(`nsc2=$(kubectl get pods -l app=alpine -n ns-floating-vl3-scale-from-zero --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=alpine -n ns-floating-vl3-scale-from-zero`)
	r.Run(`nsc1=$(kubectl get pods -l app=alpine -n ns-floating-vl3-scale-from-zero --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`)
	r.Run(`ipAddr2=$(kubectl --kubeconfig=$KUBECONFIG2 exec -n ns-floating-vl3-scale-from-zero $nsc2 -- ifconfig nsm-1)` + "\n" + `ipAddr2=$(echo $ipAddr2 | grep -Eo 'inet addr:[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'| cut -c 11-)` + "\n" + `kubectl exec $nsc1 -n ns-floating-vl3-scale-from-zero -- ping -c 4 $ipAddr2`)
	r.Run(`kubectl exec $nsc1 -n ns-floating-vl3-scale-from-zero -- ping -c 4 172.16.0.0` + "\n" + `kubectl exec $nsc1 -n ns-floating-vl3-scale-from-zero -- ping -c 4 172.16.1.0`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`ipAddr1=$(kubectl --kubeconfig=$KUBECONFIG1 exec -n ns-floating-vl3-scale-from-zero $nsc1 -- ifconfig nsm-1)` + "\n" + `ipAddr1=$(echo $ipAddr1 | grep -Eo 'inet addr:[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}'| cut -c 11-)` + "\n" + `kubectl exec $nsc2 -n ns-floating-vl3-scale-from-zero -- ping -c 4 $ipAddr1`)
	r.Run(`kubectl exec $nsc2 -n ns-floating-vl3-scale-from-zero -- ping -c 4 172.16.0.0` + "\n" + `kubectl exec $nsc2 -n ns-floating-vl3-scale-from-zero -- ping -c 4 172.16.1.0`)
}
func (s *Suite) TestInterdomain_Kernel2Vxlan2Kernel() {
	r := s.Runner("../deployments-k8s/examples/multicluster/usecases/interdomain_Kernel2Vxlan2Kernel")
	s.T().Cleanup(func() {
		r.Run(`export KUBECONFIG=$KUBECONFIG1`)
		r.Run(`kubectl delete ns ns-interdomain-kernel2vxlan2kernel`)
		r.Run(`export KUBECONFIG=$KUBECONFIG2`)
		r.Run(`kubectl delete ns ns-interdomain-kernel2vxlan2kernel`)
	})
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl create ns ns-interdomain-kernel2vxlan2kernel`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/interdomain_Kernel2Vxlan2Kernel/cluster2?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ns-interdomain-kernel2vxlan2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `[[ ! -z $NSE ]]`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-interdomain-kernel2vxlan2kernel`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl create ns ns-interdomain-kernel2vxlan2kernel`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/interdomain_Kernel2Vxlan2Kernel/cluster1?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-interdomain-kernel2vxlan2kernel`)
	r.Run(`NSC=$(kubectl get pods -l app=alpine -n ns-interdomain-kernel2vxlan2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `[[ ! -z $NSC ]]`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl exec ${NSC} -n ns-interdomain-kernel2vxlan2kernel -- ping -c 4 172.16.1.2`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl exec ${NSE} -n ns-interdomain-kernel2vxlan2kernel -- ping -c 4 172.16.1.3`)
}
func (s *Suite) TestInterdomain_Kernel2Wireguard2Kernel() {
	r := s.Runner("../deployments-k8s/examples/multicluster/usecases/interdomain_Kernel2Wireguard2Kernel")
	s.T().Cleanup(func() {
		r.Run(`export KUBECONFIG=$KUBECONFIG1`)
		r.Run(`kubectl delete ns ns-interdomain-kernel2wireguard2kernel`)
		r.Run(`export KUBECONFIG=$KUBECONFIG2`)
		r.Run(`kubectl delete ns ns-interdomain-kernel2wireguard2kernel`)
	})
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl create ns ns-interdomain-kernel2wireguard2kernel`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/interdomain_Kernel2Wireguard2Kernel/cluster2?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl wait --for=condition=ready --timeout=1m pod -l app=nse-kernel -n ns-interdomain-kernel2wireguard2kernel`)
	r.Run(`NSE=$(kubectl get pods -l app=nse-kernel -n ns-interdomain-kernel2wireguard2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `[[ ! -z $NSE ]]`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl create ns ns-interdomain-kernel2wireguard2kernel`)
	r.Run(`kubectl apply -k https://github.com/networkservicemesh/deployments-k8s/examples/multicluster/usecases/interdomain_Kernel2Wireguard2Kernel/cluster1?ref=ef02275eec2ad9601b41fabf40de496dd6b8dac4`)
	r.Run(`kubectl wait --for=condition=ready --timeout=5m pod -l app=alpine -n ns-interdomain-kernel2wireguard2kernel`)
	r.Run(`NSC=$(kubectl get pods -l app=alpine -n ns-interdomain-kernel2wireguard2kernel --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')` + "\n" + `[[ ! -z $NSC ]]`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl exec ${NSC} -n ns-interdomain-kernel2wireguard2kernel -- ping -c 4 172.16.1.2`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl exec ${NSE} -n ns-interdomain-kernel2wireguard2kernel -- ping -c 4 172.16.1.3`)
}
