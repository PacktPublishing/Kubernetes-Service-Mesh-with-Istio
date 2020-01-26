# Section 2: Traffic Management and Routing

This section is all about the powerful traffic management and routing capabilities if Istio. You will learn about controlling inbound and outbound traffic into the mesh. Next, you will learn what service versions are and how to apply different routing rules between versions. You will see how you can apply this to implement rollout scenarios such as Blue/Green and Canary deployments.

- **Video 2.1**: [Controlling Ingress Traffic](video-2.1.md)
- **Video 2.2**: [Path and Header based Routing](video-2.2.md)
- **Video 2.3**: [Weight based Routing](video-2.3.md)
- **Video 2.4**: [Blue/Green and Canary Deployments](video-2.4.md)
- **Video 2.5**: [Controlling Egress Traffic](video-2.5.md)

## Prerequisites

This section assumes that you have running Istio installation on your Kubernetes cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl create namespace hello-istio
$ kubectl label namespace hello-istio istio-injection=enabled

$ kubectl label namespace default istio-injection=enabled
```

## Running

```
$ kubectl get svc istio-ingressgateway -n istio-system
$ export INGRESS_HOST=$(kubectl -n istio-system get service istio-ingressgateway -o jsonpath='{.status.loadBalancer.ingress[0].ip}')

# deploy sample application
$ kubectl apply -f kubernetes/hello-istio.yaml

# create ingress gateway and route traffic to microservices
$ kubectl apply -f kubernetes/hello-istio-gateway.yaml
$ kubectl apply -f kubernetes/hello-istio-virtual-service.yaml

$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ watch -n 1 -d http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

# apply the version subsets as destinations
$ kubectl apply -f kubernetes/hello-istio-destination.yaml

# apply path based routing
$ kubectl apply -f kubernetes/hello-istio-uri-match.yaml

$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ http get $INGRESS_HOST/api/v1/hello Host:hello-istio.cloud
$ http get $INGRESS_HOST/api/v2/hello Host:hello-istio.cloud

# apply header based routing
$ kubectl apply -f hello-istio-user-agent.yaml
$ http get $INGRESS_HOST/api/hello User-Agent:Chrome Host:hello-istio.cloud

$ kubectl apply -f hello-istio-user-cookie.yaml
$ http get $INGRESS_HOST/api/hello Cookie:user=packtpub Host:hello-istio.cloud

# apply weight based routing
$ kubectl apply -f kubernetes/hello-istio-75-25.yaml
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ kubectl apply -f kubernetes/hello-istio-50-50.yaml
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ kubectl apply -f kubernetes/hello-istio-25-75.yaml
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

# perform blue green release deployment
$ kubectl apply -f kubernetes/hello-istio-v1.yaml
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud
$ kubectl apply -f kubernetes/hello-istio-v2.yaml
$ http get $INGRESS_HOST/api/hello Host:hello-istio.cloud

# perform canary release deployment
$ kubectl apply -f kubernetes/hello-istio-v1.yaml
$ kubectl apply -f kubernetes/hello-istio-75-25.yaml
$ kubectl apply -f kubernetes/hello-istio-50-50.yaml
$ kubectl apply -f kubernetes/hello-istio-25-75.yaml
$ kubectl apply -f kubernetes/hello-istio-v2.yaml

# get the current egress mode
$ kubectl get configmap istio -n istio-system -o yaml | grep -o "mode: ALLOW_ANY"

# disable ALLOW_ANY egress mode
$ kubectl get configmap istio -n istio-system -o yaml | sed 's/mode: ALLOW_ANY/mode: REGISTRY_ONLY/g' | kubectl replace -n istio-system -f -

$ export SOURCE_POD=$(kubectl get pod -l app=hello-istio-console -o jsonpath={.items..metadata.name})

$ kubectl exec -it $SOURCE_POD -c console /bin/sh
$ curl -I https://www.google.com

$ kubectl apply -f kubernetes/hello-istio-egress.yaml
```
