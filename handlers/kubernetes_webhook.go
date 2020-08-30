package handlers

import (
	"fmt"
	"net/http"
)

type KubernetesWebhookHandler struct {

}

func (h KubernetesWebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "deployment captured")
}
