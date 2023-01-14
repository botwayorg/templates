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

func MainRsContent(platform string) string {
	return Content("src/main.rs", platform+"-rust", "", "")
}

func CargoFileContent(botName, platform string) string {
	return Content("Cargo.toml", platform+"-rust", botName, "")
}

func RustTemplate(botName, platform, pm, hostService string) {
	_, err := looker.LookPath("cargo")
	pmPath, perr := looker.LookPath(pm)

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" cargo is not installed"))
	} else if perr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" " + pm + " is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.rs"), []byte(MainRsContent(platform)), 0644)
		cargoFile := os.WriteFile(filepath.Join(botName, "Cargo.toml"), []byte(CargoFileContent(botName, platform)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, pm+".dockerfile", platform)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "rust.md")), 0644)

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if cargoFile != nil {
			log.Fatal(cargoFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		pmBuild := pmPath + " build"
		buildCmd := exec.Command("bash", "-c", pmBuild)

		if runtime.GOOS == "windows" {
			buildCmd = exec.Command("powershell.exe", pmBuild)
		}

		buildCmd.Dir = botName
		buildCmd.Stdin = os.Stdin
		buildCmd.Stdout = os.Stdout
		buildCmd.Stderr = os.Stderr
		err = buildCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		if pm == "fleet" {
			rustUpPath, err := looker.LookPath("rustup")

			if err != nil {
				log.Printf("error: %v\n", err)
			}

			rustUpCmd := rustUpPath + " default nightly"

			rustUp := exec.Command("bash", "-c", rustUpCmd)

			if runtime.GOOS == "windows" {
				rustUp = exec.Command("powershell.exe", rustUpCmd)
			}

			rustUp.Dir = botName
			rustUp.Stdin = os.Stdin
			rustUp.Stdout = os.Stdout
			rustUp.Stderr = os.Stderr
			err = rustUp.Run()

			if err != nil {
				log.Printf("error: %v\n", err)
			}
		}

		CheckProject(botName, platform)
	}
}
