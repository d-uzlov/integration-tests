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
	r := s.Runner("../deployments-k8s/examples/k8s_monolith/dns")
	s.T().Cleanup(func() {
		r.Run(`kubectl delete service -n kube-system exposed-kube-dns`)
	})
	r.Run(`kubectl expose service kube-dns -n kube-system --port=53 --target-port=53 --protocol=TCP --name=exposed-kube-dns --type=LoadBalancer`)
	r.Run(`kubectl get services exposed-kube-dns -n kube-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "ip"}}'`)
	r.Run(`ipk8s=$(kubectl get services exposed-kube-dns -n kube-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "ip"}}')` + "\n" + `if [[ $ipk8s == *"no value"* ]]; then` + "\n" + `    $ipk8s=$(kubectl get services exposed-kube-dns -n kube-system -o go-template='{{index (index (index (index .status "loadBalancer") "ingress") 0) "hostname"}}')` + "\n" + `    $ipk8s=$(dig +short $ipk8s | head -1)` + "\n" + `fi` + "\n" + `echo Selected externalIP: $ipk8s` + "\n" + `[[ ! -z $ipk8s ]]`)
	r.Run(`ipdock=$(docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' nse-simple-vl3-docker)` + "\n" + `echo Selected dockerIP: $ipdock` + "\n" + `[[ ! -z $ipdock ]]`)
	r.Run(`cat > configmap.yaml <<EOF` + "\n" + `apiVersion: v1` + "\n" + `kind: ConfigMap` + "\n" + `metadata:` + "\n" + `  name: coredns` + "\n" + `  namespace: kube-system` + "\n" + `data:` + "\n" + `  Corefile: |` + "\n" + `    .:53 {` + "\n" + `        log` + "\n" + `        errors` + "\n" + `        health {` + "\n" + `            lameduck 5s` + "\n" + `        }` + "\n" + `        ready` + "\n" + `        kubernetes cluster.local in-addr.arpa ip6.arpa {` + "\n" + `            pods insecure` + "\n" + `            fallthrough in-addr.arpa ip6.arpa` + "\n" + `            ttl 30` + "\n" + `        }` + "\n" + `        k8s_external k8s.nsm` + "\n" + `        prometheus :9153` + "\n" + `        forward . /etc/resolv.conf` + "\n" + `        loop` + "\n" + `        reload 5s` + "\n" + `    }` + "\n" + `    docker.nsm:53 {` + "\n" + `        log` + "\n" + `        forward . ${ipdock}:53 {` + "\n" + `            force_tcp` + "\n" + `        }` + "\n" + `        reload 5s` + "\n" + `    }` + "\n" + `EOF`)
	r.Run(`kubectl apply -f configmap.yaml`)
	r.Run(`kubectl rollout restart -n kube-system deployment/coredns`)
	r.Run(`docker exec -d -i nse-simple-vl3-docker cp /etc/resolv.conf /etc/resolv_init.conf`)
	r.Run(`docker exec -d -i nse-simple-vl3-docker sh -c "echo 'nameserver 127.0.1.1' > /etc/resolv.conf"`)
	r.Run(`cat > coredns-config << EOF` + "\n" + `.:53 {` + "\n" + `    bind 127.0.1.1` + "\n" + `    log` + "\n" + `    errors` + "\n" + `    ready` + "\n" + `    file dnsentries.db` + "\n" + `    forward . /etc/resolv_init.conf {` + "\n" + `        max_concurrent 1000` + "\n" + `    }` + "\n" + `    loop` + "\n" + `    reload 5s` + "\n" + `}` + "\n" + `k8s.nsm:53 {` + "\n" + `    bind 127.0.1.1` + "\n" + `    log` + "\n" + `    forward . ${ipk8s}:53 {` + "\n" + `        force_tcp` + "\n" + `    }` + "\n" + `    reload 5s` + "\n" + `}` + "\n" + `EOF`)
	r.Run(`cat > dnsentries.db << EOF` + "\n" + `@       3600 IN SOA docker.nsm. . (` + "\n" + `                                2017042745 ; serial` + "\n" + `                                7200       ; refresh (2 hours)` + "\n" + `                                3600       ; retry (1 hour)` + "\n" + `                                1209600    ; expire (2 weeks)` + "\n" + `                                3600       ; minimum (1 hour)` + "\n" + `                                )` + "\n" + `spire-server.spire.docker.nsm   IN      A    ${ipdock}` + "\n" + `EOF`)
	r.Run(`docker cp coredns-config nse-simple-vl3-docker:/`)
	r.Run(`docker cp dnsentries.db nse-simple-vl3-docker:/`)
	r.Run(`docker exec -d nse-simple-vl3-docker coredns -conf coredns-config`)
}
func (s *Suite) Test() {}