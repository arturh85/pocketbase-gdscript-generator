package pocketbase_gdscript_generator

import (
	"github.com/arturh85/pocketbase-gdscript-generator/internal/cmd"
	"github.com/arturh85/pocketbase-gdscript-generator/internal/core"
	"github.com/arturh85/pocketbase-gdscript-generator/internal/forms"
	"github.com/arturh85/pocketbase-gdscript-generator/internal/pocketbase_api"
	"github.com/arturh85/pocketbase-gdscript-generator/internal/pocketbase_core"
	"github.com/pocketbase/pocketbase"
)

func processFileGeneration(app *pocketbase.PocketBase, generatorFlags *cmd.GeneratorFlags) error {
	collections, err := pocketbase_core.GetCollections(app)
	if err != nil {
		return err
	}

	var selectedCollections []*pocketbase_api.Collection

	selectedCollections = forms.GetSelectedCollections(generatorFlags, collections.Items)

	core.ProcessCollections(selectedCollections, collections.Items, generatorFlags)

	return nil
}
