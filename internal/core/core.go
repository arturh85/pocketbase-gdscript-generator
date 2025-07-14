package core

import (
	"fmt"
	"os"
	"strings"

	"github.com/arturh85/pocketbase-gdscript-generator/internal/cmd"
	"github.com/arturh85/pocketbase-gdscript-generator/internal/interpreter"
	"github.com/arturh85/pocketbase-gdscript-generator/internal/pocketbase_api"
	"github.com/rs/zerolog/log"
)

func ProcessCollections(selectedCollections []*pocketbase_api.Collection, allCollections []pocketbase_api.Collection, generatorFlags *cmd.GeneratorFlags) {
	interpretedCollections := interpreter.InterpretCollections(selectedCollections, allCollections)

	output := make([]string, 1+len(interpretedCollections))

	for i, collection := range interpretedCollections {
		output[i] = collection.GetGdScriptInterface(generatorFlags)
	}

	collectionDefinitions := make([]string, len(interpretedCollections))
	for i, collection := range interpretedCollections {
		collectionDefinitions[i] = collection.GetGdScriptCollectionEntry(generatorFlags)
	}

	output[len(interpretedCollections)] = fmt.Sprintf("class Collections:\n%s", strings.Join(collectionDefinitions, "\n"))

	joinedData := strings.Join(output, "\n\n")

	if generatorFlags.Output == "" {
		fmt.Println(joinedData)
	} else {
		err := os.WriteFile(generatorFlags.Output, []byte(joinedData), 0644)
		log.Info().Msgf("Saved generated interfaces to %s", generatorFlags.Output)
		if err != nil {
			log.Fatal().Err(err).Msg("Could not output contents")
		}
	}
}
