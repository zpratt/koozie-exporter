package main

import (
	"k8s.io/klog/v2"
	"net/http"
	"topokube/factories"
	"topokube/handlers"
	"topokube/internal/namespaces"
	"topokube/internal/pods"
)

func main() {
	clientGoFactory := factories.ClientGoFactory{}
	client := clientGoFactory.New()
	namespaceService := namespaces.NamespaceService{
		Client: client,
	}
	podService := pods.PodService{
		Client: client,
	}

	http.Handle("/api", handlers.NamespaceHander{
		Service: namespaceService,
	})
	http.Handle("/health", handlers.HealthHandler{})
	http.Handle("/api/pods/", handlers.PodsHandler{
		Service: podService,
	})
	http.Handle("/api/kubernetes-webhook", handlers.KubernetesWebhookHandler{})

	err := http.ListenAndServeTLS(":8443", "/certificates/tlsCert", "/certificates/tlsKey", nil)
	if err != nil {
		klog.Error(err)
	}
}
