# Initialize the Helm Chart

** This is a basic steps for installing & initializing a helm chart **

```sh
wget https://get.helm.sh/helm-v2.16.10-linux-amd64.tar.gz
tar -xvzf helm-v2.16.10-linux-amd64.tar.gz
cd linux-amd64/
sudo mv helm /usr/local/bin/
which helm
helm version
helm version --short --client
kubectl -n kube-system create serviceaccount tiller
kubectl create clusterrolebinding tiller --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl -n kube-system get sa tiller
kubectl get clusterrolebindings tiller
helm init --service-account tiller
helm help | less
helm home
ls .helm/
helm list
helm search jenkins
helm help imspect
helm inspect stable/jenkins
helm fetch stable/jenkins
tar -xvzf jenkins-2.5.2.tgz
cd jenkins/
ls -l templates/
rm -rf jenkins jenkins-2.5.2.tgz
kubectl -n kube-system get deploy,rs,po,sa,clusterrolebinding | grep -i tiller
helm help reset
helm reset --remove-helm-home
kubectl -n kube-system get deploy,rs,po,sa,clusterrolebinding | grep -i tiller
kubectl -n kube-system delete clusterrolebindings tiller
kubectl -n kube-system delete sa tiller
```
### Creating your own chart 
```sh
helm create mychart
tree mychart/
```
