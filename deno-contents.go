package templates

func DenoMainTsContent() string {
	return Content("main.ts", "discord-deno", "", "")
}

func DepsTsContent() string {
	return Content("deps.ts", "discord-deno", "", "")
}

func CommandsModTsContent() string {
	return Content("src/commands/mod.ts", "discord-deno", "", "")
}

func CommandsPingTsContent() string {
	return Content("src/commands/ping.ts", "discord-deno", "", "")
}

func EventsGuildCreateTsContent() string {
	return Content("src/events/guildCreate.ts", "discord-deno", "", "")
}

func EventsInteractionCreateTsContent() string {
	return Content("src/events/interactionCreate.ts", "discord-deno", "", "")
}

func EventsModTsContent() string {
	return Content("src/events/mod.ts", "discord-deno", "", "")
}

func EventsReadyTsContent() string {
	return Content("src/events/ready.ts", "discord-deno", "", "")
}

func UtilsHelpersTsContent() string {
	return Content("src/utils/helpers.ts", "discord-deno", "", "")
}

func UtilsLoggerTsContent() string {
	return Content("src/utils/logger.ts", "discord-deno", "", "")
}
