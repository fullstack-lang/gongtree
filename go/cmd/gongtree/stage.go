package main

import (
	"time"

	"github.com/fullstack-lang/gongtree/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_stage models.StageStruct
var ___dummy__Time_stage time.Time

// Injection point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_stage map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["stage"] = stageInjection
// }

// stageInjection will stage objects of database "stage"
func stageInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Button
	__Button__000000_arrow_circle_left := (&models.Button{Name: `arrow_circle_left`}).Stage(stage)
	__Button__000001_dataset := (&models.Button{Name: `dataset`}).Stage(stage)
	__Button__000002_dynamic_form := (&models.Button{Name: `dynamic_form`}).Stage(stage)
	__Button__000003_key := (&models.Button{Name: `key`}).Stage(stage)
	__Button__000004_logout := (&models.Button{Name: `logout`}).Stage(stage)
	__Button__000005_settings := (&models.Button{Name: `settings`}).Stage(stage)
	__Button__000006_test := (&models.Button{Name: `test`}).Stage(stage)

	// Declarations of staged instances of Node
	__Node__000000_root := (&models.Node{Name: `root`}).Stage(stage)
	__Node__000001_root2 := (&models.Node{Name: `root2`}).Stage(stage)
	__Node__000002_test := (&models.Node{Name: `test`}).Stage(stage)
	__Node__000003_test2 := (&models.Node{Name: `test2`}).Stage(stage)

	// Declarations of staged instances of Tree
	__Tree__000000_test := (&models.Tree{Name: `test`}).Stage(stage)

	// Setup of values

	// Button values setup
	__Button__000000_arrow_circle_left.Name = `arrow_circle_left`
	__Button__000000_arrow_circle_left.Icon = `arrow_circle_left`

	// Button values setup
	__Button__000001_dataset.Name = `dataset`
	__Button__000001_dataset.Icon = `dataset`

	// Button values setup
	__Button__000002_dynamic_form.Name = `dynamic_form`
	__Button__000002_dynamic_form.Icon = `dynamic_form`

	// Button values setup
	__Button__000003_key.Name = `key`
	__Button__000003_key.Icon = `key`

	// Button values setup
	__Button__000004_logout.Name = `logout`
	__Button__000004_logout.Icon = `logout`

	// Button values setup
	__Button__000005_settings.Name = `settings`
	__Button__000005_settings.Icon = `settings`

	// Button values setup
	__Button__000006_test.Name = `test`
	__Button__000006_test.Icon = `edit`

	// Node values setup
	__Node__000000_root.Name = `root`
	__Node__000000_root.IsExpanded = true
	__Node__000000_root.HasCheckboxButton = false
	__Node__000000_root.IsChecked = false
	__Node__000000_root.IsCheckboxDisabled = false
	__Node__000000_root.IsInEditMode = false

	// Node values setup
	__Node__000001_root2.Name = `root2`
	__Node__000001_root2.IsExpanded = false
	__Node__000001_root2.HasCheckboxButton = false
	__Node__000001_root2.IsChecked = false
	__Node__000001_root2.IsCheckboxDisabled = false
	__Node__000001_root2.IsInEditMode = false

	// Node values setup
	__Node__000002_test.Name = `test`
	__Node__000002_test.IsExpanded = false
	__Node__000002_test.HasCheckboxButton = false
	__Node__000002_test.IsChecked = false
	__Node__000002_test.IsCheckboxDisabled = false
	__Node__000002_test.IsInEditMode = false

	// Node values setup
	__Node__000003_test2.Name = `test2`
	__Node__000003_test2.IsExpanded = false
	__Node__000003_test2.HasCheckboxButton = false
	__Node__000003_test2.IsChecked = false
	__Node__000003_test2.IsCheckboxDisabled = false
	__Node__000003_test2.IsInEditMode = false

	// Tree values setup
	__Tree__000000_test.Name = `test`

	// Setup of pointers
	__Node__000000_root.Children = append(__Node__000000_root.Children, __Node__000002_test)
	__Node__000000_root.Children = append(__Node__000000_root.Children, __Node__000003_test2)
	__Node__000000_root.Buttons = append(__Node__000000_root.Buttons, __Button__000002_dynamic_form)
	__Node__000000_root.Buttons = append(__Node__000000_root.Buttons, __Button__000003_key)
	__Node__000002_test.Buttons = append(__Node__000002_test.Buttons, __Button__000006_test)
	__Node__000002_test.Buttons = append(__Node__000002_test.Buttons, __Button__000005_settings)
	__Node__000002_test.Buttons = append(__Node__000002_test.Buttons, __Button__000000_arrow_circle_left)
	__Node__000003_test2.Buttons = append(__Node__000003_test2.Buttons, __Button__000004_logout)
	__Node__000003_test2.Buttons = append(__Node__000003_test2.Buttons, __Button__000001_dataset)
	__Tree__000000_test.RootNodes = append(__Tree__000000_test.RootNodes, __Node__000000_root)
	__Tree__000000_test.RootNodes = append(__Tree__000000_test.RootNodes, __Node__000001_root2)
}


