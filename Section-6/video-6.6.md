# Mesh Visualization with Kiali

## Prerequisites

## Prerequisites

This video assumes that you have a running Istio installation on your Kubernetes cluster. Make sure you have installed the `demo` profile for Istio.

```bash
$ istioctl manifest apply --set profile=demo

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```

Also, make sure to deploy the demo files for this section.

```bash
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled

$ kubectl apply -f kubernetes/
```

## Running

Check that Kiali is running correctly and use `istioctl` to open the dashboard. Alternatively, use port-forwarding.

```bash
$ kubectl -n istio-system get service kiali

$ kubectl port-forward -n istio-system service kiali 20001
$ istioctl dashboard kiali

$ open http://localhost:20001/kiali

# get the username and password, which is admin/admin
$ kubectl get secret kiali -n istio-system -o json
$ echo "YWRtaW4=" | base64 --decode -i -

# generate some traffic
$ watch -n 1 -d http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
```
