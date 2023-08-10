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

func MainPyContent(platform string) string {
	return Content("main.py", platform+"-python", "", "")
}

func RequirementsContent(platform string) string {
	return Content("requirements.txt", platform+"-python", "", "")
}

func PyProjectContent(botName string) string {
	return Content("pyproject.toml", "discord-python", botName, "")
}

func PythonTemplate(botName, platform, pm string) {
	pip := "pip3"
	pythonPath := "python3"

	if runtime.GOOS == "windows" {
		pip = "pip"
		pythonPath = "python"
	}

	_, err := looker.LookPath(pythonPath)
	pipPath, perr := looker.LookPath(pip)

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" python is not installed"))
	} else if perr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(fmt.Sprintf(" %s is not installed", pip)))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.py"), []byte(MainPyContent(platform)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, pm+".dockerfile", platform)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "python.md")), 0644)

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		installCmd := ""
		pkgs := ""
		shared_pkgs := "botway.py pyyaml"

		if platform == "discord" {
			pkgs = "discord.py pynacl " + shared_pkgs
		} else if platform == "telegram" {
			pkgs = "python-telegram-bot cryptography PySocks ujson " + shared_pkgs
		} else if platform == "slack" {
			pkgs = "slack-bolt " + shared_pkgs
		} else if platform == "twitch" {
			pkgs = "twitchio " + shared_pkgs
		}

		if pm == "pip" {
			requirementsFile := os.WriteFile(filepath.Join(botName, "requirements.txt"), []byte(RequirementsContent(platform)), 0644)

			if requirementsFile != nil {
				log.Fatal(requirementsFile)
			}

			installCmd = pipPath + " install -r requirements.txt"
		} else if pm == "pipenv" {
			pipenvPath, err := looker.LookPath("pipenv")

			if err != nil {
				fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
				fmt.Println(constants.FAIL_FOREGROUND.Render(" pipenv is not installed"))
			}

			installCmd = pipenvPath + " install " + pkgs
		} else if pm == "poetry" {
			poetryPath, err := looker.LookPath("poetry")

			if err != nil {
				fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
				fmt.Println(constants.FAIL_FOREGROUND.Render(" poetry is not installed"))
			}

			pyprojectFile := os.WriteFile(filepath.Join(botName, "pyproject.toml"), []byte(PyProjectContent(botName)), 0644)

			if pyprojectFile != nil {
				log.Fatal(pyprojectFile)
			}

			installCmd = poetryPath + " add " + pkgs
		}

		cmd := exec.Command("bash", "-c", installCmd)

		if runtime.GOOS == "windows" {
			cmd = exec.Command("powershell.exe", installCmd)
		}

		cmd.Dir = botName
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		CheckProject(botName, platform)
	}
}
