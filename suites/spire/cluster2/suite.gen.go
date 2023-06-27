// Code generated by gotestmd DO NOT EDIT.
package cluster2

import (
	"github.com/stretchr/testify/suite"

	"github.com/networkservicemesh/integration-tests/extensions/base"
)

type Suite struct {
	base.Suite
}

func (s *Suite) SetupSuite() {
	parents := []interface{}{&s.Suite}
	for _, p := range parents {
		if v, ok := p.(suite.TestingSuite); ok {
			v.SetT(s.T())
		}
		if v, ok := p.(suite.SetupAllSuite); ok {
			v.SetupSuite()
		}
	}
	r := s.Runner("../deployments-k8s/examples/spire/cluster2")
	s.T().Cleanup(func() {
		r.Run(`kubectl --kubeconfig=$KUBECONFIG2 delete crd clusterspiffeids.spire.spiffe.io` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete crd clusterfederatedtrustdomains.spire.spiffe.io` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete validatingwebhookconfiguration.admissionregistration.k8s.io/spire-controller-manager-webhook` + "\n" + `kubectl --kubeconfig=$KUBECONFIG2 delete ns spire`)
	})
	r.Run(`[[ ! -z $KUBECONFIG2 ]]`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -k https://github.com/d-uzlov/deployments-k8s/examples/spire/cluster2?ref=41407ca07ccc0454f9708e7a5f82657b0f0c24fc`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-server`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 wait -n spire --timeout=1m --for=condition=ready pod -l app=spire-agent`)
	r.Run(`kubectl --kubeconfig=$KUBECONFIG2 apply -f https://raw.githubusercontent.com/d-uzlov/deployments-k8s/4a0ec48044729b965bd0a6f161f5fbf2aaf5e0b4/examples/spire/cluster2/clusterspiffeid-template.yaml`)
}
func (s *Suite) Test() {}
