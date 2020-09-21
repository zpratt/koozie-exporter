package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	http.Handle("/api/metrics", promhttp.Handler())

	http.ListenAndServeTLS(":8443", "/certificates/tlsCert", "/certificates/tlsKey", nil)
}
