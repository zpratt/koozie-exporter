package pods

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
	"log"
	"testing"
)

type DummyContainerData struct {
	containerNames  []string
	containerImages []string
}

func TestListPodsInNamespace(t *testing.T) {
	someNamespace := "someNamespace"
	somePodName := "somePod"
	someContainers := givenContainers()
	client := fake.NewSimpleClientset()

	namespace := givenANamespaceSpec(someNamespace)
	podSpec := givenAPodSpecInNamespace(someContainers, somePodName, someNamespace)
	givenANamespaceAndPod(client, namespace, someNamespace, podSpec)

	pods := ListPods(client, someNamespace)

	if len(pods) == 0 {
		log.Fatalf("ListPods returned empty result")
	}

	for _, pod := range pods {
		if pod.Name != somePodName {
			log.Fatalf("Expected pod does not match result")
		}

		assertNumberOfPodContainers(someContainers, pod)

		assertPodContainers(someContainers, pod)
	}
}

func assertPodContainers(someContainers DummyContainerData, pod Pod) {
	for i, container := range pod.Containers {
		if container.Name != someContainers.containerNames[i] {
			log.Fatalf("Expected pod container name does not match result")
		}
		if container.Image != someContainers.containerImages[i] {
			log.Fatalf("Expected pod container image does not match result")
		}
	}
}

func assertNumberOfPodContainers(someContainers DummyContainerData, pod Pod) {
	expectedNumberOfPodContainers := len(someContainers.containerNames)
	actualNumberOfPodContainers := len(pod.Containers)
	if actualNumberOfPodContainers != expectedNumberOfPodContainers {
		log.Fatalf(
			"Expected %d containers, received %d",
			expectedNumberOfPodContainers,
			actualNumberOfPodContainers,
		)
	}
}

func givenContainers() DummyContainerData {
	return DummyContainerData{
		[]string{"randomName1", "randomName2"},
		[]string{"randomContainer1", "randomContainer2"},
	}
}

func givenANamespaceAndPod(client *fake.Clientset, namespace *v1.Namespace, someNamespace string, podSpec *v1.Pod) {
	client.CoreV1().Namespaces().Create(namespace)
	client.CoreV1().Pods(someNamespace).Create(podSpec)
}

func givenAPodSpecInNamespace(containers DummyContainerData, somePodName string, someNamespace string) *v1.Pod {
	return &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      somePodName,
			Namespace: someNamespace,
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

func givenANamespaceSpec(someNamespace string) *v1.Namespace {
	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: someNamespace,
		},
	}
	return namespace
}
