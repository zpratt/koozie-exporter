package handlers

import (
	"encoding/json"
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

	body, _ := ioutil.ReadAll(r.Body)
	_, _, decodeErr := deserializer.Decode(body, nil, &admissionReview)

	if decodeErr == nil {
		kind := admissionReview.Request.Kind
		klog.Infof("captured a deployment %s", kind)
		_ = json.NewEncoder(w).Encode(response)
	} else {
		klog.Errorf("received malformed webhook request")
		response.Allowed = false
		_ = json.NewEncoder(w).Encode(response)
	}
}

func createDeserializer() runtime.Decoder {
	scheme := runtime.NewScheme()
	codecs := serializer.NewCodecFactory(scheme)
	deserializer := codecs.UniversalDeserializer()
	return deserializer
}
