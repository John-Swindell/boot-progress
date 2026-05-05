// Package k8s wires up a kubernetes.Interface using the standard discovery:
//
//  1. Explicit --kubeconfig path
//  2. In-cluster config
//  3. ~/.kube/config (or $KUBECONFIG)
package k8s

import (
	"errors"
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// NewClientset returns a typed Kubernetes clientset. If kubeconfigPath is "",
// it tries in-cluster, then $KUBECONFIG / ~/.kube/config.
func NewClientset(kubeconfigPath string) (kubernetes.Interface, error) {
	cfg, err := loadConfig(kubeconfigPath)
	if err != nil {
		return nil, fmt.Errorf("load kube config: %w", err)
	}
	cs, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, fmt.Errorf("new clientset: %w", err)
	}
	return cs, nil
}

func loadConfig(kubeconfigPath string) (*rest.Config, error) {
	if kubeconfigPath != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	}
	if cfg, err := rest.InClusterConfig(); err == nil {
		return cfg, nil
	} else if !errors.Is(err, rest.ErrNotInCluster) {
		return nil, err
	}

	if env := os.Getenv("KUBECONFIG"); env != "" {
		return clientcmd.BuildConfigFromFlags("", env)
	}
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	defaultPath := rules.GetDefaultFilename()
	if defaultPath == "" {
		return nil, errors.New("no kubeconfig found (set --kubeconfig or $KUBECONFIG)")
	}
	return clientcmd.BuildConfigFromFlags("", defaultPath)
}
