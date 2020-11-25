package main

import (
	"k8s.io/client-go/tools/clientcmd"
	"flag"
	"fmt"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/api/core/v1"
)

func main() {
	//var kubeconfig string
	//var master string
	//flag.StringVar(&kubeconfig, "kubeconfig", "", "absolute path to the kubeconfig file")
	//flag.StringVar(&master, "master", "", "master url")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", "root/.kube/config")
	if err != nil {
		fmt.Println("get config failed!, ", err)
	}

	config.APIPath = "api"
	config.GroupVersion  = &v1.SchemeGroupVersion
	config.NegotiatedSerializer = scheme.Codecs

	client, err := rest.RESTClientFor(config)
	if err != nil {
		fmt.Println("get rest client failed!", err)
	}

	podlist := v1.PodList{}
	err = client.Get().Resource("pods").Namespace("kube-system").Do().Into(&podlist)
	if err != nil {
		fmt.Println("get pod list failed!, ", err)
	}
	fmt.Println("pod list is ", podlist)
}
