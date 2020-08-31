# Initialize the Helm Chart

**This is a basic steps for installing & initializing a helm chart**

```sh
# Installing helm
wget https://get.helm.sh/helm-v2.16.10-linux-amd64.tar.gz
tar -xvzf helm-v2.16.10-linux-amd64.tar.gz
cd linux-amd64/
sudo mv helm /usr/local/bin/
which helm
helm version
helm version --short --client
kubectl -n kube-system create serviceaccount tiller
kubectl create clusterrolebinding tiller --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl -n kube-system get sa,clusterrolebindings |grep -i tiller

# Initializing helm 
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
### Helm Migrating from v2 to v3 
```sh
# Install a Release to test during migration
helm install stable/phpmyadmin --name phpmyadmin
helm list

helm install stable/phpmyadmin --name phpmyadmin
tar xzf helm-v3.3.0-linux-amd64.tar.gz
sudo mv helm /usr/local/bin/helm3
which helm3
helm3 version --short
helm repo list
helm3 repo list
helm list
helm3 list
helm3 plugin install https://github.com/helm/helm-2to3
helm3 plugin list

helm3 2to3 move --help
helm3 2to3 move config --dry-run
helm3 2to3 move config

helm3 2to3 convert --help
helm3 2to3 convert phpmyadmin --dry-run
helm3 2to3 convert phpmyadmin

helm3 2to3 cleanup
sudo rm -rf /usr/local/bin/helm
sudo mv /usr/local/bin/{helm3,helm}

which helm
helm verison --short
helm search repo phpmyadmin
helm install phpmyadmin stable/phpmyadmin
watch kubectl get all 
```
