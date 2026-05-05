package solutions

import (
	"context"
	"errors"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func LoadRESTConfig(kubeconfigPath string) (*rest.Config, error) {
	if kubeconfigPath == "" {
		return nil, errors.New("kubeconfig path required")
	}
	return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
}

func ListPodNames(ctx context.Context, clientset kubernetes.Interface, namespace, labelSelector string) ([]string, error) {
	list, err := clientset.CoreV1().Pods(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(list.Items))
	for _, p := range list.Items {
		names = append(names, p.Name)
	}
	return names, nil
}

func IsReady(p corev1.Pod) bool {
	for _, c := range p.Status.Conditions {
		if c.Type == corev1.PodReady && c.Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}

func CountReady(pods []corev1.Pod) int {
	n := 0
	for _, p := range pods {
		if IsReady(p) {
			n++
		}
	}
	return n
}
