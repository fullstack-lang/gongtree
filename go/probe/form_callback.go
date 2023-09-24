// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtree/go/models"
	"github.com/fullstack-lang/gongtree/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__ButtonFormCallback(
	button *models.Button,
	playground *Playground,
) (buttonFormCallback *ButtonFormCallback) {
	buttonFormCallback = new(ButtonFormCallback)
	buttonFormCallback.playground = playground
	buttonFormCallback.button = button

	buttonFormCallback.CreationMode = (button == nil)

	return
}

type ButtonFormCallback struct {
	button *models.Button

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (buttonFormCallback *ButtonFormCallback) OnSave() {

	log.Println("ButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	buttonFormCallback.playground.formStage.Checkout()

	if buttonFormCallback.button == nil {
		buttonFormCallback.button = new(models.Button).Stage(buttonFormCallback.playground.stageOfInterest)
	}
	button_ := buttonFormCallback.button
	_ = button_

	// get the formGroup
	formGroup := buttonFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(button_.Name), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(button_.Icon), formDiv)
		case "Node:Buttons":
			// we need to retrieve the field owner before the change
			var pastNodeOwner *models.Node
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Buttons"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				buttonFormCallback.playground.stageOfInterest,
				buttonFormCallback.playground.backRepoOfInterest,
				button_,
				&rf)

			if reverseFieldOwner != nil {
				pastNodeOwner = reverseFieldOwner.(*models.Node)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastNodeOwner != nil {
					idx := slices.Index(pastNodeOwner.Buttons, button_)
					pastNodeOwner.Buttons = slices.Delete(pastNodeOwner.Buttons, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _node := range *models.GetGongstructInstancesSet[models.Node](buttonFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _node.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newNodeOwner := _node // we have a match
						if pastNodeOwner != nil {
							if newNodeOwner != pastNodeOwner {
								idx := slices.Index(pastNodeOwner.Buttons, button_)
								pastNodeOwner.Buttons = slices.Delete(pastNodeOwner.Buttons, idx, idx+1)
								newNodeOwner.Buttons = append(newNodeOwner.Buttons, button_)
							}
						} else {
							newNodeOwner.Buttons = append(newNodeOwner.Buttons, button_)
						}
					}
				}
			}
		}
	}

	buttonFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Button](
		buttonFormCallback.playground,
	)
	buttonFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if buttonFormCallback.CreationMode {
		buttonFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ButtonFormCallback(
				nil,
				buttonFormCallback.playground,
			),
		}).Stage(buttonFormCallback.playground.formStage)
		button := new(models.Button)
		FillUpForm(button, newFormGroup, buttonFormCallback.playground)
		buttonFormCallback.playground.formStage.Commit()
	}

	fillUpTree(buttonFormCallback.playground)
}
func __gong__New__NodeFormCallback(
	node *models.Node,
	playground *Playground,
) (nodeFormCallback *NodeFormCallback) {
	nodeFormCallback = new(NodeFormCallback)
	nodeFormCallback.playground = playground
	nodeFormCallback.node = node

	nodeFormCallback.CreationMode = (node == nil)

	return
}

