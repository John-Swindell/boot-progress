# 16 — Kubernetes from Go (`client-go`)

> **Goal:** talk to the Kubernetes API from Go. Build the primitives the
> capstone needs: kubeconfig discovery, listing pods, watching for new pods,
> reading container logs.

## What you'll learn

- The `kubernetes.Interface` clientset, and why we use the *interface* not the concrete type
- `clientcmd` for kubeconfig discovery (and `rest.InClusterConfig` for in-cluster pods)
- Listing resources with label selectors
- Pod readiness — what the conditions actually mean
- A taste of `Watch` / informers — fully exercised in the capstone
- `client-go/kubernetes/fake` — testing without a real cluster

## Coming from Python

| Python (`kubernetes` lib) | Go (`client-go`) | Note |
|---|---|---|
| `client.CoreV1Api()` | `clientset.CoreV1()` | Versioned APIs grouped under typed accessors. |
| `v1.list_namespaced_pod(ns)` | `clientset.CoreV1().Pods(ns).List(ctx, opts)` | Always pass a context. |
| `label_selector="app=foo"` | `metav1.ListOptions{LabelSelector: "app=foo"}` | Same selector syntax. |
| `watch.Watch().stream(...)` | `clientset.CoreV1().Pods(ns).Watch(ctx, opts)` | Returns `watch.Interface` — `<-chan Event`. |
| `read_namespaced_pod_log` | `clientset.CoreV1().Pods(ns).GetLogs(name, opts).Stream(ctx)` | Returns `io.ReadCloser`. |

## kubeconfig discovery

```go
import (
    "k8s.io/client-go/rest"
    "k8s.io/client-go/tools/clientcmd"
)

func loadConfig(kubeconfig string) (*rest.Config, error) {
    if kubeconfig != "" {
        return clientcmd.BuildConfigFromFlags("", kubeconfig)
    }
    if cfg, err := rest.InClusterConfig(); err == nil {
        return cfg, nil
    }
    // fall back to ~/.kube/config
    return clientcmd.BuildConfigFromFlags("",
        clientcmd.NewDefaultClientConfigLoadingRules().GetDefaultFilename())
}
```

In-cluster config reads from `/var/run/secrets/kubernetes.io/serviceaccount/...`
when the binary is running as a pod with the appropriate ServiceAccount.

## Listing pods with a label selector

```go
list, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{
    LabelSelector: "app=loggy",
})
for _, p := range list.Items {
    fmt.Println(p.Name)
}
```

## Pod readiness

```go
import corev1 "k8s.io/api/core/v1"

func isReady(p corev1.Pod) bool {
    for _, c := range p.Status.Conditions {
        if c.Type == corev1.PodReady && c.Status == corev1.ConditionTrue {
            return true
        }
    }
    return false
}
```

`PodReady=True` means the pod's containers have passed their readiness probes.
Don't tail logs from a pod that's still starting — you'll race the container.

## Testing without a real cluster

```go
import "k8s.io/client-go/kubernetes/fake"

cs := fake.NewSimpleClientset(
    &corev1.Pod{ObjectMeta: metav1.ObjectMeta{Name: "p1", Namespace: "default"}},
)
list, _ := cs.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
// works; no cluster needed
```

`kubernetes.Interface` is the thing you depend on — both the real clientset
and the fake satisfy it. **Always type your params/fields as
`kubernetes.Interface`, not `*kubernetes.Clientset`** — that's what makes your
code testable.

## Your turn

```sh
go test ./16-kubernetes-go
```

The exercises here build the primitives the capstone needs. The capstone
glues them into a real CLI.
