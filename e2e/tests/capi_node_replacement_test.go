package e2e

import (
	// "time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// func TestE2E(t *testing.T) {
// 	RegisterFailHandler(Fail)
// 	RunSpecs(t, "E2E Suite")
// }

// var (
// 	clientset    *kubernetes.Clientset
// 	clusterName  = "capi-test-cluster"
// 	kubeconfig   = fmt.Sprintf("/tmp/%s.kubeconfig", clusterName)
// 	ctx          = context.TODO()
// 	namespace    = "default"
// 	initialNodes int
// )

// var _ = BeforeSuite(func() {
// 	By("Creating a test cluster using kind")
// 	createKindCluster(clusterName, kubeconfig)
// })

// var _ = AfterSuite(func() {
// 	By("Deleting the test cluster")
// 	deleteKindCluster(clusterName)
// })

var _ = Describe("CAPI Node Replacement with Test Cluster Environment", func() {
	BeforeEach(func() {
		By("Initializing the Kubernetes client")
		config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
		Expect(err).NotTo(HaveOccurred(), "Failed to load kubeconfig")
		clientset, err = kubernetes.NewForConfig(config)
		Expect(err).NotTo(HaveOccurred(), "Failed to create Kubernetes client")

		By("Getting the initial node count")
		nodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
		Expect(err).NotTo(HaveOccurred(), "Failed to list nodes")
		initialNodes = len(nodes.Items)
		Expect(initialNodes).To(BeNumerically(">=", 1), "Cluster should have at least one node")
	})

	// It("should create a new node when one is deleted", func() {
	// 	By("Deleting a node")
	// 	nodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	// 	Expect(err).NotTo(HaveOccurred(), "Failed to list nodes")
	// 	nodeToDelete := nodes.Items[0].Name
	// 	err = clientset.CoreV1().Nodes().Delete(ctx, nodeToDelete, metav1.DeleteOptions{})
	// 	Expect(err).NotTo(HaveOccurred(), "Failed to delete node")

	// 	By("Waiting for a new node to be provisioned")
	// 	Eventually(func() int {
	// 		currentNodes, err := clientset.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
	// 		Expect(err).NotTo(HaveOccurred(), "Failed to list nodes during wait")
	// 		return len(currentNodes.Items)
	// 	}, 5*time.Minute, 10*time.Second).Should(Equal(initialNodes), "New node should be provisioned")
	// })

})

// func createKindCluster(clusterName, kubeconfig string) {
// 	cmd := exec.Command("kind", "create", "cluster", "--name", clusterName, "--kubeconfig", kubeconfig)
// 	cmd.Stdout = GinkgoWriter
// 	cmd.Stderr = GinkgoWriter
// 	Expect(cmd.Run()).To(Succeed(), "Failed to create kind cluster")
// }

// func deleteKindCluster(clusterName string) {
// 	cmd := exec.Command("kind", "delete", "cluster", "--name", clusterName)
// 	cmd.Stdout = GinkgoWriter
// 	cmd.Stderr = GinkgoWriter
// 	Expect(cmd.Run()).To(Succeed(), "Failed to delete kind cluster")
// }
