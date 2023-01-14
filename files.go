package templates

import "fmt"

func DockerfileContent(botName, dockerFile, platform string) string {
	return Content(fmt.Sprintf("dockerfiles/%s", dockerFile), "botway", botName, platform)
}

func Resources(platform, file string) string {
	return Content(platform+"/"+file, "resources", "", "")
}
