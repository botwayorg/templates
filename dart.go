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

func MainDartContent(platform string) string {
	return Content("src/main.dart", platform+"-dart", "", "")
}

func PubspecFileContent(botName, platform string) string {
	return Content("pubspec.yaml", platform+"-dart", botName, "")
}

func DartTemplate(botName, platform, hostService string) {
	dartPath, err := looker.LookPath("dart")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" dart is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.dart"), []byte(MainDartContent(platform)), 0644)
		pubspecFile := os.WriteFile(filepath.Join(botName, "pubspec.yaml"), []byte(PubspecFileContent(botName, platform)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "dart.dockerfile", platform)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "dart.md")), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if pubspecFile != nil {
			log.Fatal(pubspecFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		dartGet := dartPath + " pub get"

		getCmd := exec.Command("bash", "-c", dartGet)

		if runtime.GOOS == "windows" {
			getCmd = exec.Command("powershell.exe", dartGet)
		}

		getCmd.Dir = botName
		getCmd.Stdin = os.Stdin
		getCmd.Stdout = os.Stdout
		getCmd.Stderr = os.Stderr
		err = getCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		CheckProject(botName, platform)
	}
}
