package main

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/kubernetes"
)

func getBackup(namespace string, clientset *kubernetes.Clientset) []HelmRelease {
	releaseMap := []HelmRelease{}
	labelSelector := metav1.LabelSelector{MatchLabels: map[string]string{"owner": "helm"}}

	listOptions := metav1.ListOptions{
		LabelSelector: labels.Set(labelSelector.MatchLabels).String(),
	}

	secrets, err := clientset.CoreV1().Secrets(namespace).List(context.TODO(), listOptions)
	if err != nil {
		ErrorLogger.Println(err)
	}

	for _, item := range secrets.Items {

		r := HelmRelease{
			status:  item.ObjectMeta.Labels["status"],
			version: item.ObjectMeta.Labels["version"],
			content: string(item.Data["release"]),
			name:    item.ObjectMeta.Labels["name"],
		}
		releaseMap = append(releaseMap, r)
	}
	return releaseMap
}
