# K-Foundry

## Getting Started

### Running on kcp

1. Run KCP with the following command:

```sh
make kcp-server
```

From this point onwards you can inspect kcp configuration using kubeconfig:

```sh
export KUBECONFIG=.test/kcp.kubeconfig
```

1. Bootstrap the KCP server with the following command:

```sh
export KUBECONFIG=./.test/kcp.kubeconfig
make kcp-bootstrap
```

1. Run controller:

```sh
export KUBECONFIG=./.test/kcp.kubeconfig
make run-local
```

1. In separate shell you can run tests to exercise the controller:

```sh
export KUBECONFIG=./.test/kcp.kubeconfig
make test
```

### Uninstall resources
To delete the resources from kcp:

```sh
make test-clean
```



### How it works
This project aims to follow the Kubernetes [Operator pattern](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/)

It uses [Controllers](https://kubernetes.io/docs/concepts/architecture/controller/)
which provides a reconcile function responsible for synchronizing resources until the desired state is reached.




