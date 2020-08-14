package handlers

import (
	"encoding/json"
	"net/http"
	"topokube/internal/namespaces"
)

type NamespaceHander struct {
	Service namespaces.NamespaceServiceInterface
}

func (n NamespaceHander) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	namespaceList := n.Service.ListNamespaces()
	json.NewEncoder(w).Encode(namespaceList)
}
