package e2e

import (
	"fmt"
	"os/exec"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// func TestE2E(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "E2E Service Suite")
// }

// var (
// 	clientset   *kubernetes.Clientset
// 	clusterName = "capi-test-cluster"
// 	namespace   = "default"
// 	kubeconfig  = fmt.Sprintf("/tmp/%s.kubeconfig", clusterName) // Update this path
// 	ctx         = context.TODO()
// )

// var _ = BeforeSuite(func() {
// 	By("Initializing the Kubernetes client")
// 	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
// 	Expect(err).NotTo(HaveOccurred(), "Failed to load kubeconfig")
// 	clientset, err = kubernetes.NewForConfig(config)
// 	Expect(err).NotTo(HaveOccurred(), "Failed to create Kubernetes client")
// })

// var _ = AfterSuite(func() {
// 	By("Cleaning up resources")
// 	// Add any global cleanup logic if necessary
// })

var _ = Describe("ClusterIP Service Test", func() {
	var serviceName string

	BeforeEach(func() {
		By("Initializing the Kubernetes client")
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		Expect(err).NotTo(HaveOccurred(), "Failed to load kubeconfig")
		clientset, err = kubernetes.NewForConfig(config)
		Expect(err).NotTo(HaveOccurred(), "Failed to create Kubernetes client")
		serviceName = "test-clusterip-service"

		By("Creating backend pods")
		createBackendPods(clientset, namespace)

		By("Creating a ClusterIP service")
		createClusterIPService(clientset, namespace, serviceName)
	})

	AfterEach(func() {
		By("Deleting the service")
		err := clientset.CoreV1().Services(namespace).Delete(ctx, serviceName, metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "Failed to delete service")

		By("Deleting backend pods")
		err = clientset.CoreV1().Pods(namespace).DeleteCollection(ctx, metav1.DeleteOptions{}, metav1.ListOptions{LabelSelector: "app=test-backend"})
		Expect(err).NotTo(HaveOccurred(), "Failed to delete backend pods")
	})

	It("should load balance traffic and remain available when a pod dies", func() {
		By("Validating load balancing")
		validateLoadBalancing(clientset, namespace, serviceName)

		By("Simulating a pod failure")
		deleteOneBackendPod(clientset, namespace)

		By("Ensuring service remains available")
		validateServiceAvailability(clientset, namespace, serviceName)
	})
})

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
		Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("Failed to create pod %s", pod.Name))
	}
}

func createClusterIPService(clientset *kubernetes.Clientset, namespace, serviceName string) {
	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: serviceName,
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{"app": "test-backend"},
			Ports: []v1.ServicePort{
				{
					Protocol:   v1.ProtocolTCP,
					Port:       80,
					TargetPort: intstrFromInt(80),
				},
			},
			Type: v1.ServiceTypeClusterIP,
		},
	}
	_, err := clientset.CoreV1().Services(namespace).Create(ctx, service, metav1.CreateOptions{})
	Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("Failed to create service %s", serviceName))
}

// func validateLoadBalancing(clientset *kubernetes.Clientset, namespace, serviceName string) {
// 	// Simulate sending multiple requests to the service IP
// 	service, err := clientset.CoreV1().Services(namespace).Get(ctx, serviceName, metav1.GetOptions{})
// 	Expect(err).NotTo(HaveOccurred(), "Failed to get service details")

// 	serviceIP := service.Spec.ClusterIP
// 	Expect(serviceIP).NotTo(BeEmpty(), "Service has no assigned ClusterIP")

// 	hitCounts := map[string]int{}
// 	for i := 0; i < 10; i++ {
// 		output, err := exec.Command("curl", fmt.Sprintf("http://%s", serviceIP)).Output()
// 		Expect(err).NotTo(HaveOccurred(), "Failed to reach service")
// 		hitCounts[string(output)]++
// 		time.Sleep(500 * time.Millisecond) // Small delay between requests
// 	}

// 	Expect(len(hitCounts)).To(BeNumerically(">", 1), "Traffic was not distributed across pods")
// }

