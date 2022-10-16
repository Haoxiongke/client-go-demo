package main

import (
	"context"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// config
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeDir)
	if err != nil {
		panic(err)
	}

	// client
	restClient, err := rest.RESTClientFor(config)
	if err != nil {
		panic(err)
	}

	// get data
	restClient.Get().Namespace("default").Resource("pods").Name("kube-proxy-zjm59").Do(context.TODO())
}
