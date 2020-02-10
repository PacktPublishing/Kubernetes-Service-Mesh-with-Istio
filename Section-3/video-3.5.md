# Rate Limiting

## Prerequisites

This video assumes that you have running Istio installation on your Kubernetes cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl label namespace default istio-injection=enabled

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')

$ kubectl apply -f kubernetes/hello-istio.yaml
$ kubectl apply -f kubernetes/hello-istio-gateway.yaml
$ kubectl apply -f kubernetes/hello-istio-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-istio-destination.yaml
$ kubectl apply -f kubernetes/hello-message-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-message-destination.yaml
```

## Running

First, make sure that policy enforcement is enabled. If policy enforcement is enabled (disablePolicyChecks is false), no further action is needed, otherwise
use `istioctl`.

```bash
$ kubectl -n istio-system get cm istio -o jsonpath="{@.data.mesh}" | grep disablePolicyChecks

$ istioctl experimental manifest apply --set values.global.disablePolicyChecks=false
```

Next, we apply the rate limiting configuration.
- Mixer Side
    - memquota _handler_ defines memquota adapter configuration
    - quota _instance_ defines how quota is dimensioned by Mixer
    - quota _rule_ defines which quota instance is dispatched to memquota handler
- Client Side
    - _QuotaSpec_ defines quota name and amount that the client should request
    - _QuotaSpecBinding_ conditionally associates QuotaSpec with one or more services.

```bash
$ kubectl apply -f kubernetes/hello-rate-limit-mixer.yaml
$ kubectl apply -f kubernetes/hello-rate-limit-client.yaml
```

Finally, run a watch command to see the rate limiting in action.

```bash
$ watch -n 1 -d http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```
