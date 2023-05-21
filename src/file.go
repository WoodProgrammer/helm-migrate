package main

import "io/ioutil"

func backupFile(fileName string, content []byte) {

	err := ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		panic(err)
	}
}
