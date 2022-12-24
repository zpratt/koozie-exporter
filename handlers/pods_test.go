package handlers

import (
	"bytes"
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"log"
	"net/http/httptest"
	"testing"
	"topokube/internal/pods"
)

type dummyPodService struct {
	someNamespace string
	podName       string
}

func (dp dummyPodService) ListPods(namespace string) []pods.Pod {
	if namespace != dp.someNamespace {
		errorMessage := fmt.Sprintf("expected namespace not used to find pods. received %s, expected %s", namespace, dp.someNamespace)
		log.Fatalf(errorMessage)
	}

	fakePodName := gofakeit.LetterN(10)
	container := pods.Container{
		Name:  "foo",
		Image: "bar",
	}
	pod := pods.Pod{
		Name:       fakePodName,
		Containers: []pods.Container{container},
	}
	return []pods.Pod{pod}
}

func Test_should_not_blow_up(t *testing.T) {
	expectedNamespace := gofakeit.LetterN(10)
	malformedRequest := []byte("malformed")
	recorder := httptest.NewRecorder()
	recorder.Body = bytes.NewBuffer(nil)
	urlPath := fmt.Sprintf("/api/pods/%s", expectedNamespace)

	request := httptest.NewRequest("GET", urlPath, bytes.NewBuffer(malformedRequest))

	podsHandler := PodsHandler{
		Service: dummyPodService{
			someNamespace: expectedNamespace,
			podName:       "",
		},
	}

	podsHandler.ServeHTTP(recorder, request)
}
