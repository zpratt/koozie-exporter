package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
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

	registry := prometheus.NewRegistry()
	registry.MustRegister(version.NewCollector("koozie"))

	namespaceService := namespaces.NamespaceService{
		Client: client,
	}
	podService := pods.PodService{
		Client: client,
	}

	http.Handle("/health", handlers.HealthHandler{})
	http.Handle(
		"/metrics",
		promhttp.HandlerFor(prometheus.Gatherers{registry}, promhttp.HandlerOpts{Registry: registry}))

	http.Handle("/api", handlers.NamespaceHander{
		Service: namespaceService,
	})
	http.Handle("/api/pods/", handlers.PodsHandler{
		Service: podService,
	})
	http.Handle("/api/kubernetes-webhook", handlers.NewKubernetesWebhookHandler(registry))

	err := http.ListenAndServeTLS(":8443", "/certificates/tls.crt", "/certificates/tls.key", nil)
	if err != nil {
		klog.Error(err)
	}
}
