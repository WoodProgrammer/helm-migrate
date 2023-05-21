package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func restoreBackup(namespace string, clientset *kubernetes.Clientset, releaseMap []HelmRelease) {
	for _, release := range releaseMap {
		secretName := fmt.Sprintf("sh.helm.release.v1.%s.v%s", release.name, release.version)

		secret := corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},

			Type: "helm.sh/release.v1",
			ObjectMeta: metav1.ObjectMeta{
				Name:      secretName,
				Namespace: namespace,
				Labels: map[string]string{
					"owner":   "helm",
					"name":    release.name,
					"status":  release.status,
					"version": release.version,
				},
			},
			Data: map[string][]byte{"release": []byte(release.content)},
		}
		_, err := clientset.CoreV1().Secrets(namespace).Create(context.TODO(), &secret, metav1.CreateOptions{})

		if err != nil {
			ErrorLogger.Println(err)
		}
	}
}
