package e2e

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/kubernetes"
)

var (
	clientset    *kubernetes.Clientset
	clusterName  = "capi-test-cluster"
	kubeconfig   = fmt.Sprintf("/tmp/%s.kubeconfig", clusterName)
	ctx          = context.TODO()
	namespace    = "default"
	initialNodes int
)

func TestTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tests Suite")
}

var _ = BeforeSuite(func() {
	By("Creating a test cluster using kind")
	createKindCluster(clusterName, kubeconfig)
})

// var _ = AfterSuite(func() {
// 	By("Deleting the test cluster")
// 	deleteKindCluster(clusterName)
// })

func createKindCluster(clusterName, kubeconfig string) {
	// Check if the cluster already exists
	cmd := exec.Command("kind", "get", "clusters")
	output, err := cmd.Output()
	Expect(err).NotTo(HaveOccurred(), "Failed to list kind clusters")

	existingClusters := string(output)
	if containsCluster(existingClusters, clusterName) {
		fmt.Printf("Cluster %s already exists. Skipping creation.\n", clusterName)
		return
	}

	// Create the cluster if it doesn't exist
	cmd = exec.Command("kind", "create", "cluster", "--name", clusterName, "--kubeconfig", kubeconfig)
	cmd.Stdout = GinkgoWriter
	cmd.Stderr = GinkgoWriter
	Expect(cmd.Run()).To(Succeed(), "Failed to create kind cluster")
}

// Helper function to check if a cluster exists in the list
func containsCluster(existingClusters, clusterName string) bool {
	clusters := strings.Split(existingClusters, "\n")
	for _, cluster := range clusters {
		if strings.TrimSpace(cluster) == clusterName {
			return true
		}
	}
	return false
}

func deleteKindCluster(clusterName string) {
	cmd := exec.Command("kind", "delete", "cluster", "--name", clusterName)
	cmd.Stdout = GinkgoWriter
	cmd.Stderr = GinkgoWriter
	Expect(cmd.Run()).To(Succeed(), "Failed to delete kind cluster")
}
