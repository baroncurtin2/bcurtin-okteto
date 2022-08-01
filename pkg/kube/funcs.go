package kube

import (
	"context"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sort"
	"time"
)

// GetInClusterConfig returns the Kubernetes cluster configuration
func GetInClusterConfig() *rest.Config {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	return config
}

// GetKubeClientset returns a Kubernetes Clientset
func GetKubeClientset(cfg *rest.Config) *kubernetes.Clientset {
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

// GetPods returns a list of pods in the cluster
func GetPods(cs *kubernetes.Clientset, ns string) *v1.PodList {
	pods, err := cs.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		panic(err.Error())
	}

	return pods
}

func CreateKubePods(podList *v1.PodList) []KubePod {
	kubePods := make([]KubePod, 0, len(podList.Items))

	for _, pod := range podList.Items {
		podCreationTime := pod.GetCreationTimestamp()
		podStatus := pod.Status
		var restarts int32

		name := pod.GetName()
		age := time.Since(podCreationTime.Time).Round(time.Second)

		for container := range pod.Spec.Containers {
			restarts += podStatus.ContainerStatuses[container].RestartCount
		}

		kube := NewKubePod(name, restarts, age)
		kubePods = append(kubePods, *kube)

	}

	return kubePods
}

func SortKubePods(kubes []KubePod, sortBy string) []KubePod {
	sortedPods := kubes

	switch sortBy {
	case "name.asc":
		sort.Sort(sortKubeByName(sortedPods))
	case "age.asc":
		sort.Sort(sortKubeByAge(sortedPods))
	case "restarts.asc":
		sort.Sort(sortKubeByRestarts(sortedPods))
	case "name.desc":
		sort.Sort(sort.Reverse(sortKubeByName(sortedPods)))
	case "age.desc":
		sort.Sort(sort.Reverse(sortKubeByAge(sortedPods)))
	case "restart.desc":
		sort.Sort(sort.Reverse(sortKubeByRestarts(sortedPods)))
	}

	return sortedPods
}
