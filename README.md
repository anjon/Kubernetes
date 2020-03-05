# Kubernetes
Kubernetes Knowledge

### Install Kubernete's Master (RHEL8)
To install kubernetes master on redhat8 we need to follow the below steps 

**Step-01:** Prepare the /etc/hosts file on the master node.
```sh
cat <<EOF>> /etc/hosts
192.168.1.101	k8s-master	master
192.168.1.102	k8s-node01	node01
192.168.1.103	k8s-node02	node02
EOF
```
**Step-02:** Disable the SELINUX and swap in the host.
```sh
setenforce 0
sed -i --follow-symlinks 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/sysconfig/selinux

swapoff -a 
sudo sed -i.bak '/ swap / s/^\(.*\)$/#\1/g' /etc/fstab
```
**Step-03:** Add the bridge netfilter module.
```sh
echo "br_netfilter" >> /etc/modules-load.d/br_netfilter.conf
modprobe br_netfilter
echo "net.bridge.bridge-nf-call-ip6tables = 1">> /etc/sysctl.d/01-custom.conf
echo "net.bridge.bridge-nf-call-iptables = 1">> /etc/sysctl.d/01-custom.conf
echo "net.bridge.bridge-nf-call-arptables = 1" >> /etc/sysctl.d/01-custom.conf
sysctl -p /etc/sysctl.d/01-custom.conf
```
**Step-04:** Update the filewall rule to open required port.
```sh
cat <<EOF>> /usr/lib/firewalld/services/k8s-master.xml
<?xml version="1.0" encoding="utf-8"?>
<service>
  <short>k8s-master</short>
  <description>Kubernets Master Firewall</description>
  <port protocol="tcp" port="6443"/>
  <port protocol="tcp" port="2379"/>
  <port protocol="tcp" port="2380"/>
  <port protocol="tcp" port="10250"/>
  <port protocol="tcp" port="10251"/>
  <port protocol="tcp" port="10252"/>
  <port protocol="tcp" port="10255"/>
</service>
EOF
firewall-cmd --add-service=k8s-master
firewall-cmd --add-service=k8s-master --permanent
```
**Step-05:** Install container runtime which is docker.
```sh
dnf install -y yum-utils device-mapper-persistent-data lvm2
dnf install -y https://download.docker.com/linux/centos/7/x86_64/stable/Packages/containerd.io-1.2.6-3.3.el7.x86_64.rpm
dnf config-manager --add-repo=https://download.docker.com/linux/centos/docker-ce.repo
dnf install -y docker-ce
```
**Step-06:** Install Kubeadm, Kubelet and Kubectl
```sh
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://packages.cloud.google.com/yum/repos/kubernetes-el7-x86_64
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://packages.cloud.google.com/yum/doc/yum-key.gpg https://packages.cloud.google.com/yum/doc/rpm-package-key.gpg
EOF
dnf install iproute-tc kubeadm kubelet kubectl kubernetes-cni -y
```
**Step-07:** Configure some additional setting for docker driver to be used for kubernetes
```sh
cat <<EOF > /etc/docker/daemon.json
{
  "exec-opts": ["native.cgroupdriver=systemd"],
  "log-driver": "json-file",
  "log-opts": {
     "max-size": "100m"
  },
  "storage-driver": "overlay2",
  "storage-opts": [
     "overlay2.override_kernel_check=true"
  ]
}
EOF

echo "KUBELET_EXTRA_ARGS=--cgroup-driver=systemd" > /etc/sysconfig/kubelet
```
**Step-08:** Initialize docker, kubelet and kubeadm
```sh
systemctl daemon-reload
systemctl restart docker && systemctl enable docker
systemctl restart kubelet && systemctl enable kubelet

kubeadm init --apiserver-advertise-address=$(hostname -i)
```
