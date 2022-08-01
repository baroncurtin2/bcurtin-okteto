package kube

import (
	"fmt"
	"time"
)

type KubePod struct {
	name     string
	restarts int32
	age      time.Duration
}

func NewKubePod(name string, restarts int32, age time.Duration) *KubePod {
	return &KubePod{
		name:     name,
		restarts: restarts,
		age:      age,
	}
}

func (k *KubePod) String() string {
	return fmt.Sprintf("Name: %s | Restarts: %d | Age: %d", k.name, k.restarts, k.age)
}

// name sorts
type sortKubeByName []KubePod

func (s sortKubeByName) Len() int { return len(s) }
func (s sortKubeByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortKubeByName) Less(i, j int) bool {
	return s[i].name < s[j].name
}

// restart sorts
type sortKubeByRestarts []KubePod

func (s sortKubeByRestarts) Len() int { return len(s) }
func (s sortKubeByRestarts) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortKubeByRestarts) Less(i, j int) bool {
	return s[i].restarts < s[j].restarts
}

// age sorts
type sortKubeByAge []KubePod

func (s sortKubeByAge) Len() int { return len(s) }
func (s sortKubeByAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s sortKubeByAge) Less(i, j int) bool {
	return s[i].age < s[j].age
}
