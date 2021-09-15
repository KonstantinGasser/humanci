package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/KonstantinGasser/humanci"
)

func main() {

	cli := humanci.New()

	cli.RootNOP("who", "what").
		WithNOP("is").
		WithRegex(HandleAction, "listening|receiving").
		WithNOP("on", "at").
		WithNOP("port").
		WithRegex(HandlePort, "[0-9]{2,4}")

	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}

func HandleAction(meta humanci.MetaData) error {
	switch meta.LastToken() {
	case "listening":
		meta.Value("cmd", "lsof")
		meta.Value("type", "LISTENING")
	case "receiving":
		meta.Value("cmd", "lsof")
		meta.Value("type", "RECEIVING")
	}
	return nil
}

func HandlePort(meta humanci.MetaData) error {
	port := meta.LastToken()

	switch meta.Return("type") {
	case "LISTENING":
		cmd := exec.Command(meta.Return("cmd").(string), "-i", "tcp:"+port)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	case "RECEIVING":
		cmd := exec.Command(meta.Return("cmd").(string), "-nPi", "tcp:"+port)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			return err
		}
	default:
		fmt.Println("Mhm..sorry I am not sure what to do with this.")
	}
	return nil
}
