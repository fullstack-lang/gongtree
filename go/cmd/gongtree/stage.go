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
	__Button__000000_Test_1_2_add := (&models.Button{Name: `Test 1.2 add`}).Stage(stage)
	__Button__000001_arrow_circle_left := (&models.Button{Name: `arrow_circle_left`}).Stage(stage)
	__Button__000002_dataset := (&models.Button{Name: `dataset`}).Stage(stage)
	__Button__000003_dynamic_form := (&models.Button{Name: `dynamic_form`}).Stage(stage)
	__Button__000004_key := (&models.Button{Name: `key`}).Stage(stage)
	__Button__000005_logout := (&models.Button{Name: `logout`}).Stage(stage)
	__Button__000006_settings := (&models.Button{Name: `settings`}).Stage(stage)
	__Button__000007_test := (&models.Button{Name: `test`}).Stage(stage)

	// Declarations of staged instances of Node
	__Node__000000_ := (&models.Node{Name: ``}).Stage(stage)
	__Node__000001_root1 := (&models.Node{Name: `root1`}).Stage(stage)
	__Node__000002_root2 := (&models.Node{Name: `root2`}).Stage(stage)
	__Node__000003_root3 := (&models.Node{Name: `root3`}).Stage(stage)
	__Node__000004_test := (&models.Node{Name: `test`}).Stage(stage)
	__Node__000005_test_1_1 := (&models.Node{Name: `test 1.1`}).Stage(stage)
	__Node__000006_test_1_2_clickable_ := (&models.Node{Name: `test 1.2 (clickable)`}).Stage(stage)
	__Node__000007_test_1_3 := (&models.Node{Name: `test 1.3`}).Stage(stage)
	__Node__000008_test_1_3_1 := (&models.Node{Name: `test 1.3.1`}).Stage(stage)
	__Node__000009_test_1_4 := (&models.Node{Name: `test 1.4`}).Stage(stage)
	__Node__000010_test2_1 := (&models.Node{Name: `test2.1`}).Stage(stage)
	__Node__000011_test3 := (&models.Node{Name: `test3`}).Stage(stage)

	// Declarations of staged instances of Tree
	__Tree__000000_test := (&models.Tree{Name: `test`}).Stage(stage)

	// Setup of values

	// Button values setup
	__Button__000000_Test_1_2_add.Name = `Test 1.2 add`
	__Button__000000_Test_1_2_add.Icon = `add`

	// Button values setup
	__Button__000001_arrow_circle_left.Name = `arrow_circle_left`
	__Button__000001_arrow_circle_left.Icon = `arrow_circle_left`

	// Button values setup
	__Button__000002_dataset.Name = `dataset`
	__Button__000002_dataset.Icon = `dataset`

	// Button values setup
	__Button__000003_dynamic_form.Name = `dynamic_form`
	__Button__000003_dynamic_form.Icon = `dynamic_form`

	// Button values setup
	__Button__000004_key.Name = `key`
	__Button__000004_key.Icon = `key`

	// Button values setup
	__Button__000005_logout.Name = `logout`
	__Button__000005_logout.Icon = `logout`

	// Button values setup
	__Button__000006_settings.Name = `settings`
	__Button__000006_settings.Icon = `settings`

	// Button values setup
	__Button__000007_test.Name = `test`
	__Button__000007_test.Icon = `edit`

	// Node values setup
	__Node__000000_.Name = ``
	__Node__000000_.IsExpanded = false
	__Node__000000_.HasCheckboxButton = false
	__Node__000000_.IsChecked = false
	__Node__000000_.IsCheckboxDisabled = false
	__Node__000000_.IsInEditMode = false
	__Node__000000_.IsNodeClickable = false

	// Node values setup
	__Node__000001_root1.Name = `root1`
	__Node__000001_root1.IsExpanded = true
	__Node__000001_root1.HasCheckboxButton = false
	__Node__000001_root1.IsChecked = false
	__Node__000001_root1.IsCheckboxDisabled = false
	__Node__000001_root1.IsInEditMode = false
	__Node__000001_root1.IsNodeClickable = false

	// Node values setup
	__Node__000002_root2.Name = `root2`
	__Node__000002_root2.IsExpanded = true
	__Node__000002_root2.HasCheckboxButton = false
	__Node__000002_root2.IsChecked = false
	__Node__000002_root2.IsCheckboxDisabled = false
	__Node__000002_root2.IsInEditMode = false
	__Node__000002_root2.IsNodeClickable = false

	// Node values setup
	__Node__000003_root3.Name = `root3`
	__Node__000003_root3.IsExpanded = false
	__Node__000003_root3.HasCheckboxButton = false
	__Node__000003_root3.IsChecked = false
	__Node__000003_root3.IsCheckboxDisabled = false
	__Node__000003_root3.IsInEditMode = false
	__Node__000003_root3.IsNodeClickable = false

	// Node values setup
	__Node__000004_test.Name = `test`
	__Node__000004_test.IsExpanded = false
	__Node__000004_test.HasCheckboxButton = true
	__Node__000004_test.IsChecked = true
	__Node__000004_test.IsCheckboxDisabled = false
	__Node__000004_test.IsInEditMode = true
	__Node__000004_test.IsNodeClickable = false

	// Node values setup
	__Node__000005_test_1_1.Name = `test 1.1`
	__Node__000005_test_1_1.IsExpanded = false
	__Node__000005_test_1_1.HasCheckboxButton = true
	__Node__000005_test_1_1.IsChecked = true
	__Node__000005_test_1_1.IsCheckboxDisabled = true
	__Node__000005_test_1_1.IsInEditMode = false
	__Node__000005_test_1_1.IsNodeClickable = false

	// Node values setup
	__Node__000006_test_1_2_clickable_.Name = `test 1.2 (clickable)`
	__Node__000006_test_1_2_clickable_.IsExpanded = false
	__Node__000006_test_1_2_clickable_.HasCheckboxButton = false
	__Node__000006_test_1_2_clickable_.IsChecked = false
	__Node__000006_test_1_2_clickable_.IsCheckboxDisabled = false
	__Node__000006_test_1_2_clickable_.IsInEditMode = false
	__Node__000006_test_1_2_clickable_.IsNodeClickable = true

	// Node values setup
	__Node__000007_test_1_3.Name = `test 1.3`
	__Node__000007_test_1_3.IsExpanded = true
	__Node__000007_test_1_3.HasCheckboxButton = false
	__Node__000007_test_1_3.IsChecked = false
	__Node__000007_test_1_3.IsCheckboxDisabled = false
	__Node__000007_test_1_3.IsInEditMode = false
	__Node__000007_test_1_3.IsNodeClickable = false

	// Node values setup
	__Node__000008_test_1_3_1.Name = `test 1.3.1`
	__Node__000008_test_1_3_1.IsExpanded = false
	__Node__000008_test_1_3_1.HasCheckboxButton = false
	__Node__000008_test_1_3_1.IsChecked = false
	__Node__000008_test_1_3_1.IsCheckboxDisabled = false
	__Node__000008_test_1_3_1.IsInEditMode = false
	__Node__000008_test_1_3_1.IsNodeClickable = false

	// Node values setup
	__Node__000009_test_1_4.Name = `test 1.4`
	__Node__000009_test_1_4.IsExpanded = false
	__Node__000009_test_1_4.HasCheckboxButton = false
	__Node__000009_test_1_4.IsChecked = false
	__Node__000009_test_1_4.IsCheckboxDisabled = false
	__Node__000009_test_1_4.IsInEditMode = false
	__Node__000009_test_1_4.IsNodeClickable = false

	// Node values setup
	__Node__000010_test2_1.Name = `test2.1`
	__Node__000010_test2_1.IsExpanded = false
	__Node__000010_test2_1.HasCheckboxButton = true
	__Node__000010_test2_1.IsChecked = true
	__Node__000010_test2_1.IsCheckboxDisabled = false
	__Node__000010_test2_1.IsInEditMode = false
	__Node__000010_test2_1.IsNodeClickable = false

	// Node values setup
	__Node__000011_test3.Name = `test3`
	__Node__000011_test3.IsExpanded = false
	__Node__000011_test3.HasCheckboxButton = false
	__Node__000011_test3.IsChecked = false
	__Node__000011_test3.IsCheckboxDisabled = false
	__Node__000011_test3.IsInEditMode = false
	__Node__000011_test3.IsNodeClickable = false

	// Tree values setup
	__Tree__000000_test.Name = `test`

	// Setup of pointers
	__Node__000001_root1.Children = append(__Node__000001_root1.Children, __Node__000007_test_1_3)
	__Node__000001_root1.Children = append(__Node__000001_root1.Children, __Node__000004_test)
	__Node__000001_root1.Children = append(__Node__000001_root1.Children, __Node__000006_test_1_2_clickable_)
	__Node__000001_root1.Children = append(__Node__000001_root1.Children, __Node__000009_test_1_4)
	__Node__000001_root1.Children = append(__Node__000001_root1.Children, __Node__000005_test_1_1)
	__Node__000001_root1.Buttons = append(__Node__000001_root1.Buttons, __Button__000003_dynamic_form)
	__Node__000001_root1.Buttons = append(__Node__000001_root1.Buttons, __Button__000004_key)
	__Node__000002_root2.Children = append(__Node__000002_root2.Children, __Node__000010_test2_1)
	__Node__000003_root3.Children = append(__Node__000003_root3.Children, __Node__000011_test3)
	__Node__000004_test.Buttons = append(__Node__000004_test.Buttons, __Button__000007_test)
	__Node__000004_test.Buttons = append(__Node__000004_test.Buttons, __Button__000006_settings)
	__Node__000004_test.Buttons = append(__Node__000004_test.Buttons, __Button__000001_arrow_circle_left)
	__Node__000005_test_1_1.Buttons = append(__Node__000005_test_1_1.Buttons, __Button__000005_logout)
	__Node__000005_test_1_1.Buttons = append(__Node__000005_test_1_1.Buttons, __Button__000002_dataset)
	__Node__000006_test_1_2_clickable_.Buttons = append(__Node__000006_test_1_2_clickable_.Buttons, __Button__000000_Test_1_2_add)
	__Node__000007_test_1_3.Children = append(__Node__000007_test_1_3.Children, __Node__000008_test_1_3_1)
	__Tree__000000_test.RootNodes = append(__Tree__000000_test.RootNodes, __Node__000001_root1)
	__Tree__000000_test.RootNodes = append(__Tree__000000_test.RootNodes, __Node__000002_root2)
	__Tree__000000_test.RootNodes = append(__Tree__000000_test.RootNodes, __Node__000003_root3)
}


