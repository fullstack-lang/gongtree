// generated code - do not edit
package probe

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtree/go/models"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewButtonFormCallback(
	button *models.Button,
	playground *Playground,
) (buttonFormCallback *ButtonFormCallback) {
	buttonFormCallback = new(ButtonFormCallback)
	buttonFormCallback.playground = playground
	buttonFormCallback.button = button
	return
}

type ButtonFormCallback struct {
	button *models.Button

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
		}
	}

	buttonFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Button](
		buttonFormCallback.playground,
	)
	buttonFormCallback.playground.tableStage.Commit()
}
func NewNodeFormCallback(
	node *models.Node,
	playground *Playground,
) (nodeFormCallback *NodeFormCallback) {
	nodeFormCallback = new(NodeFormCallback)
	nodeFormCallback.playground = playground
	nodeFormCallback.node = node
	return
}

type NodeFormCallback struct {
	node *models.Node

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
		}
	}

	nodeFormCallback.playground.stageOfInterest.Commit()
	fillUpTable[models.Node](
		nodeFormCallback.playground,
	)
	nodeFormCallback.playground.tableStage.Commit()
}
func NewTreeFormCallback(
	tree *models.Tree,
	playground *Playground,
) (treeFormCallback *TreeFormCallback) {
	treeFormCallback = new(TreeFormCallback)
	treeFormCallback.playground = playground
	treeFormCallback.tree = tree
	return
}

type TreeFormCallback struct {
	tree *models.Tree

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
}
