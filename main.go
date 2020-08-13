package main

import (
	"encoding/json"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"net/http"
	"strings"
	"topokube/handlers"
	"topokube/internal/namespaces"
	"topokube/internal/pods"
)

func apiRoot(w http.ResponseWriter, r *http.Request) {
	client, err := createClient()

	if err != nil {
		http.Error(w, "failed to connect to kubernetes", http.StatusInternalServerError)
	}

	namespaceService := namespaces.NamespaceServiceImpl{
		Client: client,
	}

	namespaceList := namespaceService.ListNamespaces()

	json.NewEncoder(w).Encode(namespaceList)
}

func getPods(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")[1:]
	namespace := pathParts[len(pathParts) - 1]
	client, _ := createClient()

	podList := pods.ListPods(client, namespace)

	json.NewEncoder(w).Encode(podList)
}

// NOTE: i should use an environment variable to determine if i should start this with the fake server
// to allow for running outside of k8s
func createClient() (*kubernetes.Clientset, error) {
	config, err := rest.InClusterConfig()

	if err != nil {
		klog.Errorf("error getting cluster config")
		return nil, err
	}

	client, err := kubernetes.NewForConfig(config)

	if err != nil {
		klog.Errorf("error creating kubernetes client")
		return nil, err
	}

	return client, nil
}

func main() {
	http.HandleFunc("/api", apiRoot)
	//http.HandleFunc("/health", health)
	http.Handle("/health", handlers.HealthHandler{})
	http.HandleFunc("/api/pods/", getPods)
	http.ListenAndServe(":8080", nil)
}
