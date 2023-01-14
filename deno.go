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

func DenoTemplate(botName, platform, hostService string) {
	if platform == "discord" {
		if err := os.Mkdir(filepath.Join(botName, "src", "commands"), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(botName, "src", "events"), os.ModePerm); err != nil {
			log.Fatal(err)
		}

		if err := os.Mkdir(filepath.Join(botName, "src", "utils"), os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	deno, err := looker.LookPath("deno")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" deno is not installed"))
	} else {
		mainFile := os.WriteFile(filepath.Join(botName, "main.ts"), []byte(DenoMainTsContent(platform)), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, "deno.dockerfile", platform)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "deno.md")), 0644)
		fileToCache := "deps.ts"

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		if platform == "discord" {
			commandsModTsFile := os.WriteFile(filepath.Join(botName, "src", "commands", "mod.ts"), []byte(CommandsModTsContent()), 0644)
			commandsPingTsFile := os.WriteFile(filepath.Join(botName, "src", "commands", "ping.ts"), []byte(CommandsPingTsContent()), 0644)
			eventsGuildCreateTsFile := os.WriteFile(filepath.Join(botName, "src", "events", "guildCreate.ts"), []byte(EventsGuildCreateTsContent()), 0644)
			eventsInteractionCreateTsFile := os.WriteFile(filepath.Join(botName, "src", "events", "interactionCreate.ts"), []byte(EventsInteractionCreateTsContent()), 0644)
			eventsModTsFile := os.WriteFile(filepath.Join(botName, "src", "events", "mod.ts"), []byte(EventsModTsContent()), 0644)
			eventsReadyTsFile := os.WriteFile(filepath.Join(botName, "src", "events", "ready.ts"), []byte(EventsReadyTsContent()), 0644)
			utilsHelpersTsFile := os.WriteFile(filepath.Join(botName, "src", "utils", "helpers.ts"), []byte(UtilsHelpersTsContent()), 0644)
			utilsLoggerTsFile := os.WriteFile(filepath.Join(botName, "src", "utils", "logger.ts"), []byte(UtilsLoggerTsContent()), 0644)

			if commandsModTsFile != nil {
				log.Fatal(commandsModTsFile)
			}

			if commandsPingTsFile != nil {
				log.Fatal(commandsPingTsFile)
			}

			if eventsGuildCreateTsFile != nil {
				log.Fatal(eventsGuildCreateTsFile)
			}

			if eventsInteractionCreateTsFile != nil {
				log.Fatal(eventsInteractionCreateTsFile)
			}

			if eventsModTsFile != nil {
				log.Fatal(eventsModTsFile)
			}

			if eventsReadyTsFile != nil {
				log.Fatal(eventsReadyTsFile)
			}

			if utilsHelpersTsFile != nil {
				log.Fatal(utilsHelpersTsFile)
			}

			if utilsLoggerTsFile != nil {
				log.Fatal(utilsLoggerTsFile)
			}
		}

		if platform == "twitch" {
			loggerFile := os.WriteFile(filepath.Join(botName, "logger.ts"), []byte(LoggerTsContent()), 0644)

			if loggerFile != nil {
				log.Fatal(loggerFile)
			}
		}

		if platform != "telegram" {
			depsFile := os.WriteFile(filepath.Join(botName, "deps.ts"), []byte(DepsTsContent(platform)), 0644)

			if depsFile != nil {
				log.Fatal(depsFile)
			}
		}

		if platform == "telegram" {
			fileToCache = "main.ts"
		}

		if platform != "discord" {
			os.RemoveAll("src")
		}

		denoInstall := deno + " cache " + fileToCache

		installCmd := exec.Command("bash", "-c", denoInstall)

		if runtime.GOOS == "windows" {
			installCmd = exec.Command("powershell.exe", denoInstall)
		}

		installCmd.Dir = botName
		installCmd.Stdin = os.Stdin
		installCmd.Stdout = os.Stdout
		installCmd.Stderr = os.Stderr
		err = installCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		// format files
		denoFormat := deno + " fmt"

		formatCmd := exec.Command("bash", "-c", denoFormat)

		if runtime.GOOS == "windows" {
			formatCmd = exec.Command("powershell.exe", denoFormat)
		}

		formatCmd.Dir = botName
		formatCmd.Stdin = os.Stdin
		formatCmd.Stdout = os.Stdout
		formatCmd.Stderr = os.Stderr
		err = formatCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		CheckProject(botName, platform)
	}
}
