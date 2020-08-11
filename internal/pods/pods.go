package pods

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Pod struct {
	Name string
}

func ListPods(k kubernetes.Interface, namespace string) []Pod {
	podList, _ := k.CoreV1().Pods(namespace).List(metav1.ListOptions{})
	result := make([]Pod, len(podList.Items))

	for i, pod := range podList.Items {
		result[i] = Pod{
			Name: pod.Name,
		}
	}

	return result
}