func validateLoadBalancing(clientset *kubernetes.Clientset, namespace, serviceName string) {
	By("Creating a debug pod to test load balancing from inside the cluster")
	debugPod := &v1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name: "debug-pod",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name:    "debug-container",
					Image:   "curlimages/curl:latest",  // Minimal image with curl
					Command: []string{"sleep", "3600"}, // Keep the pod running
				},
			},
			RestartPolicy: v1.RestartPolicyNever,
		},
	}
	_, err := clientset.CoreV1().Pods(namespace).Create(ctx, debugPod, metav1.CreateOptions{})
	Expect(err).NotTo(HaveOccurred(), "Failed to create debug pod")

	// Wait for the pod to be ready
	Eventually(func() bool {
		pod, err := clientset.CoreV1().Pods(namespace).Get(ctx, "debug-pod", metav1.GetOptions{})
		Expect(err).NotTo(HaveOccurred(), "Failed to fetch debug pod status")
		return pod.Status.Phase == v1.PodRunning
	}, 2*time.Minute, 5*time.Second).Should(BeTrue(), "Debug pod did not start running in time")

	defer func() {
		By("Deleting the debug pod")
		err = clientset.CoreV1().Pods(namespace).Delete(ctx, "debug-pod", metav1.DeleteOptions{})
		Expect(err).NotTo(HaveOccurred(), "Failed to delete debug pod")
	}()

	By("Validating load balancing")
	service, err := clientset.CoreV1().Services(namespace).Get(ctx, serviceName, metav1.GetOptions{})
	Expect(err).NotTo(HaveOccurred(), "Failed to get service details")

	serviceIP := service.Spec.ClusterIP
	Expect(serviceIP).NotTo(BeEmpty(), "Service has no assigned ClusterIP")

	hitCounts := map[string]int{}
	for i := 0; i < 10; i++ {
		cmd := []string{"kubectl", "--kubeconfig", kubeconfig, "exec", "-n", namespace, "debug-pod", "--", "curl", fmt.Sprintf("http://%s", serviceIP)}
		output, err := exec.Command(cmd[0], cmd[1:]...).Output()
		Expect(err).NotTo(HaveOccurred(), "Failed to curl service from debug pod")
		hitCounts[string(output)]++
		time.Sleep(500 * time.Millisecond) // Small delay between requests
	}

	Expect(len(hitCounts)).To(BeNumerically(">", 1), "Traffic was not distributed across pods")
}

func deleteOneBackendPod(clientset *kubernetes.Clientset, namespace string) {
	pods, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{LabelSelector: "app=test-backend"})
	Expect(err).NotTo(HaveOccurred(), "Failed to list pods")
	Expect(len(pods.Items)).To(BeNumerically(">", 0), "No pods found to delete")

	err = clientset.CoreV1().Pods(namespace).Delete(ctx, pods.Items[0].Name, metav1.DeleteOptions{})
	Expect(err).NotTo(HaveOccurred(), fmt.Sprintf("Failed to delete pod %s", pods.Items[0].Name))
}

func validateServiceAvailability(clientset *kubernetes.Clientset, namespace, serviceName string) {
	service, err := clientset.CoreV1().Services(namespace).Get(ctx, serviceName, metav1.GetOptions{})
	Expect(err).NotTo(HaveOccurred(), "Failed to get service details")

	serviceIP := service.Spec.ClusterIP
	Expect(serviceIP).NotTo(BeEmpty(), "Service has no assigned ClusterIP")

	// Test if the service is still reachable
	output, err := exec.Command("curl", fmt.Sprintf("http://%s", serviceIP)).Output()
	Expect(err).NotTo(HaveOccurred(), "Service is not reachable after pod deletion")
	Expect(output).NotTo(BeEmpty(), "Service returned empty response")
}

func intstrFromInt(i int) intstr.IntOrString {
	return intstr.IntOrString{Type: intstr.Int, IntVal: int32(i)}
}
