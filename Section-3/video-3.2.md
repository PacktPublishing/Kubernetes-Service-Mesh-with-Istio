# Setting request timeouts

## Prerequisites

This video assumes that you have running Istio installation on your Kubernetes cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl label namespace default istio-injection=enabled

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

## Running

```
$ kubectl apply -f kubernetes/hello-istio.yaml
$ kubectl apply -f kubernetes/hello-istio-gateway.yaml
$ kubectl apply -f kubernetes/hello-istio-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-istio-destination.yaml

$ kubectl get all
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud sleep==3
```

Next, edit the virtual service definitions for `hello-istio` and the `hello-message` service to configure the timeouts.

```yaml
    # configure a 2s timeout
    timeout: 2s
```

Issue the following commands to apply and see the timeouts in action.
```bash
$ kubectl apply -f kubernetes/hello-istio-virtual-service.yaml
$ kubectl apply -f kubernetes/hello-message-virtual-service.yaml

$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud sleep==3

$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud sleep==1
```