type NodeFormCallback struct {
	node *models.Node

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (nodeFormCallback *NodeFormCallback) OnSave() {

	log.Println("NodeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	nodeFormCallback.playground.formStage.Checkout()

	if nodeFormCallback.node == nil {
		nodeFormCallback.node = new(models.Node).Stage(nodeFormCallback.playground.stageOfInterest)
	}
	node_ := nodeFormCallback.node
	_ = node_

	// get the formGroup
	formGroup := nodeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(node_.Name), formDiv)
		case "BackgroundColor":
			FormDivBasicFieldToField(&(node_.BackgroundColor), formDiv)
		case "IsExpanded":
			FormDivBasicFieldToField(&(node_.IsExpanded), formDiv)
		case "HasCheckboxButton":
			FormDivBasicFieldToField(&(node_.HasCheckboxButton), formDiv)
		case "IsChecked":
			FormDivBasicFieldToField(&(node_.IsChecked), formDiv)
		case "IsCheckboxDisabled":
			FormDivBasicFieldToField(&(node_.IsCheckboxDisabled), formDiv)
		case "IsInEditMode":
			FormDivBasicFieldToField(&(node_.IsInEditMode), formDiv)
		case "IsNodeClickable":
			FormDivBasicFieldToField(&(node_.IsNodeClickable), formDiv)
		case "Node:Children":
			// we need to retrieve the field owner before the change
			var pastNodeOwner *models.Node
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Children"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				nodeFormCallback.playground.stageOfInterest,
				nodeFormCallback.playground.backRepoOfInterest,
				node_,
				&rf)

			if reverseFieldOwner != nil {
				pastNodeOwner = reverseFieldOwner.(*models.Node)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastNodeOwner != nil {
					idx := slices.Index(pastNodeOwner.Children, node_)
					pastNodeOwner.Children = slices.Delete(pastNodeOwner.Children, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _node := range *models.GetGongstructInstancesSet[models.Node](nodeFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _node.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newNodeOwner := _node // we have a match
						if pastNodeOwner != nil {
							if newNodeOwner != pastNodeOwner {
								idx := slices.Index(pastNodeOwner.Children, node_)
								pastNodeOwner.Children = slices.Delete(pastNodeOwner.Children, idx, idx+1)
								newNodeOwner.Children = append(newNodeOwner.Children, node_)
							}
						} else {
							newNodeOwner.Children = append(newNodeOwner.Children, node_)
						}
					}
				}
			}
		case "Tree:RootNodes":
			// we need to retrieve the field owner before the change
			var pastTreeOwner *models.Tree
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Tree"
			rf.Fieldname = "RootNodes"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				nodeFormCallback.playground.stageOfInterest,
				nodeFormCallback.playground.backRepoOfInterest,
				node_,
				&rf)

			if reverseFieldOwner != nil {
				pastTreeOwner = reverseFieldOwner.(*models.Tree)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastTreeOwner != nil {
					idx := slices.Index(pastTreeOwner.RootNodes, node_)
					pastTreeOwner.RootNodes = slices.Delete(pastTreeOwner.RootNodes, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _tree := range *models.GetGongstructInstancesSet[models.Tree](nodeFormCallback.playground.stageOfInterest) {

					// the match is base on the name
					if _tree.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newTreeOwner := _tree // we have a match
						if pastTreeOwner != nil {
							if newTreeOwner != pastTreeOwner {
								idx := slices.Index(pastTreeOwner.RootNodes, node_)
								pastTreeOwner.RootNodes = slices.Delete(pastTreeOwner.RootNodes, idx, idx+1)
								newTreeOwner.RootNodes = append(newTreeOwner.RootNodes, node_)
							}
						} else {
							newTreeOwner.RootNodes = append(newTreeOwner.RootNodes, node_)
						}
					}
				}
			}
		}
	}

	nodeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Node](
		nodeFormCallback.playground,
	)
	nodeFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if nodeFormCallback.CreationMode {
		nodeFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__NodeFormCallback(
				nil,
				nodeFormCallback.playground,
			),
		}).Stage(nodeFormCallback.playground.formStage)
		node := new(models.Node)
		FillUpForm(node, newFormGroup, nodeFormCallback.playground)
		nodeFormCallback.playground.formStage.Commit()
	}

	fillUpTree(nodeFormCallback.playground)
}
func __gong__New__TreeFormCallback(
	tree *models.Tree,
	playground *Playground,
) (treeFormCallback *TreeFormCallback) {
	treeFormCallback = new(TreeFormCallback)
	treeFormCallback.playground = playground
	treeFormCallback.tree = tree

	treeFormCallback.CreationMode = (tree == nil)

	return
}

type TreeFormCallback struct {
	tree *models.Tree

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	playground *Playground
}

func (treeFormCallback *TreeFormCallback) OnSave() {

	log.Println("TreeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	treeFormCallback.playground.formStage.Checkout()

	if treeFormCallback.tree == nil {
		treeFormCallback.tree = new(models.Tree).Stage(treeFormCallback.playground.stageOfInterest)
	}
	tree_ := treeFormCallback.tree
	_ = tree_

	// get the formGroup
	formGroup := treeFormCallback.playground.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tree_.Name), formDiv)
		}
	}

	treeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Tree](
		treeFormCallback.playground,
	)
	treeFormCallback.playground.tableStage.Commit()

	// display a new form by reset the form stage
	if treeFormCallback.CreationMode {
		treeFormCallback.playground.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TreeFormCallback(
				nil,
				treeFormCallback.playground,
			),
		}).Stage(treeFormCallback.playground.formStage)
		tree := new(models.Tree)
		FillUpForm(tree, newFormGroup, treeFormCallback.playground)
		treeFormCallback.playground.formStage.Commit()
	}

	fillUpTree(treeFormCallback.playground)
}
