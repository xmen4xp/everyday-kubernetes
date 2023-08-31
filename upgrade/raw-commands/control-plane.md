k drain k8s-control --ignore-daemonsets
sudo apt-get update
sudo apt-get install -y --allow-change-held-packages kubeadm=1.27.2-00
sudo kubeadm upgrade plan
sudo kubeadm upgrade apply v1.27.2
sudo apt-get install -y --allow-change-held-packages kubelet=1.27.2-00 kubectl=1.27.2-00 
sudo systemctl daemon-reload
sudo systemctl restart kubelet
k uncordon k8s-control
---
kubectl drain k8s-worker1 --ignore-daemonsets --force
k uncordon k8s-worker1
