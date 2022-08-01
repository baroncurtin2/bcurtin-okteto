package kube

import "time"

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
