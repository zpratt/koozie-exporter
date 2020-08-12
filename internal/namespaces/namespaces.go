package namespaces

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Namespace struct {
	Name string `json:"name"`
}

func ListNamespaces(k kubernetes.Interface) []Namespace {
	namespaceList, _ := k.CoreV1().Namespaces().List(metav1.ListOptions{})
	result := make([]Namespace, len(namespaceList.Items))

	for i, namespace := range namespaceList.Items {
		result[i] = Namespace{
			Name: namespace.Name,
		}
	}

	return result
}
