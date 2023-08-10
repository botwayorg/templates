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

func MainCContent() string {
	return Content("src/main.c", "discord-c", "", "")
}

func BWCContent(botName string) string {
	return Content("packages/bwc/main.h", "botway", botName, "")
}

func CRunPsFileContent() string {
	return Content("run.ps1", "discord-c", "", "")
}

func CTemplate(botName string) {
	_, err := looker.LookPath("gcc")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" gcc is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.c"), []byte(MainCContent()), 0644)
		botwayHeaderFile := os.WriteFile(filepath.Join(botName, "src", "botway.h"), []byte(BWCContent(botName)), 0644)
		runPsFile := os.WriteFile(filepath.Join(botName, "run.ps1"), []byte(CRunPsFileContent()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "c-discord.dockerfile", "discord")), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources("discord", "c.md")), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if botwayHeaderFile != nil {
			log.Fatal(botwayHeaderFile)
		}

		if runPsFile != nil {
			log.Fatal(runPsFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		installScript := "https://raw.githubusercontent.com/botwayorg/discord-c/main/scripts/install-concord."

		getConcord := exec.Command("bash", "-c", "curl -sL "+installScript+".sh | bash")

		if runtime.GOOS == "windows" {
			getConcord = exec.Command("powershell.exe", "irm "+installScript+".ps1 | iex")
		}

		getConcord.Dir = botName
		getConcord.Stdin = os.Stdin
		getConcord.Stdout = os.Stdout
		getConcord.Stderr = os.Stderr
		err = getConcord.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		run := exec.Command("bash", "-c", "gcc src/main.c -o bot -pthread -ldiscord -lcurl")

		if runtime.GOOS == "windows" {
			run = exec.Command("powershell.exe", "./run.ps1")
		}

		run.Dir = botName
		run.Stdin = os.Stdin
		run.Stdout = os.Stdout
		run.Stderr = os.Stderr
		err = run.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		CheckProject(botName, "discord")
	}
}
