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

func MainCsContent(platform string) string {
	return Content("src/Main.cs", platform+"-csharp", "", "")
}

func BotCSharpProj(platform string) string {
	return Content(platform+"-csharp.csproj", platform+"-csharp", "", "")
}

func CsharpTemplate(botName, platform string) {
	dotnetPath, err := looker.LookPath("dotnet")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" dotnet is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "Main.cs"), []byte(MainCsContent(platform)), 0644)
		csprojFile := os.WriteFile(filepath.Join(botName, botName+".csproj"), []byte(BotCSharpProj(platform)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "csharp.dockerfile", platform)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "csharp.md")), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if csprojFile != nil {
			log.Fatal(csprojFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		dotNetRestore := dotnetPath + " restore"

		restoreCmd := exec.Command("bash", "-c", dotNetRestore)

		if runtime.GOOS == "windows" {
			restoreCmd = exec.Command("powershell.exe", dotNetRestore)
		}

		restoreCmd.Dir = botName
		restoreCmd.Stdin = os.Stdin
		restoreCmd.Stdout = os.Stdout
		restoreCmd.Stderr = os.Stderr
		err = restoreCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		CheckProject(botName, platform)
	}
}
