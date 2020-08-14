package namespaces

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"testing"
)

func TestListNamespaces(t *testing.T) {
	client := fake.NewSimpleClientset()
	someNamespace := "someNamespace"
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: someNamespace,
		},
	}
	client.CoreV1().Namespaces().Create(namespace)

	namespaceService := NamespaceService{
		Client: client,
	}

	namespaceList := namespaceService.ListNamespaces()
	numberOfListedNamespaces := len(namespaceList)

	if numberOfListedNamespaces == 0 {
		t.Fatalf("no namespaces")
	}

	for _, namespace := range namespaceList {
		if namespace.Name != someNamespace {
			t.Fatalf("expected namespace not created")
		}
	}
}
