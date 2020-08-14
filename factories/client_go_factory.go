package factories

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
)

type IClientGoFactory interface {
	New() kubernetes.Interface
}

type ClientGoFactory struct {
}

func (c ClientGoFactory) New() kubernetes.Interface {
	config, err := rest.InClusterConfig()

	if err != nil {
		klog.Errorf("error getting cluster config")
		return nil
	}

	client, err := kubernetes.NewForConfig(config)

	if err != nil {
		klog.Errorf("error creating kubernetes client")
		return nil
	}

	return client
}
