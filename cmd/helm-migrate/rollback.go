package main

import (
	"log"
	"os"

	"helm.sh/helm/v3/pkg/action"
	"k8s.io/cli-runtime/pkg/genericclioptions"
)

func rollBack(release string, releaseVersion int, namespace string, targetClusterContext string) error {

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(
		&genericclioptions.ConfigFlags{
			Namespace: &namespace,
			Context:   &targetClusterContext,
		},
		namespace,
		os.Getenv("HELM_DRIVER"),
		log.Printf,
	); err != nil {
		return err
	}

	client := action.NewRollback(actionConfig)

	client.Version = releaseVersion

	err := client.Run(release)
	if err != nil {
		ErrorLogger.Println(err)
		return err
	}

	return nil
}
