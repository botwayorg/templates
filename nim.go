package templates

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/abdfnx/botway/constants"
	"github.com/abdfnx/looker"
)

func MainNimContent(platform string) string {
	return Content("src/main.nim", platform+"-nim", "", "")
}

func BotnimContent(botName string) string {
	return Content("packages/botnim/main.nim", "botway", botName, "")
}

func PngFileContent() string {
	return Content("assets/facepalm.png", "discord-nim", "", "")
}

func NimbleFileContent(platform string) string {
	return Content(platform+"_nim.nimble", platform+"-nim", "", "")
}

func NimTemplate(botName, platform string) {
	_, err := looker.LookPath("nim")
	nimblePath, nerr := looker.LookPath("nimble")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" nim is not installed"))
	} else if nerr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" nimble is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.nim"), []byte(MainNimContent(platform)), 0644)
		botnimFile := os.WriteFile(filepath.Join(botName, "src", "botnim.nim"), []byte(BotnimContent(botName)), 0644)
		nimbleFile := os.WriteFile(filepath.Join(botName, botName+".nimble"), []byte(NimbleFileContent(platform)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "nim.dockerfile", platform)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "nim.md")), 0644)

		if platform == "discord" {
			if err := os.Mkdir(filepath.Join(botName, "assets"), os.ModePerm); err != nil {
				log.Fatal(err)
			}

			pngFile := os.WriteFile(filepath.Join(botName, "assets", "facepalm.png"), []byte(PngFileContent()), 0644)

			if pngFile != nil {
				log.Fatal(pngFile)
			}
		}

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if botnimFile != nil {
			log.Fatal(botnimFile)
		}

		if nimbleFile != nil {
			log.Fatal(nimbleFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		nimbleInstall := nimblePath + " install -y"

		installCmd := exec.Command("bash", "-c", nimbleInstall)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", nimbleInstall)
		}

		installCmd.Dir = botName
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		err = installCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		CheckProject(botName, platform)
	}
}
