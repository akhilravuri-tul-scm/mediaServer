package audioconverter

import (
	"os/exec"
)

func WavToMp3(inputPath, outputPath string) error {
	cmd := exec.Command("lame", inputPath, outputPath)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
