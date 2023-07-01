package templates

import (
	"log"
	"os"
	"path/filepath"
)

func createDirs(botName, lang, platform string) {
	if err := os.Mkdir(filepath.Join(botName, "gradle"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "gradle", "wrapper"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "app"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "app", "src"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "app", "src", "main"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir(filepath.Join(botName, "app", "src", "main", lang), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if (lang == "kotlin") {
		if err := os.Mkdir(filepath.Join(botName, "app", "src", "main", lang, "botway"), os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}

	if err := os.Mkdir(filepath.Join(botName, "app", "src", "main", lang, "core"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if platform == "twitch" {
		if err := os.Mkdir(filepath.Join(botName, "app", "src", "main", lang, "core", "features"), os.ModePerm); err != nil {
			log.Fatal(err)
		}
	}
}

func BotlinContent() string {
	return Content("packages/botlin/main.kt", "botway", "", "")
}

func MainJavaContent(platform string) string {
	return Content("app/src/main/java/core/Bot.java", platform+"-java", "", "")
}

func BotHandlerContent() string {
	return Content("app/src/main/java/core/BotHandler.java", "telegram-java", "", "")
}

func TGBotContent() string {
	return Content("app/src/main/java/core/TGBot.java", "telegram-java", "", "")
}

func BuildGradleContent(platform string) string {
	return Content("app/build.gradle", platform+"-java", "", "")
}

func GradleWrapperPropsContent() string {
	return Content("gradle/wrapper/gradle-wrapper.properties", "telegram-java", "", "")
}

func DotGitattributesContent() string {
	return Content(".gitattributes", "telegram-java", "", "")
}

func GradlewContent() string {
	return Content("gradlew", "telegram-java", "", "")
}

func GradlewBatContent() string {
	return Content("gradlew.bat", "telegram-java", "", "")
}

func SettingsGradle() string {
	return Content("settings.gradle", "telegram-java", "", "")
}

func StartJavaContent() string {
	return Content("app/src/main/java/core/Start.java", "twitch-java", "", "")
}

func ChannelNotificationOnDonation() string {
	return Content("app/src/main/java/core/features/ChannelNotificationOnDonation.java", "twitch-java", "", "")
}

func ChannelNotificationOnFollow() string {
	return Content("app/src/main/java/core/features/ChannelNotificationOnFollow.java", "twitch-java", "", "")
}

func ChannelNotificationOnSubscription() string {
	return Content("app/src/main/java/core/features/ChannelNotificationOnSubscription.java", "twitch-java", "", "")
}

func WriteChannelChatToConsole() string {
	return Content("app/src/main/java/core/features/WriteChannelChatToConsole.java", "twitch-java", "", "")
}
