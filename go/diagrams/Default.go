package diagrams

import (
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"

	// injection point for ident package import declaration
	ref_models "github.com/fullstack-lang/gongtree/go/models"
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var ___dummy__Stage_Default models.StageStruct
var ___dummy__Time_Default time.Time

// Injection point for meta package dummy declaration
var ___dummy__ref_models_Default ref_models.StageStruct

// currently, DocLink renaming is not enabled in gopls
// the following map are devised to overcome this limitation
// those maps and the processing code will be eleminated when
// DocLink renaming will be enabled in gopls
// [Corresponding Issue](https://github.com/golang/go/issues/57559)
//
// When parsed, those maps will help with the renaming process
var map_DocLink_Identifier_Default map[string]any = map[string]any{
	// injection point for docLink to identifiers

	"ref_models.Button": &(ref_models.Button{}),

	"ref_models.Button.Icon": (ref_models.Button{}).Icon,

	"ref_models.Button.Name": (ref_models.Button{}).Name,

	"ref_models.Node": &(ref_models.Node{}),

	"ref_models.Node.Buttons": (ref_models.Node{}).Buttons,

	"ref_models.Node.Children": (ref_models.Node{}).Children,

	"ref_models.Node.HasCheckboxButton": (ref_models.Node{}).HasCheckboxButton,

	"ref_models.Node.IsCheckboxDisabled": (ref_models.Node{}).IsCheckboxDisabled,

	"ref_models.Node.IsChecked": (ref_models.Node{}).IsChecked,

	"ref_models.Node.IsExpanded": (ref_models.Node{}).IsExpanded,

	"ref_models.Node.IsInEditMode": (ref_models.Node{}).IsInEditMode,

	"ref_models.Node.Name": (ref_models.Node{}).Name,

	"ref_models.Tree": &(ref_models.Tree{}),

	"ref_models.Tree.Name": (ref_models.Tree{}).Name,

	"ref_models.Tree.RootNodes": (ref_models.Tree{}).RootNodes,
}

// init might be handy if one want to have the data embedded in the binary
// but it has to properly reference the Injection gateway in the main package
// func init() {
// 	_ = __Dummy_time_variable
// 	InjectionGateway["Default"] = DefaultInjection
// }

// DefaultInjection will stage objects of database "Default"
func DefaultInjection(stage *models.StageStruct) {

	// Declaration of instances to stage

	// Declarations of staged instances of Button

	// Declarations of staged instances of Classdiagram
	__Classdiagram__000000_Default := (&models.Classdiagram{Name: `Default`}).Stage(stage)

	// Declarations of staged instances of DiagramPackage

	// Declarations of staged instances of Field
	__Field__000000_HasCheckboxButton := (&models.Field{Name: `HasCheckboxButton`}).Stage(stage)
	__Field__000001_Icon := (&models.Field{Name: `Icon`}).Stage(stage)
	__Field__000002_IsCheckboxDisabled := (&models.Field{Name: `IsCheckboxDisabled`}).Stage(stage)
	__Field__000003_IsChecked := (&models.Field{Name: `IsChecked`}).Stage(stage)
	__Field__000004_IsExpanded := (&models.Field{Name: `IsExpanded`}).Stage(stage)
	__Field__000005_IsInEditMode := (&models.Field{Name: `IsInEditMode`}).Stage(stage)
	__Field__000006_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000007_Name := (&models.Field{Name: `Name`}).Stage(stage)
	__Field__000008_Name := (&models.Field{Name: `Name`}).Stage(stage)

	// Declarations of staged instances of GongEnumShape

	// Declarations of staged instances of GongEnumValueEntry

	// Declarations of staged instances of GongStructShape
	__GongStructShape__000000_Default_Button := (&models.GongStructShape{Name: `Default-Button`}).Stage(stage)
	__GongStructShape__000001_Default_Node := (&models.GongStructShape{Name: `Default-Node`}).Stage(stage)
	__GongStructShape__000002_Default_Tree := (&models.GongStructShape{Name: `Default-Tree`}).Stage(stage)

	// Declarations of staged instances of Link
	__Link__000000_Buttons := (&models.Link{Name: `Buttons`}).Stage(stage)
	__Link__000001_Children := (&models.Link{Name: `Children`}).Stage(stage)
	__Link__000002_RootNodes := (&models.Link{Name: `RootNodes`}).Stage(stage)

	// Declarations of staged instances of Node

	// Declarations of staged instances of NoteShape

	// Declarations of staged instances of NoteShapeLink

	// Declarations of staged instances of Position
	__Position__000000_Pos_Default_Button := (&models.Position{Name: `Pos-Default-Button`}).Stage(stage)
	__Position__000001_Pos_Default_Node := (&models.Position{Name: `Pos-Default-Node`}).Stage(stage)
	__Position__000002_Pos_Default_Tree := (&models.Position{Name: `Pos-Default-Tree`}).Stage(stage)

	// Declarations of staged instances of Tree

	// Declarations of staged instances of UmlState

	// Declarations of staged instances of Umlsc

	// Declarations of staged instances of Vertice
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Node and Default-Button`}).Stage(stage)
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Node and Default-Node`}).Stage(stage)
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node := (&models.Vertice{Name: `Verticle in class diagram Default in middle between Default-Tree and Default-Node`}).Stage(stage)

	// Setup of values

	// Classdiagram values setup
	__Classdiagram__000000_Default.Name = `Default`
	__Classdiagram__000000_Default.IsInDrawMode = true

	// Field values setup
	__Field__000000_HasCheckboxButton.Name = `HasCheckboxButton`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.HasCheckboxButton]
	__Field__000000_HasCheckboxButton.Identifier = `ref_models.Node.HasCheckboxButton`
	__Field__000000_HasCheckboxButton.FieldTypeAsString = ``
	__Field__000000_HasCheckboxButton.Structname = `Node`
	__Field__000000_HasCheckboxButton.Fieldtypename = `bool`

	// Field values setup
	__Field__000001_Icon.Name = `Icon`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Button.Icon]
	__Field__000001_Icon.Identifier = `ref_models.Button.Icon`
	__Field__000001_Icon.FieldTypeAsString = ``
	__Field__000001_Icon.Structname = `Button`
	__Field__000001_Icon.Fieldtypename = `string`

	// Field values setup
	__Field__000002_IsCheckboxDisabled.Name = `IsCheckboxDisabled`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.IsCheckboxDisabled]
	__Field__000002_IsCheckboxDisabled.Identifier = `ref_models.Node.IsCheckboxDisabled`
	__Field__000002_IsCheckboxDisabled.FieldTypeAsString = ``
	__Field__000002_IsCheckboxDisabled.Structname = `Node`
	__Field__000002_IsCheckboxDisabled.Fieldtypename = `bool`

	// Field values setup
	__Field__000003_IsChecked.Name = `IsChecked`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.IsChecked]
	__Field__000003_IsChecked.Identifier = `ref_models.Node.IsChecked`
	__Field__000003_IsChecked.FieldTypeAsString = ``
	__Field__000003_IsChecked.Structname = `Node`
	__Field__000003_IsChecked.Fieldtypename = `bool`

	// Field values setup
	__Field__000004_IsExpanded.Name = `IsExpanded`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.IsExpanded]
	__Field__000004_IsExpanded.Identifier = `ref_models.Node.IsExpanded`
	__Field__000004_IsExpanded.FieldTypeAsString = ``
	__Field__000004_IsExpanded.Structname = `Node`
	__Field__000004_IsExpanded.Fieldtypename = `bool`

	// Field values setup
	__Field__000005_IsInEditMode.Name = `IsInEditMode`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.IsInEditMode]
	__Field__000005_IsInEditMode.Identifier = `ref_models.Node.IsInEditMode`
	__Field__000005_IsInEditMode.FieldTypeAsString = ``
	__Field__000005_IsInEditMode.Structname = `Node`
	__Field__000005_IsInEditMode.Fieldtypename = `bool`

	// Field values setup
	__Field__000006_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Tree.Name]
	__Field__000006_Name.Identifier = `ref_models.Tree.Name`
	__Field__000006_Name.FieldTypeAsString = ``
	__Field__000006_Name.Structname = `Tree`
	__Field__000006_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000007_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.Name]
	__Field__000007_Name.Identifier = `ref_models.Node.Name`
	__Field__000007_Name.FieldTypeAsString = ``
	__Field__000007_Name.Structname = `Node`
	__Field__000007_Name.Fieldtypename = `string`

	// Field values setup
	__Field__000008_Name.Name = `Name`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Button.Name]
	__Field__000008_Name.Identifier = `ref_models.Button.Name`
	__Field__000008_Name.FieldTypeAsString = ``
	__Field__000008_Name.Structname = `Button`
	__Field__000008_Name.Fieldtypename = `string`

	// GongStructShape values setup
	__GongStructShape__000000_Default_Button.Name = `Default-Button`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Button]
	__GongStructShape__000000_Default_Button.Identifier = `ref_models.Button`
	__GongStructShape__000000_Default_Button.ShowNbInstances = false
	__GongStructShape__000000_Default_Button.NbInstances = 0
	__GongStructShape__000000_Default_Button.Width = 240.000000
	__GongStructShape__000000_Default_Button.Height = 93.000000
	__GongStructShape__000000_Default_Button.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000001_Default_Node.Name = `Default-Node`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__GongStructShape__000001_Default_Node.Identifier = `ref_models.Node`
	__GongStructShape__000001_Default_Node.ShowNbInstances = false
	__GongStructShape__000001_Default_Node.NbInstances = 0
	__GongStructShape__000001_Default_Node.Width = 240.000000
	__GongStructShape__000001_Default_Node.Height = 337.000000
	__GongStructShape__000001_Default_Node.IsSelected = false

	// GongStructShape values setup
	__GongStructShape__000002_Default_Tree.Name = `Default-Tree`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Tree]
	__GongStructShape__000002_Default_Tree.Identifier = `ref_models.Tree`
	__GongStructShape__000002_Default_Tree.ShowNbInstances = false
	__GongStructShape__000002_Default_Tree.NbInstances = 0
	__GongStructShape__000002_Default_Tree.Width = 240.000000
	__GongStructShape__000002_Default_Tree.Height = 78.000000
	__GongStructShape__000002_Default_Tree.IsSelected = false

	// Link values setup
	__Link__000000_Buttons.Name = `Buttons`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.Buttons]
	__Link__000000_Buttons.Identifier = `ref_models.Node.Buttons`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Button]
	__Link__000000_Buttons.Fieldtypename = `ref_models.Button`
	__Link__000000_Buttons.FieldOffsetX = -74.000000
	__Link__000000_Buttons.FieldOffsetY = -19.000000
	__Link__000000_Buttons.TargetMultiplicity = models.MANY
	__Link__000000_Buttons.TargetMultiplicityOffsetX = -36.000000
	__Link__000000_Buttons.TargetMultiplicityOffsetY = 30.000000
	__Link__000000_Buttons.SourceMultiplicity = models.ZERO_ONE
	__Link__000000_Buttons.SourceMultiplicityOffsetX = -42.000000
	__Link__000000_Buttons.SourceMultiplicityOffsetY = -16.000000
	__Link__000000_Buttons.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Buttons.StartRatio = 0.925816
	__Link__000000_Buttons.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000000_Buttons.EndRatio = 0.500000
	__Link__000000_Buttons.CornerOffsetRatio = -0.735764

	// Link values setup
	__Link__000001_Children.Name = `Children`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node.Children]
	__Link__000001_Children.Identifier = `ref_models.Node.Children`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000001_Children.Fieldtypename = `ref_models.Node`
	__Link__000001_Children.FieldOffsetX = -73.000000
	__Link__000001_Children.FieldOffsetY = 22.000000
	__Link__000001_Children.TargetMultiplicity = models.MANY
	__Link__000001_Children.TargetMultiplicityOffsetX = -24.000000
	__Link__000001_Children.TargetMultiplicityOffsetY = -7.000000
	__Link__000001_Children.SourceMultiplicity = models.ZERO_ONE
	__Link__000001_Children.SourceMultiplicityOffsetX = -33.000000
	__Link__000001_Children.SourceMultiplicityOffsetY = 24.000000
	__Link__000001_Children.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Children.StartRatio = 0.451039
	__Link__000001_Children.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000001_Children.EndRatio = 0.676558
	__Link__000001_Children.CornerOffsetRatio = -0.898264

	// Link values setup
	__Link__000002_RootNodes.Name = `RootNodes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Tree.RootNodes]
	__Link__000002_RootNodes.Identifier = `ref_models.Tree.RootNodes`

	// comment added to overcome the problem with the comment map association

	//gong:ident [ref_models.Node]
	__Link__000002_RootNodes.Fieldtypename = `ref_models.Node`
	__Link__000002_RootNodes.FieldOffsetX = -100.000000
	__Link__000002_RootNodes.FieldOffsetY = -19.000000
	__Link__000002_RootNodes.TargetMultiplicity = models.MANY
	__Link__000002_RootNodes.TargetMultiplicityOffsetX = -28.000000
	__Link__000002_RootNodes.TargetMultiplicityOffsetY = 24.000000
	__Link__000002_RootNodes.SourceMultiplicity = models.ZERO_ONE
	__Link__000002_RootNodes.SourceMultiplicityOffsetX = 30.000000
	__Link__000002_RootNodes.SourceMultiplicityOffsetY = 28.000000
	__Link__000002_RootNodes.StartOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_RootNodes.StartRatio = 0.500000
	__Link__000002_RootNodes.EndOrientation = models.ORIENTATION_HORIZONTAL
	__Link__000002_RootNodes.EndRatio = 0.130564
	__Link__000002_RootNodes.CornerOffsetRatio = 1.255903

	// Position values setup
	__Position__000000_Pos_Default_Button.X = 458.999969
	__Position__000000_Pos_Default_Button.Y = 445.000000
	__Position__000000_Pos_Default_Button.Name = `Pos-Default-Button`

	// Position values setup
	__Position__000001_Pos_Default_Node.X = 460.999969
	__Position__000001_Pos_Default_Node.Y = 76.000000
	__Position__000001_Pos_Default_Node.Name = `Pos-Default-Node`

	// Position values setup
	__Position__000002_Pos_Default_Tree.X = 15.000000
	__Position__000002_Pos_Default_Tree.Y = 84.000000
	__Position__000002_Pos_Default_Tree.Name = `Pos-Default-Tree`

	// Vertice values setup
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button.X = 424.000000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button.Y = 82.500000
	__Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button.Name = `Verticle in class diagram Default in middle between Default-Node and Default-Button`

	// Vertice values setup
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node.X = 454.000000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node.Y = 112.500000
	__Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node.Name = `Verticle in class diagram Default in middle between Default-Node and Default-Node`

	// Vertice values setup
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node.X = 443.500000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node.Y = 110.000000
	__Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node.Name = `Verticle in class diagram Default in middle between Default-Tree and Default-Node`

	// Setup of pointers
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000000_Default_Button)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000001_Default_Node)
	__Classdiagram__000000_Default.GongStructShapes = append(__Classdiagram__000000_Default.GongStructShapes, __GongStructShape__000002_Default_Tree)
	__GongStructShape__000000_Default_Button.Position = __Position__000000_Pos_Default_Button
	__GongStructShape__000000_Default_Button.Fields = append(__GongStructShape__000000_Default_Button.Fields, __Field__000008_Name)
	__GongStructShape__000000_Default_Button.Fields = append(__GongStructShape__000000_Default_Button.Fields, __Field__000001_Icon)
	__GongStructShape__000001_Default_Node.Position = __Position__000001_Pos_Default_Node
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000007_Name)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000004_IsExpanded)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000000_HasCheckboxButton)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000003_IsChecked)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000002_IsCheckboxDisabled)
	__GongStructShape__000001_Default_Node.Fields = append(__GongStructShape__000001_Default_Node.Fields, __Field__000005_IsInEditMode)
	__GongStructShape__000001_Default_Node.Links = append(__GongStructShape__000001_Default_Node.Links, __Link__000000_Buttons)
	__GongStructShape__000001_Default_Node.Links = append(__GongStructShape__000001_Default_Node.Links, __Link__000001_Children)
	__GongStructShape__000002_Default_Tree.Position = __Position__000002_Pos_Default_Tree
	__GongStructShape__000002_Default_Tree.Fields = append(__GongStructShape__000002_Default_Tree.Fields, __Field__000006_Name)
	__GongStructShape__000002_Default_Tree.Links = append(__GongStructShape__000002_Default_Tree.Links, __Link__000002_RootNodes)
	__Link__000000_Buttons.Middlevertice = __Vertice__000000_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Button
	__Link__000001_Children.Middlevertice = __Vertice__000001_Verticle_in_class_diagram_Default_in_middle_between_Default_Node_and_Default_Node
	__Link__000002_RootNodes.Middlevertice = __Vertice__000002_Verticle_in_class_diagram_Default_in_middle_between_Default_Tree_and_Default_Node
}
