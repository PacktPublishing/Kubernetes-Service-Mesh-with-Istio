# Section 5: Securing the Mesh

This section is about securing the service mesh at various levels. You will learn how to enable mutual TLS between microservices, and how to use different other available fine-grained authentication and access policies. You will also learn how to enable and configure RBAC authorization with the service mesh.

- **Video 5.1**: Security by Default - Zero Trust Networks
- **Video 5.2**: [Mutual TLS between services](video-5.2.md)
- **Video 5.3**: [Enabling Strict Mode](video-5.3.md)
- **Video 5.4**: [Authorization on Ingress Gateway](video-5.4.md)
- **Video 5.5**: [Authorization for HTTP traffic](video-5.5.md)
- **Video 5.6**: [Authorization with JWT](video-5.6.md)

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
