package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

func dump(filename string, value []HelmRelease) {
	file, _ := os.Create(filename)

	enc := gob.NewEncoder(file)

	if err := enc.Encode(value); err != nil {
		log.Fatal(err)
	}
}

func encodedBackup(filename string) []byte {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	stats, statsErr := file.Stat()
	if statsErr != nil {
		fmt.Println(statsErr)
	}

	var size int64 = stats.Size()
	bytes := make([]byte, size)

	bufr := bufio.NewReader(file)
	_, err = bufr.Read(bytes)

	return bytes
}

func decodeFromBackup(filename string) error {

	data := encodedBackup(filename)

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)

	m := []HelmRelease{}

	if err := dec.Decode(&m); err != nil {
		log.Fatal(err)
		return err
	}

	return nil

}
