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

func isTypescript(isTs bool) string {
	if isTs {
		return "nodejs-ts"
	} else {
		return "nodejs"
	}
}

func MainJSContent(platform string) string {
	return Content("main.js", platform+"-nodejs", "", "")
}

func MainTSContent(platform string) string {
	return Content("main.ts", platform+"-nodejs-ts", "", "")
}

func NodejsTemplate(botName, pm, platform, hostService string, isTs bool) {
	_, nerr := looker.LookPath("npm")
	pmPath, err := looker.LookPath(pm)

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" " + pm + " is not installed"))
	} else {
		if nerr != nil {
			fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
			fmt.Println(constants.FAIL_FOREGROUND.Render(" npm is not installed"))
		} else {
			mainFile := os.WriteFile(filepath.Join(botName, "src", "main.ts"), []byte(MainJSContent(platform)), 0644)
			dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, hostService, pm+".dockerfile", platform)), 0644)
			resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "nodejs.md")), 0644)
			packageFile := os.WriteFile(filepath.Join(botName, "package.json"), []byte(Content("package.json", platform+"-"+isTypescript(isTs), "", "")), 0644)

			if resourcesFile != nil {
				log.Fatal(resourcesFile)
			}

			if mainFile != nil {
				log.Fatal(mainFile)
			}

			if dockerFile != nil {
				log.Fatal(dockerFile)
			}

			if packageFile != nil {
				log.Fatal(packageFile)
			}

			if isTs {
				tsConfigFile := os.WriteFile(filepath.Join(botName, "tsconfig.json"), []byte(Content("tsconfig.json", platform+"-nodejs-ts", "", "")), 0644)

				if tsConfigFile != nil {
					log.Fatal(tsConfigFile)
				}
			}

			pmInstall := pmPath + " i"

			if pm == "yarn" {
				pmInstall = pmPath
			}

			installCmd := exec.Command("bash", "-c", pmInstall)

			if runtime.GOOS == "windows" {
				installCmd = exec.Command("powershell.exe", pmInstall)
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
}
