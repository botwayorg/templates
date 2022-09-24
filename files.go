package templates

import "fmt"

func DockerfileContent(botName, hostService, dockerFile, platform string) string {
	return Content(fmt.Sprintf("dockerfiles/%s/%s.dockerfile", hostService, dockerFile), "botway", botName, platform)
}

func Resources(platform, file string) string {
	return Content(platform+"/"+file, "resources", "", "")
}
