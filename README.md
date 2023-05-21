# helm restore

This is a basic plugin that allows you backup and restore whole cluster and specified namespace.


## Build & Installation
The installation part is very simple you can easily install with via native helm commands like this;

* Installation
```sh
    git clone git@github.com:WoodProgrammer/helm-restore.git
    
    pushd helm-restore
        wget https://github.com/WoodProgrammer/helm-restore/releases/download/0.0.1/helm-restore_0.0.1_OS_amd64.tar.gz
        tar -xvf helm-restore_0.0.1_OS_amd64.tar.gz
    
        helm plugin install .
    popd

```

After you download and install the entire package to verify the steps you can check this one

```sh
➜  helm-restore git:(main) helm restore --help
    helm restore rocking yeaaa :))

    Usage:
    helm restore [flags]
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

* Backup Only Mode

Coming sooon....

* Restore Mode

```sh
helm restore mode restore --namespace kube-system --sourcecluster kind-kind --targetcluster kind-kind-test-cluster
```

## TODO

* Backup only mode
* Better output
* Data Backup and Restore