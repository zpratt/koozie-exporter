package namespaces

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Namespace struct {
	Name string `json:"name"`
}

type NamespaceServiceInterface interface {
	ListNamespaces() []Namespace
}

type NamespaceService struct {
	Client kubernetes.Interface
}

func (n NamespaceService) ListNamespaces() []Namespace {
	ctx := context.Background()
	namespaceList, _ := n.Client.CoreV1().Namespaces().List(ctx, metav1.ListOptions{})
	result := make([]Namespace, len(namespaceList.Items))

	for i, namespace := range namespaceList.Items {
		result[i] = Namespace{
			Name: namespace.Name,
		}
	}

	return result
}
