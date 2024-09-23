// MFP - Miulti-Function Printers and scanners toolkit
// Device discovery
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Functions for management collections of endpoints

package discovery

import "sort"

// endpointsCmp compares two endpoint strings for sorting and searching.
func endpointsCmp(e1, e2 string) int {
	switch {
	case e1 < e2:
		return -1
	case e1 > e2:
		return 1
	}
	return 0
}

// endpointsContain reports if collection of endpoints contains
// the specified endpoint.
//
// endpoints assumed to be sorted.
func endpointsContain(endpoints []string, endpoint string) bool {
	_, found := sort.Find(len(endpoints), func(i int) int {
		return endpointsCmp(endpoint, endpoints[i])
	})
	return found
}

// endpointsAdd adds endpoint into the collection of endpoints.
//
// It assumes sorted endpoints on input and returns updated and
// properly sorted collection on output. It works by modifying
// the input collection in-place, so be aware.
//
// Additionally, it returns a bool flag that indicates that
// endpoint was actually added.
func endpointsAdd(endpoints []string, endpoint string) ([]string, bool) {
	i, found := sort.Find(len(endpoints), func(i int) int {
		return endpointsCmp(endpoint, endpoints[i])
	})

	if found {
		return endpoints, false
	}

	endpoints = append(endpoints, "")
	copy(endpoints[i+1:], endpoints[i:])
	endpoints[i] = endpoint

	return endpoints, true
}

// endpointsDel deletes endpoint from the collection of endpoints.
//
// It assumes sorted endpoints on input and returns updated and
// properly sorted collection on output. It works by modifying
// the input collection in-place, so be aware.
//
// Additionally, it returns a bool flag that indicates that
// endpoint was actually found in the collection and deleted.
func endpointsDel(endpoints []string, endpoint string) ([]string, bool) {
	i, found := sort.Find(len(endpoints), func(i int) int {
		return endpointsCmp(endpoint, endpoints[i])
	})

	if !found {
		return endpoints, false
	}

	copy(endpoints[i:], endpoints[i+1:])
	endpoints = endpoints[:len(endpoints)-1]

	return endpoints, true
}

// endpointsMerge merges two collections of endpoints.
//
// Input collections must be sorted and returned collection
// is sorted as well.
func endpointsMerge(endpoints1, endpoints2 []string) []string {
	out := make([]string, 0, len(endpoints1)+len(endpoints2))

	for len(endpoints1) > 0 && len(endpoints2) > 0 {
		cmp := endpointsCmp(endpoints1[0], endpoints2[0])
		switch {
		case cmp < 0:
			out = append(out, endpoints1[0])
			endpoints1 = endpoints1[1:]
		case cmp > 0:
			out = append(out, endpoints2[0])
			endpoints2 = endpoints2[1:]
		default:
			out = append(out, endpoints1[0])
			endpoints1 = endpoints1[1:]
			endpoints2 = endpoints2[1:]
		}
	}

	out = append(out, endpoints1...)
	out = append(out, endpoints2...)

	return out
}
