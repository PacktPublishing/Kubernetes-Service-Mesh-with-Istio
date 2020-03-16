# Installing Istio by Hand

In this video are going to install Istio by hand, by applying the
Kubernetes resources such as CRDs and component deployments.

## Prerequesites

You will need a running Kubernetes with enough resources. If you install
Istio locally in Docker Desktop make sure you followed the instructions
under https://istio.io/docs/setup/platform-setup/docker/

## Step 1: Setup and Verify

First, you need to download and setup the latest Istio release.
At the time of writing this is `1.3.4`, upgrade to latest version if desired.
```
$ curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.3.4 sh -

$ cd istio-1.3.4
$ export PATH=$PWD/bin:$PATH

$ istioctl verify-install
```

## Step 2: Installation

In this step we are installing the Istio CRDs as well as the components and
services for the Istio demo.

```
# install and bootstrap all the Istio CRDs
$ for i in install/kubernetes/helm/istio-init/files/crd*yaml; do kubectl apply -f $i; done
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
$ kubectl label namespace default istio-injection-
$ kubectl delete -f istio-$(VERSION)install/kubernetes/istio-demo.yaml
$ for i in install/kubernetes/helm/istio-init/files/crd*yaml; do kubectl delete -f $i; done
```
