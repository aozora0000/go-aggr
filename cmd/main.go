package main

import (
	"github.com/aozora0000/go-agqr"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	path, _ := filepath.Abs("temp.flv")
	client, err := agqr.New(path)
	if err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}
	defer client.Close()

	if err = client.Start(time.Minute*30, convert); err != nil {
		log.Fatalln(err.Error())
		os.Exit(1)
	}

}

func convert(path string) error {
	args := []string{
		"-y", "-i", path,
		"-vcodec", "copy", "-acodec", "copy", path + ".mp4",
	}
	cmd := exec.Command("ffmpeg", args...)
	defer func() {
		os.Remove(path)
	}()
	return cmd.Start()
}
