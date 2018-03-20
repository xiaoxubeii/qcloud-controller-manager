[![Build Status](https://travis-ci.org/xiaoxubeii/qcloud-controller-manager.svg?branch=master)](https://travis-ci.org/xiaoxubeii/qcloud-controller-manager) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
# Kubernetes Cloud Controller Manager for Tencent Cloud
`qcloud-controller-manager` is a external cloud controller manager implementation of Kubernetes for Tencent Cloud. It is not a official project but a third party implementation. The project aims to provide a agile way to utilize QCloud services in Kubernetes.

**WARNING:** This project is not a offical version and still work in progress, be careful using it in production environment.

## Requirements
At the current state of Kubernetes, running cloud controller manager requires a few things. Please read through the requirements carefully as they are critical to running cloud controller manager on a Kubernetes cluster on QCloud.

### Version
Kubernetes version v1.9.x.

### Running qcloud-controller-manager 
Successfully running cloud-controller-manager requires some changes to your cluster configuration.

* `kube-apiserver` and `kube-controller-manager` MUST NOT specify the `--cloud-provider` flag. This ensures that it does not run any cloud specific loops that would be run by cloud controller manager. In the future, this flag will be deprecated and removed.
* `kubelet` must run with `--cloud-provider=external`. This is to ensure that the kubelet is aware that it must be initialized by the cloud controller manager before it is scheduled any work.

You can find more details from [here](https://kubernetes.io/docs/tasks/administer-cluster/running-cloud-controller/).

## Planned Implementation Details
* node controller - responsible for updating kubernetes nodes using cloud APIs and deleting kubernetes nodes that were deleted on your cloud.
* service controller - responsible for loadbalancers on your cloud against services of type LoadBalancer.

## Contributing
This project welcomes contributions and suggestions. You can submit a pull request directly or communicate with me by:

* wechat: 280816925
* email: xiaoxubeii@gmail.com

## License
The project is Open Source software released under the [Apache 2.0 license](http://www.apache.org/licenses/LICENSE-2.0.html).



