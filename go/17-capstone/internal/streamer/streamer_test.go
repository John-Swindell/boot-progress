package streamer

import (
	"context"
	"testing"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestIsReady(t *testing.T) {
	ready := corev1.Pod{Status: corev1.PodStatus{Conditions: []corev1.PodCondition{
		{Type: corev1.PodReady, Status: corev1.ConditionTrue},
	}}}
	notReady := corev1.Pod{Status: corev1.PodStatus{Conditions: []corev1.PodCondition{
		{Type: corev1.PodReady, Status: corev1.ConditionFalse},
	}}}
	no := corev1.Pod{}

	if !isReady(ready) {
		t.Error("ready pod said not ready")
	}
	if isReady(notReady) {
		t.Error("not-ready pod said ready")
	}
	if isReady(no) {
		t.Error("no-condition pod said ready")
	}
}

func TestContainersOf(t *testing.T) {
	p := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{Name: "p1"},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{Name: "app"},
				{Name: "sidecar"},
			},
		},
	}
	if got := containersOf(p, ""); len(got) != 2 || got[0] != "app" || got[1] != "sidecar" {
		t.Errorf("all = %v", got)
	}
	if got := containersOf(p, "sidecar"); len(got) != 1 || got[0] != "sidecar" {
		t.Errorf("filter = %v", got)
	}
	if got := containersOf(p, "missing"); len(got) != 0 {
		t.Errorf("missing = %v", got)
	}
}

func TestTrackedSet(t *testing.T) {
	s := newTrackedSet()
	if !s.add("a") {
		t.Error("first add should be true")
	}
	if s.add("a") {
		t.Error("duplicate add should be false")
	}
	s.remove("a")
	if !s.add("a") {
		t.Error("readd after remove should be true")
	}
}

// Note: full Run() is exercised end-to-end via `make e2e` against a kind
// cluster. The fake clientset does not support GetLogs streaming, so we
// keep Run() integration testing in e2e land.
func TestRunNotImplementedSentinel(t *testing.T) {
	// This test exists so that, at the start, ./... still tests this package
	// without panicking. Once you implement Run, you can delete it.
	cfg := Config{}
	if err := Run(context.Background(), cfg); err == nil {
		// Once implemented, Run() will error on a nil clientset. That's also
		// fine — flip this assertion to require some error and move on.
		t.Skip("Run() returned nil — assuming implemented; revise this test if needed.")
	}
}
