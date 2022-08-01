package kube

import "testing"

func TestNewKubePod(t *testing.T) {
	tests := []struct {
		name     string
		kube     KubePod
		expected KubePod
	}{
		{
			name: "Brand new kube",
			kube: KubePod{
				name:     "a",
				restarts: 0,
				age:      0,
			},
			expected: KubePod{
				name:     "a",
				restarts: 0,
				age:      0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kube := tt.kube

			if kube != tt.expected {
				t.Errorf("Kube = %v; expected %v", kube, tt.expected)
			}
		})
	}
}

func TestKubeString(t *testing.T) {
	tests := []struct {
		name     string
		kube     KubePod
		expected string
	}{
		{
			name: "KubePod String Method Test 1",
			kube: KubePod{
				name:     "a",
				restarts: 0,
				age:      0,
			},
			expected: "Name: a | Restarts: 0 | Age: 0",
		},
		{
			name: "KubePod String Method Test 2",
			kube: KubePod{
				name:     "b",
				restarts: 100,
				age:      5,
			},
			expected: "Name: b | Restarts: 100 | Age: 5",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kube := tt.kube

			if kube.String() != tt.expected {
				t.Errorf("Kube = %v; expected %v", kube.String(), tt.expected)
			}
		})
	}
}

func TestKubeSorting(t *testing.T) {
	tests := []struct {
		name     string
		sortBy   string
		kubes    []KubePod
		expected []KubePod
	}{
		{
			name:   "Kube Name Sort",
			sortBy: "name.asc",
			kubes: []KubePod{
				KubePod{
					name:     "b",
					restarts: 0,
					age:      1,
				},
				KubePod{
					name:     "a",
					restarts: 0,
					age:      2,
				},
			},
			expected: []KubePod{
				KubePod{
					name:     "a",
					restarts: 0,
					age:      2,
				},
				KubePod{
					name:     "b",
					restarts: 0,
					age:      1,
				},
			},
		},
		{
			name:   "Kube Name Sort Desc",
			sortBy: "name.desc",
			kubes: []KubePod{
				KubePod{
					name:     "a",
					restarts: 0,
					age:      2,
				},
				KubePod{
					name:     "b",
					restarts: 0,
					age:      1,
				},
			},
			expected: []KubePod{
				KubePod{
					name:     "b",
					restarts: 0,
					age:      1,
				},
				KubePod{
					name:     "a",
					restarts: 0,
					age:      2,
				},
			},
		},
		{
			name:   "Kube Restarts Sort",
			sortBy: "restarts.asc",
			kubes: []KubePod{
				KubePod{
					name:     "a",
					restarts: 5,
					age:      1,
				},
				KubePod{
					name:     "b",
					restarts: 3,
					age:      2,
				},
			},
			expected: []KubePod{
				KubePod{
					name:     "b",
					restarts: 3,
					age:      2,
				},
				KubePod{
					name:     "a",
					restarts: 5,
					age:      1,
				},
			},
		},
		{
			name:   "Kube Restarts Sort Desc",
			sortBy: "restarts.desc",
			kubes: []KubePod{
				KubePod{
					name:     "b",
					restarts: 3,
					age:      2,
				},
				KubePod{
					name:     "a",
					restarts: 5,
					age:      1,
				},
			},
			expected: []KubePod{
				KubePod{
					name:     "a",
					restarts: 5,
					age:      1,
				},
				KubePod{
					name:     "b",
					restarts: 3,
					age:      2,
				},
			},
		},
		{
			name:   "Kube Age Sort",
			sortBy: "age.asc",
			kubes: []KubePod{
				KubePod{
					name:     "a",
					restarts: 0,
					age:      2,
				},
				KubePod{
					name:     "b",
					restarts: 0,
					age:      1,
				},
			},
			expected: []KubePod{
				KubePod{
					name:     "b",
					restarts: 0,
					age:      1,
				},
				KubePod{
					name:     "a",
					restarts: 0,
					age:      2,
				},
			},
		},
		{
			name:   "Kube Age Sort Desc",
			sortBy: "age.asc",
			kubes: []KubePod{
				KubePod{
					name:     "b",
					restarts: 0,
					age:      1,
				},
				KubePod{
					name:     "a",
					restarts: 0,
					age:      2,
				},
			},
			expected: []KubePod{
				KubePod{
					name:     "a",
					restarts: 0,
					age:      2,
				},
				KubePod{
					name:     "b",
					restarts: 0,
					age:      1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kubes := tt.kubes
			sortBy := tt.sortBy
			sortedKubes := SortKubePods(kubes, sortBy)

			for i, pod := range sortedKubes {
				if pod.name != tt.expected[i].name &&
					pod.restarts != tt.expected[i].restarts &&
					pod.age != tt.expected[i].age {
					t.Errorf("Kube = %v; expected %v", pod, tt.expected[i])
				}
			}
		})
	}
}
