package main

import (
	"fmt"
	"os"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func configHandler(contextToUse string, kubeconfig string) *kubernetes.Clientset {
	var filename string
	var kubeconfigPath string
	dirname, err := os.UserHomeDir()

	if kubeconfig == "" {
		filename = ".kube/config"
		kubeconfigPath = fmt.Sprintf("%s/%s", dirname, filename)

	}

	WarningLogger.Println("The obtained kubeconfig path is %s", kubeconfigPath)

	config, err := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeconfigPath},
		&clientcmd.ConfigOverrides{
			CurrentContext: contextToUse,
		}).ClientConfig()

	if err != nil {
		panic(err.Error())
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}
