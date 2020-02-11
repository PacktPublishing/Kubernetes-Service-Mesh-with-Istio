# Section 4: Chaos Testing

Your services need to react gracefully in the face of failure. In this section, you will learn how to perform chaos monkey tests with Istio. You will see how to inject faults into communication channels. You will learn how envoy filters can be used to maliciously modify requests and responses and how traffic mirroring can be performed.

- **Video 4.1**: [Injecting HTTP Delay Faults](video-4.1.md)
- **Video 4.2**: [Injecting HTTP Abort Faults](video-4.2.md)
- **Video 4.3**: [Using Envoy Filters](video-4.3.md)
- **Video 4.4**: [Perform Traffic Mirroring](video-4.4.md)

## Prerequisites

This section assumes that you have running Istio installation on your Kubernetes cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled
$ kubectl label namespace default istio-injection=enabled

$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
```
