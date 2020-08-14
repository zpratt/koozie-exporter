package pods

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Container struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Pod struct {
	Name       string      `json:"name"`
	Containers []Container `json:"containers"`
}

type PodServiceInterface interface {
	ListPods(namespace string) []Pod
}

type PodService struct {
	Client kubernetes.Interface
}

func (p PodService) ListPods(namespace string) []Pod {
	podList, _ := p.Client.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	result := make([]Pod, len(podList.Items))

	for i, pod := range podList.Items {
		result[i] = Pod{
			Name:       pod.Name,
			Containers: extractContainersFromPod(pod),
		}
	}

	return result
}

func extractContainersFromPod(pod v1.Pod) []Container {
	containers := make([]Container, len(pod.Spec.Containers))

	for i, container := range pod.Spec.Containers {
		containers[i] = Container{
			Name:  container.Name,
			Image: container.Image,
		}
	}

	return containers
}
