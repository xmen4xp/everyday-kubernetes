# Build - a - Kubernetes using Ansible

1. Need sudo access to remote machines
2. ssh keys should be pre-populated in the remote machines.

## Step 1: Fill in your setup info in file: inventory

Inventory is composed of two catogeries of nodes:

**master** - K8s master node

**nodes** - K8s worker node

<details><summary>Example</summary>

[master]

18.191.204.6

[nodes]

18.217.35.46

3.145.142.23
</details>

<details><summary>Verify</summary>

1. List all hosts in the inventory to verify all hosts are visible to ansible.

```
ansible -i inventory --list-hosts all
```

</details>

