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
	probe *Probe,
) (buttonFormCallback *ButtonFormCallback) {
	buttonFormCallback = new(ButtonFormCallback)
	buttonFormCallback.probe = probe
	buttonFormCallback.button = button

	buttonFormCallback.CreationMode = (button == nil)

	return
}

type ButtonFormCallback struct {
	button *models.Button

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (buttonFormCallback *ButtonFormCallback) OnSave() {

	log.Println("ButtonFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	buttonFormCallback.probe.formStage.Checkout()

	if buttonFormCallback.button == nil {
		buttonFormCallback.button = new(models.Button).Stage(buttonFormCallback.probe.stageOfInterest)
	}
	button_ := buttonFormCallback.button
	_ = button_

	// get the formGroup
	formGroup := buttonFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
				buttonFormCallback.probe.stageOfInterest,
				buttonFormCallback.probe.backRepoOfInterest,
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
				for _node := range *models.GetGongstructInstancesSet[models.Node](buttonFormCallback.probe.stageOfInterest) {

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

	buttonFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Button](
		buttonFormCallback.probe,
	)
	buttonFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if buttonFormCallback.CreationMode {
		buttonFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__ButtonFormCallback(
				nil,
				buttonFormCallback.probe,
			),
		}).Stage(buttonFormCallback.probe.formStage)
		button := new(models.Button)
		FillUpForm(button, newFormGroup, buttonFormCallback.probe)
		buttonFormCallback.probe.formStage.Commit()
	}

	fillUpTree(buttonFormCallback.probe)
}
func __gong__New__NodeFormCallback(
	node *models.Node,
	probe *Probe,
) (nodeFormCallback *NodeFormCallback) {
	nodeFormCallback = new(NodeFormCallback)
	nodeFormCallback.probe = probe
	nodeFormCallback.node = node

	nodeFormCallback.CreationMode = (node == nil)

	return
}

type NodeFormCallback struct {
	node *models.Node

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (nodeFormCallback *NodeFormCallback) OnSave() {

	log.Println("NodeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	nodeFormCallback.probe.formStage.Checkout()

	if nodeFormCallback.node == nil {
		nodeFormCallback.node = new(models.Node).Stage(nodeFormCallback.probe.stageOfInterest)
	}
	node_ := nodeFormCallback.node
	_ = node_

	// get the formGroup
	formGroup := nodeFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

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
		case "IsWithPreceedingIcon":
			FormDivBasicFieldToField(&(node_.IsWithPreceedingIcon), formDiv)
		case "PreceedingIcon":
			FormDivBasicFieldToField(&(node_.PreceedingIcon), formDiv)
		case "PreceedingSVGIcon":
			FormDivSelectFieldToField(&(node_.PreceedingSVGIcon), nodeFormCallback.probe.stageOfInterest, formDiv)
		case "Node:Children":
			// we need to retrieve the field owner before the change
			var pastNodeOwner *models.Node
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Children"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				nodeFormCallback.probe.stageOfInterest,
				nodeFormCallback.probe.backRepoOfInterest,
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
				for _node := range *models.GetGongstructInstancesSet[models.Node](nodeFormCallback.probe.stageOfInterest) {

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
				nodeFormCallback.probe.stageOfInterest,
				nodeFormCallback.probe.backRepoOfInterest,
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
				for _tree := range *models.GetGongstructInstancesSet[models.Tree](nodeFormCallback.probe.stageOfInterest) {

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

	nodeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Node](
		nodeFormCallback.probe,
	)
	nodeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if nodeFormCallback.CreationMode {
		nodeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__NodeFormCallback(
				nil,
				nodeFormCallback.probe,
			),
		}).Stage(nodeFormCallback.probe.formStage)
		node := new(models.Node)
		FillUpForm(node, newFormGroup, nodeFormCallback.probe)
		nodeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(nodeFormCallback.probe)
}
func __gong__New__SVGIconFormCallback(
	svgicon *models.SVGIcon,
	probe *Probe,
) (svgiconFormCallback *SVGIconFormCallback) {
	svgiconFormCallback = new(SVGIconFormCallback)
	svgiconFormCallback.probe = probe
	svgiconFormCallback.svgicon = svgicon

	svgiconFormCallback.CreationMode = (svgicon == nil)

	return
}

type SVGIconFormCallback struct {
	svgicon *models.SVGIcon

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (svgiconFormCallback *SVGIconFormCallback) OnSave() {

	log.Println("SVGIconFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	svgiconFormCallback.probe.formStage.Checkout()

	if svgiconFormCallback.svgicon == nil {
		svgiconFormCallback.svgicon = new(models.SVGIcon).Stage(svgiconFormCallback.probe.stageOfInterest)
	}
	svgicon_ := svgiconFormCallback.svgicon
	_ = svgicon_

	// get the formGroup
	formGroup := svgiconFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(svgicon_.Name), formDiv)
		case "SVG":
			FormDivBasicFieldToField(&(svgicon_.SVG), formDiv)
		}
	}

	svgiconFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.SVGIcon](
		svgiconFormCallback.probe,
	)
	svgiconFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if svgiconFormCallback.CreationMode {
		svgiconFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__SVGIconFormCallback(
				nil,
				svgiconFormCallback.probe,
			),
		}).Stage(svgiconFormCallback.probe.formStage)
		svgicon := new(models.SVGIcon)
		FillUpForm(svgicon, newFormGroup, svgiconFormCallback.probe)
		svgiconFormCallback.probe.formStage.Commit()
	}

	fillUpTree(svgiconFormCallback.probe)
}
func __gong__New__TreeFormCallback(
	tree *models.Tree,
	probe *Probe,
) (treeFormCallback *TreeFormCallback) {
	treeFormCallback = new(TreeFormCallback)
	treeFormCallback.probe = probe
	treeFormCallback.tree = tree

	treeFormCallback.CreationMode = (tree == nil)

	return
}

type TreeFormCallback struct {
	tree *models.Tree

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (treeFormCallback *TreeFormCallback) OnSave() {

	log.Println("TreeFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	treeFormCallback.probe.formStage.Checkout()

	if treeFormCallback.tree == nil {
		treeFormCallback.tree = new(models.Tree).Stage(treeFormCallback.probe.stageOfInterest)
	}
	tree_ := treeFormCallback.tree
	_ = tree_

	// get the formGroup
	formGroup := treeFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(tree_.Name), formDiv)
		}
	}

	treeFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Tree](
		treeFormCallback.probe,
	)
	treeFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if treeFormCallback.CreationMode {
		treeFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__TreeFormCallback(
				nil,
				treeFormCallback.probe,
			),
		}).Stage(treeFormCallback.probe.formStage)
		tree := new(models.Tree)
		FillUpForm(tree, newFormGroup, treeFormCallback.probe)
		treeFormCallback.probe.formStage.Commit()
	}

	fillUpTree(treeFormCallback.probe)
}
