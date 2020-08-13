package namespaces

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Namespace struct {
	Name string `json:"name"`
}

type NamespaceService interface {
	ListNamespaces() []Namespace
}

type NamespaceServiceImpl struct {
	Client kubernetes.Interface
}

func (n NamespaceServiceImpl) ListNamespaces() []Namespace {
	namespaceList, _ := n.Client.CoreV1().Namespaces().List(metav1.ListOptions{})
	result := make([]Namespace, len(namespaceList.Items))

	for i, namespace := range namespaceList.Items {
		result[i] = Namespace{
			Name: namespace.Name,
		}
	}

	return result
}
