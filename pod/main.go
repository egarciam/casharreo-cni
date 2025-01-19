package main

import (
	"context"
	"fmt"

	// . "github.com/onsi/ginkgo/v2"
	// . "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientset *kubernetes.Clientset
	// clusterName  = "capi-test-cluster"
	kubeconfig   = "/home/yayito/.kube/kubeconfig"
	ctx          = context.TODO()
	namespace    = "default"
	initialNodes int
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	clientset, err = kubernetes.NewForConfig(config)
	createBackendPods(clientset, "default")
}

func createBackendPods(clientset *kubernetes.Clientset, namespace string) {
	for i := 1; i <= 2; i++ {
		pod := &v1.Pod{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Pod",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:   fmt.Sprintf("test-backend-%d", i),
				Labels: map[string]string{"app": "test-backend"},
			},
			Spec: v1.PodSpec{
				Containers: []v1.Container{
					{
						Name:  "test-container",
						Image: "nginx", // A simple web server image
						Ports: []v1.ContainerPort{
							{ContainerPort: 80},
						},
					},
				},
			},
		}
		_, err := clientset.CoreV1().Pods(namespace).Create(ctx, pod, metav1.CreateOptions{})
		if err != nil {
			fmt.Printf("%v", err)
		}
		//Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("Failed to create pod %s", pod.Name))
	}
}
