package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func download(v string) error {
	cmd := exec.Command("java", "-version")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println(out.String())
	return nil
}
