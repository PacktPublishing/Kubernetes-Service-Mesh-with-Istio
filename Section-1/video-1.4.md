# Installing Istio by Hand

In this video are going to install Istio by hand, by applying the
Kubernetes resources such as CRDs and component deployments.

## Prerequesites

You will need a running Kubernetes with enough resources. If you install
Istio locally in Docker Desktop make sure you followed the instructions
under https://istio.io/docs/setup/platform-setup/docker/

## Step 1: Setup and Verify

First, you need to download and setup the latest Istio release.
At the time of writing this is `1.3.0`.
```
$ curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.3.0 sh -

$ cd istio-1.3.0
$ export PATH=$PWD/bin:$PATH

$ istioctl verify-install
```

## Step 2: Installation

In this step we are installing the Istio CRDs as well as the components and
services for the Istio demo.

```
# install and bootstrap all the Istio CRDs
$ kubectl apply -f install/kubernetes/helm/istio-init/files/crd-10.yaml
$ kubectl apply -f install/kubernetes/helm/istio-init/files/crd-11.yaml
$ kubectl apply -f install/kubernetes/helm/istio-init/files/crd-12.yaml
$ kubectl apply -f install/kubernetes/helm/istio-init/files/crd-certmanager-10.yaml
$ kubectl apply -f install/kubernetes/helm/istio-init/files/crd-certmanager-11.yaml
$ kubectl get crds | grep 'istio.io' | wc -l

# install the istio demo profile
$ kubectl apply -f install/kubernetes/istio-demo.yaml

# Verifying the installation
$ kubectl get pods -n istio-system
$ kubectl get svc -n istio-system
```

## (Optional) Step 4: Uninstall Istio

In case you want to uninstall Istio, issue the following commands:
```
@kubectl delete -f istio-$(VERSION)install/kubernetes/istio-demo.yaml
@kubectl delete -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-10.yaml
@kubectl delete -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-11.yaml
@kubectl delete -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-12.yaml
@kubectl delete -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-certmanager-10.yaml
@kubectl delete -f istio-$(VERSION)/install/kubernetes/helm/istio-init/files/crd-certmanager-11.yaml
@kubectl label namespace default istio-injection-
```
