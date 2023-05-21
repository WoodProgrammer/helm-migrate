package main

func main() {

	sourceClusterclientset := configHandler("kind-kind")
	targetClusterclientset := configHandler("kind-kind-test-cluster")

	backup := getBackup(sourceClusterclientset)

	restoreBackup(targetClusterclientset, backup)

}
