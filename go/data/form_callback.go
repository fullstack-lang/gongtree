// generated code - do not edit
package data

import (
	"log"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"
	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gongtree/go/models"
	"github.com/fullstack-lang/gongtree/go/orm"
)

const __dummmy__time = time.Nanosecond

// insertion point
func NewButtonFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	button *models.Button,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (buttonFormCallback *ButtonFormCallback) {
	buttonFormCallback = new(ButtonFormCallback)
	buttonFormCallback.stageOfInterest = stageOfInterest
	buttonFormCallback.tableStage = tableStage
	buttonFormCallback.formStage = formStage
	buttonFormCallback.button = button
	buttonFormCallback.r = r
	buttonFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type ButtonFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	button            *models.Button
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (buttonFormCallback *ButtonFormCallback) OnSave() {

	log.Println("ButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	buttonFormCallback.formStage.Checkout()

	if buttonFormCallback.button == nil {
		buttonFormCallback.button = new(models.Button).Stage(buttonFormCallback.stageOfInterest)
	}
	button_ := buttonFormCallback.button
	_ = button_

	// get the formGroup
	formGroup := buttonFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(button_.Name), formDiv)
		case "Icon":
			FormDivBasicFieldToField(&(button_.Icon), formDiv)
		}
	}

	buttonFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Button](
		buttonFormCallback.stageOfInterest,
		buttonFormCallback.tableStage,
		buttonFormCallback.formStage,
		buttonFormCallback.r,
		buttonFormCallback.backRepoOfInterest,
	)
	buttonFormCallback.tableStage.Commit()
}
func NewNodeFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	node *models.Node,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (nodeFormCallback *NodeFormCallback) {
	nodeFormCallback = new(NodeFormCallback)
	nodeFormCallback.stageOfInterest = stageOfInterest
	nodeFormCallback.tableStage = tableStage
	nodeFormCallback.formStage = formStage
	nodeFormCallback.node = node
	nodeFormCallback.r = r
	nodeFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type NodeFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	node            *models.Node
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (nodeFormCallback *NodeFormCallback) OnSave() {

	log.Println("NodeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	nodeFormCallback.formStage.Checkout()

	if nodeFormCallback.node == nil {
		nodeFormCallback.node = new(models.Node).Stage(nodeFormCallback.stageOfInterest)
	}
	node_ := nodeFormCallback.node
	_ = node_

	// get the formGroup
	formGroup := nodeFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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

	nodeFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Node](
		nodeFormCallback.stageOfInterest,
		nodeFormCallback.tableStage,
		nodeFormCallback.formStage,
		nodeFormCallback.r,
		nodeFormCallback.backRepoOfInterest,
	)
	nodeFormCallback.tableStage.Commit()
}
func NewTreeFormCallback(
	stageOfInterest *models.StageStruct,
	tableStage *table.StageStruct,
	formStage *table.StageStruct,
	tree *models.Tree,
	r *gin.Engine,
	backRepoOfInterest *orm.BackRepoStruct,
) (treeFormCallback *TreeFormCallback) {
	treeFormCallback = new(TreeFormCallback)
	treeFormCallback.stageOfInterest = stageOfInterest
	treeFormCallback.tableStage = tableStage
	treeFormCallback.formStage = formStage
	treeFormCallback.tree = tree
	treeFormCallback.r = r
	treeFormCallback.backRepoOfInterest = backRepoOfInterest
	return
}

type TreeFormCallback struct {
	stageOfInterest    *models.StageStruct
	tableStage         *table.StageStruct
	formStage          *table.StageStruct
	tree            *models.Tree
	r                  *gin.Engine
	backRepoOfInterest *orm.BackRepoStruct
}

func (treeFormCallback *TreeFormCallback) OnSave() {

	log.Println("TreeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	treeFormCallback.formStage.Checkout()

	if treeFormCallback.tree == nil {
		treeFormCallback.tree = new(models.Tree).Stage(treeFormCallback.stageOfInterest)
	}
	tree_ := treeFormCallback.tree
	_ = tree_

	// get the formGroup
	formGroup := treeFormCallback.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tree_.Name), formDiv)
		}
	}

	treeFormCallback.stageOfInterest.Commit()
	fillUpTable[models.Tree](
		treeFormCallback.stageOfInterest,
		treeFormCallback.tableStage,
		treeFormCallback.formStage,
		treeFormCallback.r,
		treeFormCallback.backRepoOfInterest,
	)
	treeFormCallback.tableStage.Commit()
}
