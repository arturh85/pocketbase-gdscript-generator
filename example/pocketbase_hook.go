package main

import (
	"github.com/arturh85/pocketbase-gdscript-generator/pkg/pocketbase-gdscript-generator"
	"github.com/pocketbase/pocketbase"
	"github.com/rs/zerolog/log"
)

func main() {
	app := pocketbase.New()

	pocketbase_gdscript_generator.RegisterHook(app, &pocketbase_gdscript_generator.GeneratorOptions{
		Output: "test.ts",
	})

	if err := app.Start(); err != nil {
		log.Fatal().Err(err)
	}
}
