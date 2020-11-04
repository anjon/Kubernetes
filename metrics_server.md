### Deploy Metrics Server using helm
```sh
wget https://get.helm.sh/helm-v3.4.0-linux-amd64.tar.gz
tar -xzvf helm-v3.4.0-linux-amd64.tar.gz
sudo mv linux-amd64/helm /usr/local/bin/helm
which helm
helm version --short
sudo kubectl completion bash >/etc/bash_completion.d/kubectl
helm repo add stable https://charts.helm.sh/stable
helm search repo stable
helm search repo metrics-server
helm show values stable/metrics-server > /tmp/metrics-server.values
vim /tmp/metrics-server.values
hostNetwork:
  enabled: true             <--- Change to true
args:
- --kubelet-insecure-tls    <--- Remove the braces and remove the insucure-tls hash

kubectl create ns metrics-server
helm install metrics-server stable/metrics-server --namespace metrics-server --values /tmp/metrics-server.values
helm list --all-namespaces
kubectl -n metrics-server logs metrics-server-77dcb9f49f-p25kh -f
kubectl top node
```
