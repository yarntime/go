/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"

	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/pkg/labels"
)

var (
	kubeconfig = flag.String("kubeconfig", "./config", "absolute path to the kubeconfig file")
)

const (
	ReplicationController = "replicationcontrollers"
	Deployment            = "deployments"
)

type Metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"name"`
	Type      string `json:"type"`

	Labels map[string]string `json:"labels,omitempty"`

	Annotations map[string]string `json:"annotations,omitempty"`

	DesiredReplicas   *int32 `json:"desiredreplicas"`
	AvailableReplicas int32  `json:"availablereplicas"`
}

type App struct {
	Metadata `json:"matadata"`
	Service  *v1.Service `json:"service"`
	Pods     *v1.PodList `json:"pods"`
}

type ReleaseResources struct {
	apps []App `json:"apps"`
}

func main() {
	fmt.Printf("starting\n")

	resources := ReleaseResources{}

	name := "release-test"
	namespace := "default"
	labelSelector := labels.Set(map[string]string{"release": name}).AsSelector()
	listOptions := v1.ListOptions{
		LabelSelector: labelSelector.String(),
	}

	c := client.NewK8sClint()

	// add replication controllers to resources
	replicationControllers, _ := c.ListReplicationControllers(namespace, listOptions)
	fmt.Printf("rc length is %d\n", len(replicationControllers.Items))
	for i := 0; i < len(replicationControllers.Items); i++ {
		rc := replicationControllers.Items[i]
		app := App{
			Metadata: Metadata{
				Name:              rc.Name,
				Namespace:         rc.Namespace,
				Type:              ReplicationController,
				DesiredReplicas:   rc.Spec.Replicas,
				AvailableReplicas: rc.Status.AvailableReplicas,
				Labels:            rc.Labels,
				Annotations:       rc.Annotations,
			},
		}
		service, _ := c.GetService(rc.Name, rc.Namespace)
		app.Service = service

		pods, _ := c.ListPodsByService(service)
		app.Pods = pods

		resources.apps = append(resources.apps, app)
	}

	// add deployments to resources
	deployments, _ := c.ListDeployments(namespace, listOptions)
	fmt.Printf("dep length is %d\n", len(deployments.Items))
	for i := 0; i < len(deployments.Items); i++ {
		dep := deployments.Items[i]
		app := App{
			Metadata: Metadata{
				Name:              dep.Name,
				Namespace:         dep.Namespace,
				Type:              Deployment,
				DesiredReplicas:   dep.Spec.Replicas,
				AvailableReplicas: dep.Status.AvailableReplicas,
				Labels:            dep.Labels,
				Annotations:       dep.Annotations,
			},
		}
		service, _ := c.GetService(dep.Name, dep.Namespace)
		app.Service = service

		pods, _ := c.ListPodsByService(service)
		app.Pods = pods

		resources.apps = append(resources.apps, app)
	}

	fmt.Printf("%v\n", resources)
}
