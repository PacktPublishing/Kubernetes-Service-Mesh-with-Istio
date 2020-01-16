# Deploying Services to the Mesh

In this video are going to deploy our first application to the Istio service mesh.

```
$ kubectl create namespace istio-demo
$ kubectl label namespace istio-demo istio-injection=enabled
$ kubectl label namespace default istio-injection=enabled

$ kubectl apply -f kubernetes/nginx-app.yaml -n istio-demo
$ kubectl get pods -n istio-demo
$ kubectl describe pod <POD_NAME>

$ kubectl apply -f kubernetes/nginx-app-istio-gateway.yaml -n istio-demo
$ kubectl apply -f kubernetes/nginx-app-istio-virtual-service.yaml -n istio-demo

$ kubectl get svc istio-ingressgateway -n istio-system
$ curl -H "Host: nginx-app.demo" <EXTERNAL-IP>

```
