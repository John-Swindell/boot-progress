package k8slab

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// LoadRESTConfig returns a *rest.Config for talking to a Kubernetes API server.
//
// Logic:
//  1. If kubeconfigPath != "", use clientcmd.BuildConfigFromFlags("", kubeconfigPath).
//  2. Otherwise return an error (we don't try in-cluster here — the capstone does).
//
// Don't reach for InClusterConfig in this exercise; just the explicit path.
func LoadRESTConfig(kubeconfigPath string) (*rest.Config, error) {
	// TODO
	_ = clientcmd.BuildConfigFromFlags
	return nil, nil
}

// ListPodNames returns the names of pods in namespace matching labelSelector.
// An empty labelSelector matches all pods in the namespace.
//
// Use clientset.CoreV1().Pods(namespace).List with metav1.ListOptions{LabelSelector: ...}.
//
// Names are returned in the order returned by the API.
func ListPodNames(ctx context.Context, clientset kubernetes.Interface, namespace, labelSelector string) ([]string, error) {
	// TODO
	return nil, nil
}

// IsReady reports whether p has condition PodReady == ConditionTrue.
func IsReady(p corev1.Pod) bool {
	// TODO
	return false
}

// CountReady returns how many pods in pods are ready (per IsReady).
func CountReady(pods []corev1.Pod) int {
	// TODO
	return 0
}
