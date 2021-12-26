package handlers

import (
	"bytes"
	"encoding/json"
	"k8s.io/api/admission/v1beta1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"log"
	"net/http/httptest"
	"testing"
)

type DummyContainerData struct {
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

	handler := KubernetesWebhookHandler{}

	handler.ServeHTTP(recorder, request)
	json.Unmarshal(recorder.Body.Bytes(), &admissionResponse)

	if admissionResponse.Allowed != true {
		log.Fatalf("admission response was not allowed: %s", admissionResponse)
	}
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

func givenContainers() DummyContainerData {
	return DummyContainerData{
		[]string{"randomName1", "randomName2"},
		[]string{"randomContainer1", "randomContainer2"},
	}
}

func givenAPodSpecInNamespace(containers DummyContainerData, somePodName string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: somePodName,
		},
		Spec: v1.PodSpec{
			Containers: givenASliceOfContainers(containers),
		},
	}
}

func givenASliceOfContainers(someContainers DummyContainerData) []v1.Container {
	containers := make([]v1.Container, len(someContainers.containerNames))

	for i, containerName := range someContainers.containerNames {
		containers[i] = v1.Container{
			Name:  containerName,
			Image: someContainers.containerImages[i],
		}
	}

	return containers
}
