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

func MainPHPContent(platform string) string {
	return Content("src/main.php", platform+"-php", "", "")
}

func BotwayPHPContent() string {
	return Content("packages/bw-php/main.php", "botway", "", "")
}

func ComposerFileContent(botName, platform string) string {
	return Content("composer.json", platform+"-php", botName, "")
}

func PHPTemplate(botName, platform string) {
	_, err := looker.LookPath("php")
	composerPath, serr := looker.LookPath("composer")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" php is not installed"))
	} else if serr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" composer is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.php"), []byte(MainPHPContent(platform)), 0644)
		botwayFile := os.WriteFile(filepath.Join(botName, "src", "botway.php"), []byte(BotwayPHPContent()), 0644)
		composerFile := os.WriteFile(filepath.Join(botName, "composer.json"), []byte(ComposerFileContent(botName, platform)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "php.dockerfile", platform)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "php.md")), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if botwayFile != nil {
			log.Fatal(botwayFile)
		}

		if composerFile != nil {
			log.Fatal(composerFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		composerInstall := composerPath + " install"

		installCmd := exec.Command("bash", "-c", composerInstall)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", composerInstall)
		}

		installCmd.Dir = botName
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		err = installCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		CheckProject(botName, "discord")
	}
}
