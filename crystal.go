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

func MainCrContent() string {
	return Content("src/main.cr", "discord-crystal", "", "")
}

func ShardFileContent(botName string) string {
	return Content("shard.yml", "discord-crystal", botName, "")
}

func CrystalTemplate(botName string) {
	_, err := looker.LookPath("crystal")
	shardsPath, serr := looker.LookPath("shards")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" crystal is not installed"))
	} else if serr != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" shards is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "src", "main.cr"), []byte(MainCrContent()), 0644)
		shardFile := os.WriteFile(filepath.Join(botName, "shard.yml"), []byte(ShardFileContent(botName)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "crystal.dockerfile", "discord")), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources("discord", "crystal.md")), 0644)

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if shardFile != nil {
			log.Fatal(shardFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		shardsInstall := shardsPath + " install"

		installCmd := exec.Command("bash", "-c", shardsInstall)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", shardsInstall)
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
