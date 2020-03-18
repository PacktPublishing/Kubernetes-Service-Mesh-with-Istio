# Authorization on Ingress Gateway

## Prerequisites

This video assumes that you have a running Istio installation on your Kubernetes cluster. Also, make sure you have followed the setup instructions from the previous videos.

```bash
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled

$ kubectl apply -f kubernetes/demo/
$ kubectl get all -n hello-istio

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

## Running

First, make sure that you can call the services without any applied `AuthorizationPolicy` and
that you have configured to the gateway to forward source IP addresses.

```bash
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

$ kubectl patch svc istio-ingressgateway -n istio-system -p '{"spec":{"externalTrafficPolicy":"Local"}}'
```

Next, find out your client IP address and issue the following commands to `DENY` or `ALLOW` any traffic entering the ingress gateway from your IP.

```bash
$ curl -s 'https://api.ipify.org?format=json'
$ export CLIENT_IP=$(curl -s 'https://api.ipify.org?format=text')

$ sed "s/<<CLIENT_IP>>/$CLIENT_IP/" kubernetes/hello-istio-gateway-policy-deny.yaml | kubectl apply -f -
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

$ sed "s/<<CLIENT_IP>>/$CLIENT_IP/" kubernetes/hello-istio-gateway-policy-allow.yaml | kubectl apply -f -
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

# do some cleanup
$ kubectl delete AuthorizationPolicy -n istio-system --all
```

