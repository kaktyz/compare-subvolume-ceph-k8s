# Find persistent volumes from k8s in Ceph 

1) Take array of subvolumes in Ceph from [file](./subvolumeListFronCeph.txt)
2) Take array of Persistent Volumes if k8s cluster use your `~/.kube/config` file
3) Compare arrays and show in log array with diffirent values and array with equals values

You can specify global env while start `LOG_LEVEL` by default it INFO, but you can switch it to `DEBUG`.
You can use global env `KUBE_CONFIG_FILE_PATH` to declare full path to kubeconfig file, or dont use it and then will be use your default kubeconfig file.

## HTU
```bash
go build -mod=vendor -x -v
sudo chmod +x ./main
./main
```
