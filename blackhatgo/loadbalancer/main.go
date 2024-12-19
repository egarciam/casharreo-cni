package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type LoadBalancer struct {
	Backends []*url.URL
	Current  uint64
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	target := lb.Backends[int(atomic.AddUint64(&lb.Current, 1))%len(lb.Backends)]
	fmt.Println(lb.Current)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

func NewLoadBalancer(backendUrls []string) *LoadBalancer {
	var backends []*url.URL
	for _, backendUrl := range backendUrls {
		url, err := url.Parse(backendUrl)
		if err != nil {
			log.Fatal(err)
		}
		backends = append(backends, url)
	}
	lb := &LoadBalancer{Backends: backends}
	fmt.Println(lb)
	return lb
	//return &LoadBalancer{Backends: backends}
}

func main() {
	backendUrls := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8088",
	}
	lb := NewLoadBalancer(backendUrls)

	fmt.Println("Load Balancer started at :8080")
	http.ListenAndServe(":8080", lb)
}
