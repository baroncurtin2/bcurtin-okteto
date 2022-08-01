package main

import (
	"fmt"
	kube "github.com/baroncurtin2/bcurtin-okteto/pkg/kube"
	"net/http"
)

// https://stackoverflow.com/questions/67543729/kubernetes-go-client-to-list-out-pod-details-similar-to-kubectl-get-pods
func main() {
	fmt.Println("Starting hello-world server...")

	http.HandleFunc("/", helloServer)

	// pods api endpoint
	http.HandleFunc("/pods", podsCounter)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func podsCounter(w http.ResponseWriter, r *http.Request) {
	cfg := kube.GetInClusterConfig()
	client := kube.GetKubeClientset(cfg)
	pods := kube.GetPods(client, "baroncurtin2")
	kubePods := kube.CreateKubePods(pods)

	sortBy := getSortBy(r)
	sortedPods := kube.SortKubePods(kubePods, sortBy)

	for _, pod := range sortedPods {
		fmt.Fprintln(w, pod.String())
	}
}

func getSortBy(r *http.Request) string {
	// sortBy is expected to look like field.orderdirection i. e. id.asc
	sortBy := r.URL.Query().Get("sortBy")
	if sortBy == "" {
		// id.asc is the default sort query
		sortBy = "name.asc"
	}

	return sortBy
}
