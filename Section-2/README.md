# Section 2: Traffic Management and Routing

This section is all about the powerful traffic management and routing capabilities if Istio. You will learn about controlling inbound and outbound traffic into the mesh. Next, you will learn what service versions are and how to apply different routing rules between versions. You will see how you can apply this to implement rollout scenarios such as Blue/Green and Canary deployments.

- **Video 2.1**: [Controlling Ingress Traffic](video-2.1.md)
- **Video 2.2**: [Path and Header based Routing](video-2.2.md)
- **Video 2.3**: [Weight based Routing](video-2.3.md)
- **Video 2.4**: [Blue/Green and Canary Deployments](video-2.4.md)
- **Video 2.5**: [Controlling Egress Traffic](video-2.5.md)

## Prerequisites

This showcase assumes that you are running Kubernetes 1.9 or later and that you have
the sidecar injector webhook installed properly.

## Preparation

Optionally, create a dedicated namespace for this showcase and label it appropriately
for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled
$ kubectl label namespace default istio-injection=enabled
```

## Running

In this showcase are going to deploy two versions of the same microservice and
use different traffic management features to demonstrate the power and simplicity
of Istio.

```
$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')

$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ watch -n 1 -d http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

# apply the version subsets as destinations
$ kubectl apply -f hello-istio-destination.yaml

# apply the version specific virtual services
$ kubectl apply -f hello-istio-v1.yaml
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

$ kubectl apply -f hello-istio-v2.yaml
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

$ kubectl apply -f hello-istio-v1.yaml
$ kubectl apply -f hello-istio-75-25.yaml
$ kubectl apply -f hello-istio-50-50.yaml
$ kubectl apply -f hello-istio-25-75.yaml
$ kubectl apply -f hello-istio-v2.yaml

$ kubectl apply -f hello-istio-user-agent.yaml
$ http get $INGRESS_HOST/api/hello User-Agent:Chrome

$ kubectl apply -f hello-istio-user-cookie.yaml
$ http get $INGRESS_HOST/api/hello Cookie:user=oreilly
```
