// Code generated by gotestmd DO NOT EDIT.
package dns

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
	r := s.Runner("../deployments-k8s/examples/interdomain/dns")
	r.Run(`[[ ! -z $KUBECONFIG1 ]]`)
	r.Run(`[[ ! -z $KUBECONFIG2 ]]`)
	r.Run(`[[ ! -z $KUBECONFIG3 ]]`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`node1=$(kubectl get pods -n kube-system -l k8s-app=kube-dns -o go-template='{{index (index (index  .items 0) "spec") "nodeName"}}')`)
	r.Run(`ip1=$(kubectl get nodes $node1 -o go-template='{{range .status.addresses}}{{if eq .type "ExternalIP"}}{{.address}}{{end}}{{end}}')` + "\n" + `echo Selected node IP: ${ip1:=$(kubectl get nodes $node1 -o go-template='{{range .status.addresses}}{{if eq .type "InternalIP"}}{{.address}}{{end}}{{end}}')}`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`node2=$(kubectl get pods -n kube-system -l k8s-app=kube-dns -o go-template='{{index (index (index  .items 0) "spec") "nodeName"}}')`)
	r.Run(`ip2=$(kubectl get nodes $node2 -o go-template='{{range .status.addresses}}{{if eq .type "ExternalIP"}}{{.address}}{{end}}{{end}}')` + "\n" + `echo Selected node IP: ${ip2:=$(kubectl get nodes $node2 -o go-template='{{range .status.addresses}}{{if eq .type "InternalIP"}}{{.address}}{{end}}{{end}}')}`)
	r.Run(`export KUBECONFIG=$KUBECONFIG3`)
	r.Run(`node3=$(kubectl get pods -n kube-system -l k8s-app=kube-dns -o go-template='{{index (index (index  .items 0) "spec") "nodeName"}}')`)
	r.Run(`ip3=$(kubectl get nodes $node3 -o go-template='{{range .status.addresses}}{{if eq .type "ExternalIP"}}{{.address}}{{end}}{{end}}')` + "\n" + `echo Selected node IP: ${ip3:=$(kubectl get nodes $node3 -o go-template='{{range .status.addresses}}{{if eq .type "InternalIP"}}{{.address}}{{end}}{{end}}')}`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`kubectl apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/343015352e917b521c802fe07246a092b6f7b740/examples/interdomain/dns/service.yaml`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`kubectl apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/343015352e917b521c802fe07246a092b6f7b740/examples/interdomain/dns/service.yaml`)
	r.Run(`export KUBECONFIG=$KUBECONFIG3`)
	r.Run(`kubectl apply -f https://raw.githubusercontent.com/networkservicemesh/deployments-k8s/343015352e917b521c802fe07246a092b6f7b740/examples/interdomain/dns/service.yaml`)
	r.Run(`export KUBECONFIG=$KUBECONFIG1`)
	r.Run(`---` + "\n" + `cat > configmap.yaml <<EOF` + "\n" + `apiVersion: v1` + "\n" + `kind: ConfigMap` + "\n" + `metadata:` + "\n" + `  name: coredns` + "\n" + `  namespace: kube-system` + "\n" + `data:` + "\n" + `  Corefile: |` + "\n" + `    .:53 .:30053 {` + "\n" + `        errors` + "\n" + `        health {` + "\n" + `            lameduck 5s` + "\n" + `        }` + "\n" + `        ready` + "\n" + `        kubernetes cluster.local in-addr.arpa ip6.arpa {` + "\n" + `            pods insecure` + "\n" + `            fallthrough in-addr.arpa ip6.arpa` + "\n" + `            ttl 30` + "\n" + `        }` + "\n" + `        k8s_external my.cluster1` + "\n" + `        prometheus :9153` + "\n" + `        forward . /etc/resolv.conf {` + "\n" + `            max_concurrent 1000` + "\n" + `        }` + "\n" + `        loop` + "\n" + `        reload 5s` + "\n" + `    }` + "\n" + `    my.cluster2:53 {` + "\n" + `      forward . ${ip2}:30053` + "\n" + `    }` + "\n" + `    my.cluster3:53 {` + "\n" + `      forward . ${ip3}:30053` + "\n" + `    }` + "\n" + `EOF`)
	r.Run(`kubectl apply -f configmap.yaml`)
	r.Run(`export KUBECONFIG=$KUBECONFIG2`)
	r.Run(`cat > configmap.yaml <<EOF` + "\n" + `apiVersion: v1` + "\n" + `kind: ConfigMap` + "\n" + `metadata:` + "\n" + `  name: coredns` + "\n" + `  namespace: kube-system` + "\n" + `data:` + "\n" + `  Corefile: |` + "\n" + `    .:53 .:30053 {` + "\n" + `        errors` + "\n" + `        health {` + "\n" + `            lameduck 5s` + "\n" + `        }` + "\n" + `        ready` + "\n" + `        kubernetes cluster.local in-addr.arpa ip6.arpa {` + "\n" + `            pods insecure` + "\n" + `            fallthrough in-addr.arpa ip6.arpa` + "\n" + `            ttl 30` + "\n" + `        }` + "\n" + `        k8s_external my.cluster2` + "\n" + `        prometheus :9153` + "\n" + `        forward . /etc/resolv.conf {` + "\n" + `            max_concurrent 1000` + "\n" + `        }` + "\n" + `        loop` + "\n" + `        reload 5s` + "\n" + `    }` + "\n" + `    my.cluster1:53 {` + "\n" + `      forward . ${ip1}:30053` + "\n" + `    }` + "\n" + `    my.cluster3:53 {` + "\n" + `      forward . ${ip3}:30053` + "\n" + `    }` + "\n" + `EOF`)
	r.Run(`kubectl apply -f configmap.yaml`)
	r.Run(`export KUBECONFIG=$KUBECONFIG3`)
	r.Run(`cat > configmap.yaml <<EOF` + "\n" + `apiVersion: v1` + "\n" + `kind: ConfigMap` + "\n" + `metadata:` + "\n" + `  name: coredns` + "\n" + `  namespace: kube-system` + "\n" + `data:` + "\n" + `  Corefile: |` + "\n" + `    .:53 .:30053 {` + "\n" + `        errors` + "\n" + `        health {` + "\n" + `            lameduck 5s` + "\n" + `        }` + "\n" + `        ready` + "\n" + `        kubernetes cluster.local in-addr.arpa ip6.arpa {` + "\n" + `            pods insecure` + "\n" + `            fallthrough in-addr.arpa ip6.arpa` + "\n" + `            ttl 30` + "\n" + `        }` + "\n" + `        k8s_external my.cluster3` + "\n" + `        prometheus :9153` + "\n" + `        forward . /etc/resolv.conf {` + "\n" + `            max_concurrent 1000` + "\n" + `        }` + "\n" + `        loop` + "\n" + `        reload 5s` + "\n" + `    }` + "\n" + `    my.cluster1:53 {` + "\n" + `      forward . ${ip1}:30053` + "\n" + `    }` + "\n" + `    my.cluster2:53 {` + "\n" + `      forward . ${ip2}:30053` + "\n" + `    }` + "\n" + `EOF`)
	r.Run(`kubectl apply -f configmap.yaml`)
}
func (s *Suite) Test() {}
