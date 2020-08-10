package main

import (
	"encoding/json"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/klog/v2"
	"net/http"
	"topokube/internal/namespaces"
)

func apiRoot(w http.ResponseWriter, r *http.Request) {
	client, err := createClient()

	if err != nil {
		http.Error(w, "failed to connect to kubernetes", http.StatusInternalServerError)
	}

	namespaceList := namespaces.ListNamespaces(client)

	json.NewEncoder(w).Encode(namespaceList)
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
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
	http.HandleFunc("/", apiRoot)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}
