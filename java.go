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

func JavaTemplate(botName, platform, hostService string) {
	createDirs(botName, "java", platform)

	gradle, err := looker.LookPath("gradle")

	if err != nil {
		fmt.Print(constants.FAIL_BACKGROUND.Render("ERROR"))
		fmt.Println(constants.FAIL_FOREGROUND.Render(" gradle is not wrappered"))
	} else {
		botlinFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "java", "core", "Botway.kt"), []byte(BotlinContent()), 0644)
		mainFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "java", "core", "Bot.java"), []byte(MainJavaContent(platform)), 0644)
		buildGradleFile := os.WriteFile(filepath.Join(botName, "app", "build.gradle"), []byte(BuildGradleContent(platform)), 0644)
		gradleWrapperPropsFile := os.WriteFile(filepath.Join(botName, "gradle", "wrapper", "gradle-wrapper.properties"), []byte(GradleWrapperPropsContent()), 0644)
		gradlewFile := os.WriteFile(filepath.Join(botName, "gradlew"), []byte(GradlewContent()), 0644)
		gradlewBatFile := os.WriteFile(filepath.Join(botName, "gradlew.bat"), []byte(GradlewBatContent()), 0644)
		settingsFile := os.WriteFile(filepath.Join(botName, "settings.gradle"), []byte(SettingsGradle()), 0644)
		dockerFile := os.WriteFile(filepath.Join(botName, "Dockerfile"), []byte(DockerfileContent(botName, hostService, "gradle.dockerfile", platform)), 0644)
		resourcesFile := os.WriteFile(filepath.Join(botName, "resources.md"), []byte(Resources(platform, "java.md")), 0644)
		gitattributesFile := os.WriteFile(filepath.Join(botName, ".gitattributes"), []byte(DotGitattributesContent()), 0644)

		if platform == "telegram" {
			botHandlerFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "java", "core", "BotHandler.java"), []byte(BotHandlerContent()), 0644)
			tgBotFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "java", "core", "TGBot.java"), []byte(TGBotContent()), 0644)

			if botHandlerFile != nil {
				log.Fatal(botHandlerFile)
			}

			if tgBotFile != nil {
				log.Fatal(tgBotFile)
			}
		}

		if platform == "twitch" {
			startFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "java", "core", "Start.java"), []byte(StartJavaContent()), 0644)
			cnodFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "java", "core", "features", "ChannelNotificationOnDonation.java"), []byte(ChannelNotificationOnDonation()), 0644)
			cnofFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "java", "core", "features", "ChannelNotificationOnFollow.java"), []byte(ChannelNotificationOnFollow()), 0644)
			cnosFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "java", "core", "features", "ChannelNotificationOnSubscription.java"), []byte(ChannelNotificationOnSubscription()), 0644)
			wcctcFile := os.WriteFile(filepath.Join(botName, "app", "src", "main", "java", "core", "features", "WriteChannelChatToConsole.java"), []byte(WriteChannelChatToConsole()), 0644)

			if startFile != nil {
				log.Fatal(startFile)
			}

			if cnodFile != nil {
				log.Fatal(cnodFile)
			}

			if cnofFile != nil {
				log.Fatal(cnofFile)
			}

			if cnosFile != nil {
				log.Fatal(cnosFile)
			}

			if wcctcFile != nil {
				log.Fatal(wcctcFile)
			}
		}

		if botlinFile != nil {
			log.Fatal(botlinFile)
		}

		if mainFile != nil {
			log.Fatal(mainFile)
		}

		if buildGradleFile != nil {
			log.Fatal(buildGradleFile)
		}

		if gradleWrapperPropsFile != nil {
			log.Fatal(gradleWrapperPropsFile)
		}

		if gradlewFile != nil {
			log.Fatal(gradlewFile)
		}

		if gradlewBatFile != nil {
			log.Fatal(gradlewBatFile)
		}

		if settingsFile != nil {
			log.Fatal(settingsFile)
		}

		if dockerFile != nil {
			log.Fatal(dockerFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		if resourcesFile != nil {
			log.Fatal(resourcesFile)
		}

		if gitattributesFile != nil {
			log.Fatal(gitattributesFile)
		}

		gradleWrapper := gradle + " wrapper"

		wrapperCmd := exec.Command("bash", "-c", gradleWrapper)

		if runtime.GOOS == "windows" {
			wrapperCmd = exec.Command("powershell.exe", gradleWrapper)
		}

		wrapperCmd.Dir = botName
		wrapperCmd.Stdin = os.Stdin
		wrapperCmd.Stdout = os.Stdout
		wrapperCmd.Stderr = os.Stderr
		err = wrapperCmd.Run()

		if err != nil {
			log.Printf("error: %v\n", err)
		}

		CheckProject(botName, platform)
	}
}
