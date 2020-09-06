package handlers

import (
	"fmt"
	"k8s.io/klog/v2"
	"net/http"
)

type KubernetesWebhookHandler struct {
}

func (h KubernetesWebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	klog.Infof("captured a deployment")
	fmt.Fprintf(w, "deployment captured")
}
