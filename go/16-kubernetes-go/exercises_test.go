package k8slab

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

const sampleKubeconfig = `
apiVersion: v1
kind: Config
clusters:
- name: test
  cluster:
    server: https://example.invalid
contexts:
- name: test
  context:
    cluster: test
    user: u
current-context: test
users:
- name: u
  user:
    token: abc
`

func TestLoadRESTConfig(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "kubeconfig")
	if err := os.WriteFile(path, []byte(sampleKubeconfig), 0o600); err != nil {
		t.Fatal(err)
	}
	cfg, err := LoadRESTConfig(path)
	if err != nil {
		t.Fatalf("err = %v", err)
	}
	if cfg == nil || cfg.Host == "" {
		t.Errorf("got %+v", cfg)
	}

	if _, err := LoadRESTConfig(""); err == nil {
		t.Error("expected error on empty path")
	}
}

func makePod(ns, name, app string, ready bool) *corev1.Pod {
	cond := corev1.ConditionFalse
	if ready {
		cond = corev1.ConditionTrue
	}
	return &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
			Labels:    map[string]string{"app": app},
		},
		Status: corev1.PodStatus{
			Conditions: []corev1.PodCondition{
				{Type: corev1.PodReady, Status: cond},
			},
		},
	}
}

func TestListPodNames(t *testing.T) {
	cs := fake.NewSimpleClientset(
		makePod("default", "p1", "loggy", true),
		makePod("default", "p2", "loggy", false),
		makePod("default", "p3", "other", true),
		makePod("kube-system", "p4", "loggy", true),
	)

	t.Run("all in namespace", func(t *testing.T) {
		names, err := ListPodNames(context.Background(), cs, "default", "")
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		sort.Strings(names)
		want := []string{"p1", "p2", "p3"}
		if len(names) != len(want) {
			t.Fatalf("got %v, want %v", names, want)
		}
		for i := range want {
			if names[i] != want[i] {
				t.Errorf("got %v, want %v", names, want)
			}
		}
	})

	t.Run("with selector", func(t *testing.T) {
		names, err := ListPodNames(context.Background(), cs, "default", "app=loggy")
		if err != nil {
			t.Fatalf("err = %v", err)
		}
		sort.Strings(names)
		if len(names) != 2 || names[0] != "p1" || names[1] != "p2" {
			t.Errorf("got %v", names)
		}
	})
}

func TestIsReadyAndCountReady(t *testing.T) {
	pods := []corev1.Pod{
		*makePod("default", "p1", "x", true),
		*makePod("default", "p2", "x", false),
		*makePod("default", "p3", "x", true),
	}
	if !IsReady(pods[0]) || IsReady(pods[1]) {
		t.Errorf("IsReady wrong")
	}
	if got := CountReady(pods); got != 2 {
		t.Errorf("CountReady = %d, want 2", got)
	}
}
