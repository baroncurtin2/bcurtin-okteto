package main

import (
	"fmt"
	kube "github.com/baroncurtin2/bcurtin-okteto/pkg/kube"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

// https://stackoverflow.com/questions/67543729/kubernetes-go-client-to-list-out-pod-details-similar-to-kubectl-get-pods
func main() {
	fmt.Println("Starting hello-world server...")

	http.HandleFunc("/", helloServer)

	// pods api endpoint
	http.HandleFunc("/podscount", podsCounter)
	http.HandleFunc("/podslist", podsDisplay)

	// metrics endpoint
	recordMetrics()
	http.Handle("/metrics", promhttp.Handler())

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

	fmt.Fprint(w, "The number of pods in the cluser is:", len(kubePods))
}

func podsDisplay(w http.ResponseWriter, r *http.Request) {
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

func recordMetrics() {
	go func() {
		for {
			kubeCount.Set(getKubeCount())
			time.Sleep(60 * time.Second) // refresh pod count every 60 seconds
		}
	}()

}

// https://prometheus.io/docs/guides/go-application/
var kubeCount = promauto.NewGauge(
	prometheus.GaugeOpts{
		Name: "myapp_pod_count_total",
		Help: "The total number of pods in the namespace",
	})

func getKubeCount() float64 {
	cfg := kube.GetInClusterConfig()
	client := kube.GetKubeClientset(cfg)
	pods := kube.GetPods(client, "baroncurtin2")
	kubePods := kube.CreateKubePods(pods)
	return float64(len(kubePods))
}
