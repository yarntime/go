package fish

import (
	k8s "k8s.io/client-go/kubernetes"
	v1 "k8s.io/client-go/pkg/api/v1"
	v1beta1 "k8s.io/client-go/pkg/apis/extensions/v1beta1"
	"k8s.io/client-go/pkg/labels"
	"k8s.io/client-go/tools/clientcmd"
)

type k8sClient struct {
	clientset *k8s.Clientset
}

func NewK8sClint() *k8sClient {

	config, err := clientcmd.BuildConfigFromFlags("kubernetes:8080", "")
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	client := &k8sClient{
		clientset: clientset,
	}
	return client
}

func (c *k8sClient) ListReplicationControllers(namespace string, listOption v1.ListOptions) (*v1.ReplicationControllerList, error) {
	return c.clientset.CoreV1().ReplicationControllers(namespace).List(listOption)
}

func (c *k8sClient) ListDeployments(namespace string, listOption v1.ListOptions) (*v1beta1.DeploymentList, error) {
	return c.clientset.Extensions().Deployments(namespace).List(listOption)
}

func (c *k8sClient) GetService(name string, namespace string) (*v1.Service, error) {
	return c.clientset.Services(namespace).Get(name)
}

func (c *k8sClient) ListPodsByService(service *v1.Service) (*v1.PodList, error) {
	lo := v1.ListOptions{
		LabelSelector: labels.Set(service.Spec.Selector).AsSelector().String(),
	}
	return c.clientset.Pods(service.Namespace).List(lo)
}
