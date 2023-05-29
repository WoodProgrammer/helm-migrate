# helm migrate

This is a basic plugin that allows you backup and restore whole cluster and specified namespace.


## Build & Installation
The installation part is very simple you can easily install with via native helm commands like this;

* Installation
```sh
    git clone git@github.com:WoodProgrammer/helm-migrate.git
    
    pushd helm-migrate
        wget https://github.com/WoodProgrammer/helm-migrate/releases/download/0.0.1/helm-migrate_0.0.2_OS_amd64.tar.gz
        tar -xvf helm-migrate_0.0.2_OS_amd64.tar.gz
    
        helm plugin install .
    popd

```

After you download and install the entire package to verify the steps you can check this one

```sh
➜  helm migrate --help
    helm migrate rocking yeaaa :))

    Usage:
    helm migrate [flags]
    helm [command]

    Available Commands:
    completion  Generate the autocompletion script for the specified shell
    help        Help about any command
    mode        Type of the operation in cluster

    Flags:
    -h, --help   help for helm

    Use "helm [command] --help" for more information about a command.
```

## Usage 

```sh
helm migrate mode restore --namespace kube-system --sourcecluster kind-kind --targetcluster kind-kind-test-cluster \
--kubeconfig $KUBECONFIGPATH
```