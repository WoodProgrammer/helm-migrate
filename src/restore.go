package main

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

func restoreBackup(clientset *kubernetes.Clientset, releaseMap []HelmRelease) {
	for _, release := range releaseMap {
		secretName := fmt.Sprintf("sh.helm.release.v1.%s.v%s", release.name, release.version)

		/*base64EncodedData := make([]byte, b64.StdEncoding.EncodedLen(len(release.content)))
		b64.StdEncoding.Encode(base64EncodedData, []byte(release.content))*/

		secret := corev1.Secret{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Secret",
				APIVersion: "v1",
			},

			Type: "helm.sh/release.v1",
			ObjectMeta: metav1.ObjectMeta{
				Name:      secretName,
				Namespace: "kube-system",
				Labels: map[string]string{
					"owner":   "helm",
					"name":    release.name,
					"status":  release.status,
					"version": release.version,
				},
			},
			Data: map[string][]byte{"release": []byte(release.content)},
		}
		result, err := clientset.CoreV1().Secrets("kube-system").Create(context.TODO(), &secret, metav1.CreateOptions{})

		if err != nil {
			fmt.Println("ERR::", err)
		} else {
			fmt.Println(result)
		}
	}
}
