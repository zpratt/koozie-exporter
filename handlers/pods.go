package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
	"topokube/internal/pods"
)

type PodsHandler struct {
	Service pods.PodServiceInterface
}

func (p PodsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")[1:]
	namespace := pathParts[len(pathParts)-1]

	podList := p.Service.ListPods(namespace)

	json.NewEncoder(w).Encode(podList)
}
