package handlers

import (
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"io/ioutil"
	v1 "k8s.io/api/admission/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/klog/v2"
	"net/http"
)

type KubernetesWebhookHandler struct {
}

func (h KubernetesWebhookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	deserializer := createDeserializer()
	var admissionReview v1.AdmissionReview

	response := v1.AdmissionResponse{
		Allowed: true,
	}

	body, err := ioutil.ReadAll(r.Body)
	deserializer.Decode(body, nil, &admissionReview)

	if err != nil {
		klog.Errorf("error handling request %s", err.Error())
		klog.Errorf("request was %s", r.Body)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	klog.Infof("captured a deployment")
	kind := admissionReview.Request.Kind

	klog.Infof("captured a deployment %s", kind)
	promauto.NewCounter(prometheus.CounterOpts{
		Name: "topokube_somenamespace_podname_deployments_total",
		Help: "Number of deployments",
	})
	json.NewEncoder(w).Encode(response)
}

func createDeserializer() runtime.Decoder {
	scheme := runtime.NewScheme()
	codecs := serializer.NewCodecFactory(scheme)
	deserializer := codecs.UniversalDeserializer()
	return deserializer
}
