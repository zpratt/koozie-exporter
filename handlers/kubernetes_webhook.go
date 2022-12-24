package handlers

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"io/ioutil"
	v1 "k8s.io/api/admission/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"net/http"
)

type KubernetesWebhookHandler struct {
	registry        *prometheus.Registry
	deploymentCount *prometheus.CounterVec
}

func NewKubernetesWebhookHandler(registry *prometheus.Registry) KubernetesWebhookHandler {
	counterOpts := prometheus.CounterOpts{
		Namespace: "koozie",
		Subsystem: "deployment",
		Name:      "count",
		Help:      "Number of deployments",
	}

	labelNames := []string{"deployment_name"}
	deploymentCollector := prometheus.NewCounterVec(counterOpts, labelNames)

	registry.MustRegister(deploymentCollector)

	return KubernetesWebhookHandler{
		deploymentCount: deploymentCollector,
		registry:        registry,
	}
}

func (h KubernetesWebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var admissionReview v1.AdmissionReview
	var podMetadata metav1.PartialObjectMetadata

	response := v1.AdmissionResponse{
		Allowed: true,
	}
	body, _ := ioutil.ReadAll(r.Body)
	decodeErr := json.Unmarshal(body, &admissionReview)

	if decodeErr == nil {
		kind := admissionReview.Request.Kind
		_ = json.Unmarshal(admissionReview.Request.Object.Raw, &podMetadata)
		klog.Infof("captured a deployment %s", kind)
		_ = json.NewEncoder(w).Encode(response)
		h.deploymentCount.With(prometheus.Labels{"deployment_name": podMetadata.Name}).Inc()
	} else {
		klog.Errorf("received malformed webhook request")
		response.Allowed = false
		_ = json.NewEncoder(w).Encode(response)
	}
}
