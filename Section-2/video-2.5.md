# Controlling Egress Traffic

## Prerequisites

This video assumes that you have running Istio installation on your Kubernetes cluster.

Optionally, create a dedicated namespace for this showcase and label it appropriately for the sidecar injector webhook to work. Or simply use the default namespace.

```
$ kubectl label namespace default istio-injection=enabled
```

## Running

```
# deploy sample application
$ kubectl apply -f kubernetes/hello-istio.yaml

# get the current egress mode
$ kubectl get configmap istio -n istio-system -o yaml | grep -o "mode: ALLOW_ANY"

# disable ALLOW_ANY egress mode
$ kubectl get configmap istio -n istio-system -o yaml | sed 's/mode: ALLOW_ANY/mode: REGISTRY_ONLY/g' | kubectl replace -n istio-system -f -

$ export SOURCE_POD=$(kubectl get pod -l app=hello-istio-console -o jsonpath={.items..metadata.name})

$ kubectl exec -it $SOURCE_POD -c console /bin/sh
$ curl -I https://www.google.com

$ kubectl apply -f kubernetes/hello-istio-egress.yaml
```