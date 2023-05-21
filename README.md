# helm restore

This is a basic plugin that allows you backup and restore whole cluster and specified namespace.

## Usage 

```sh
./src mode restore --namespace kube-system --sourcecluster kind-kind --targetcluster kind-kind-test-cluster
```

## TODO

* Backup only mode
* Better output
* Data Backup and Restore