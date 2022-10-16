package main

import (
	"context"
	"fmt"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// RestClient
	//// config
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	//if err != nil {
	//	panic(err)
	//}
	//
	//config.GroupVersion = &v1.SchemeGroupVersion
	//config.NegotiatedSerializer = scheme.Codecs
	//config.APIPath = "/api"
	//
	//// client
	//restClient, err := rest.RESTClientFor(config)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// get data
	//pod := v1.Pod{}
	//err = restClient.Get().Namespace("default").Resource("pods").Name("nginx-8f458dc5b-tkmnr").Do(context.TODO()).Into(&pod)
	//if err != nil {
	//	println(err)
	//} else {
	//	fmt.Println(pod.Name)
	//}

	// clientset
	//config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	//if err != nil {
	//	panic(err)
	//}
	//clientset, err := kubernetes.NewForConfig(config)
	//coreV1 := clientset.CoreV1()
	//pod, err := coreV1.Pods("default").Get(context.TODO(), "nginx-8f458dc5b-tkmnr", v1.GetOptions{})
	//if err != nil {
	//	panic(err)
	//} else {
	//	println(pod.Name)
	//}

	// DynamicCLient
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		panic(err)
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	gvr := schema.GroupVersionResource{Version: "v1", Resource: "pods"}
	pod, err := dynamicClient.Resource(gvr).Namespace("default").Get(context.TODO(), "nginx-8f458dc5b-tkmnr", v1.GetOptions{})
	if err != nil {
		panic(err)
	} else {
		fmt.Println(pod.GetName())
	}
}
