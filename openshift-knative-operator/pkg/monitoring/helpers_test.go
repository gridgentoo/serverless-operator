package monitoring

import (
	"os"
	"testing"

	mf "github.com/manifestival/manifestival"
	"github.com/manifestival/manifestival/fake"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	servingNamespace  = "knative-serving"
	eventingNamespace = "knative-eventing"
)

func init() {
	os.Setenv(smRbacManifestPath, "testdata/rbac-proxy.yaml")
}

func TestSetupServingRbacTransformation(t *testing.T) {
	client := fake.New()
	manifest, err := mf.NewManifest("testdata/rbac.yaml", mf.UseClient(client))
	if err != nil {
		t.Errorf("Unable to load test manifest: %v", err)
	}
	transforms := []mf.Transformer{injectNamespaceWithSubject(servingNamespace, OpenshiftMonitoringNamespace)}
	if manifest, err = manifest.Transform(transforms...); err != nil {
		t.Errorf("Unable to transform test manifest: %v", err)
	}
	if err := manifest.Apply(); err != nil {
		t.Errorf("Unable to apply the test manifest %v", err)
	}
	u := createRole(prometheusRoleName, servingNamespace)
	_, err = client.Get(u)
	if err != nil {
		t.Errorf("Unable to get the role %v", err)
	}
	u = createRole("test-role", "default")
	_, err = client.Get(u)
	if err != nil {
		t.Errorf("Unable to get the role %v", err)
	}
	u = createClusterRole()
	_, err = client.Get(u)
	if err != nil {
		t.Errorf("Unable to get the cluster role %v", err)
	}
	u = createRoleBinding(prometheusRoleName, servingNamespace)
	resultRoleBinding, err := client.Get(u)
	if err != nil {
		t.Errorf("Unable to get the rolebinding %v", err)
	}
	checkSubjects(t, resultRoleBinding.Object, OpenshiftMonitoringNamespace)
	u = createRoleBinding("test-rb", "default")
	resultRoleBinding, err = client.Get(u)
	if err != nil {
		t.Errorf("Unable to get the rolebinding %v", err)
	}
	checkSubjects(t, resultRoleBinding.Object, "default")
	u = createClusterRoleBinding()
	resultClusterRoleBinding, err := client.Get(u)
	if err != nil {
		t.Errorf("Unable to get the cluster rolebinding %v", err)
	}
	checkSubjects(t, resultClusterRoleBinding.Object, OpenshiftMonitoringNamespace)
	// Make sure unrelated resources are not touched
	u = createService("activator-sm-service", "test")
	_, err = client.Get(u)
	if err != nil {
		t.Errorf("Unable to get the service %v", err)
	}
}

func checkSubjects(t *testing.T, object map[string]interface{}, ns string) {
	subjects, _, _ := unstructured.NestedFieldNoCopy(object, "subjects")
	subjs := subjects.([]interface{})
	m := subjs[0].(map[string]interface{})
	if m["namespace"] != ns {
		t.Errorf("Got %q, want %q", m["namespace"], ns)
	}
}

func createService(name string, ns string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetKind("Service")
	u.SetAPIVersion("v1")
	u.SetName(name)
	u.SetNamespace(ns)
	return u
}

func createRole(name string, ns string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetKind("Role")
	u.SetAPIVersion("rbac.authorization.k8s.io/v1")
	u.SetName(name)
	u.SetNamespace(ns)
	return u
}

func createClusterRole() *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetKind("ClusterRole")
	u.SetAPIVersion("rbac.authorization.k8s.io/v1")
	u.SetName(prometheusClusterRoleName)
	return u
}

func createRoleBinding(name string, ns string) *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetKind("RoleBinding")
	u.SetAPIVersion("rbac.authorization.k8s.io/v1")
	u.SetName(name)
	u.SetNamespace(ns)
	return u
}

func createClusterRoleBinding() *unstructured.Unstructured {
	u := &unstructured.Unstructured{}
	u.SetKind("ClusterRoleBinding")
	u.SetAPIVersion("rbac.authorization.k8s.io/v1")
	u.SetName(prometheusClusterRoleName + "-rb")
	return u
}
