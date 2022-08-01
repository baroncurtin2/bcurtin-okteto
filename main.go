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

	fmt.Fprint(w, "The number of pods running in your current namespace: ", len(kubePods))
}
