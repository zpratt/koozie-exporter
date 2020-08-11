package pods

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"log"
	"testing"
)

func TestListPodsInNamespace(t *testing.T) {
	someNamespace := "someNamespace"
	somePodName := "somePod"
	client := fake.NewSimpleClientset()

	namespace := givenANamespaceSpec(someNamespace)
	podSpec := givenAPodSpecInNamespace(somePodName, someNamespace)
	givenANamespaceAndPod(client, namespace, someNamespace, podSpec)

	pods := ListPods(client, someNamespace)

	if len(pods) == 0 {
		log.Fatalf("ListPods returned empty result")
	}

	for _, pod := range pods {
		if pod.Name != somePodName {
			log.Fatalf("Expected pod does not match result")
		}
	}
}

func givenANamespaceAndPod(client *fake.Clientset, namespace *v1.Namespace, someNamespace string, podSpec *v1.Pod) {
	client.CoreV1().Namespaces().Create(namespace)
	client.CoreV1().Pods(someNamespace).Create(podSpec)
}

func givenAPodSpecInNamespace(somePodName string, someNamespace string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      somePodName,
			Namespace: someNamespace,
		},
	}
}

func givenANamespaceSpec(someNamespace string) *v1.Namespace {
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: someNamespace,
		},
	}
	return namespace
}
