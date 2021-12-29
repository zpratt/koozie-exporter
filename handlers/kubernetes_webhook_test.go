package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"net/http/httptest"
	"testing"
)

type dummyContainerData struct {
	containerNames  []string
	containerImages []string
}

func TestKubernetesWebhookHandler_always_approves(t *testing.T) {
	somePodName := "somePod"
	someContainers := givenContainers()
	podSpec := givenAPodSpecInNamespace(someContainers, somePodName)
	podAsJson, _ := json.Marshal(podSpec)
	admissionReview := givenAnAdmissionReview(podAsJson)

	admissionReviewAsJson, _ := json.Marshal(admissionReview)
	admissionResponse := &v1beta1.AdmissionResponse{}
	recorder := httptest.NewRecorder()
	recorder.Body = bytes.NewBuffer(nil)

	request := httptest.NewRequest("POST", "/api/kubernetes-webhook", bytes.NewBuffer(admissionReviewAsJson))

	handler, _ := makeWebhookHandler()

	handler.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &admissionResponse)

	if admissionResponse.Allowed != true {
		t.Fatalf("admission response was not allowed: %s", admissionResponse)
	}
}

func Test_should_update_count_of_deployments(t *testing.T) {
	somePodName := "somePod"
	someContainers := givenContainers()
	podSpec := givenAPodSpecInNamespace(someContainers, somePodName)
	podAsJson, _ := json.Marshal(podSpec)
	admissionReview := givenAnAdmissionReview(podAsJson)

	admissionReviewAsJson, _ := json.Marshal(admissionReview)
	recorder := httptest.NewRecorder()
	recorder.Body = bytes.NewBuffer(nil)
	request := httptest.NewRequest("POST", "/api/kubernetes-webhook", bytes.NewBuffer(admissionReviewAsJson))

	handler, actualRegistry := makeWebhookHandler()
	handler.ServeHTTP(recorder, request)

	gatherers := prometheus.Gatherers{actualRegistry}
	actualMetricLookup, err := testutil.GatherAndCount(gatherers, "koozie_deployment_count")

	if err != nil {
		t.Fatalf(err.Error())
	}

	if actualMetricLookup != 1 {
		t.Fatalf("metric not registered")
	}

	actualDeploymentCount := testutil.ToFloat64(handler.deploymentCount)
	if actualDeploymentCount != 1 {
		t.Fatalf("deployment count not incremented, actual value is %f", actualDeploymentCount)
	}
}

func TestKubernetesWebhookHandler_should_handle_malformed_requests(t *testing.T) {
	malformedRequest := []byte("malformed")
	admissionResponse := &v1beta1.AdmissionResponse{}
	recorder := httptest.NewRecorder()
	recorder.Body = bytes.NewBuffer(nil)

	request := httptest.NewRequest("POST", "/api/kubernetes-webhook", bytes.NewBuffer(malformedRequest))

	handler, _ := makeWebhookHandler()

	handler.ServeHTTP(recorder, request)
	_ = json.Unmarshal(recorder.Body.Bytes(), &admissionResponse)

	if len(recorder.Body.Bytes()) == 0 {
		t.Fatalf("no response was returned from webhook")
	}

	if admissionResponse.Allowed != false {
		t.Fatalf("admission response was not allowed: %s", admissionResponse)
	}
}

func makeWebhookHandler() (KubernetesWebhookHandler, *prometheus.Registry) {
	registry := prometheus.NewRegistry()
	return NewKubernetesWebhookHandler(registry), registry
}

func givenAnAdmissionReview(podAsJson []byte) *v1beta1.AdmissionReview {
	return &v1beta1.AdmissionReview{
		Request: &v1beta1.AdmissionRequest{
			Kind: metav1.GroupVersionKind{
				Kind:  "Pod",
				Group: "",
			},
			Namespace: "",
			Object: runtime.RawExtension{
				Raw: podAsJson,
			},
		},
	}
}

func givenContainers() dummyContainerData {
	return dummyContainerData{
		[]string{"randomName1", "randomName2"},
		[]string{"randomContainer1", "randomContainer2"},
	}
}

func givenAPodSpecInNamespace(containers dummyContainerData, somePodName string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: somePodName,
		},
		Spec: v1.PodSpec{
			Containers: givenASliceOfContainers(containers),
		},
	}
}

func givenASliceOfContainers(someContainers dummyContainerData) []v1.Container {
	containers := make([]v1.Container, len(someContainers.containerNames))

	for i, containerName := range someContainers.containerNames {
		containers[i] = v1.Container{
			Name:  containerName,
			Image: someContainers.containerImages[i],
		}
	}

	return containers
}
