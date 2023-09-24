// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtree/go/models"
	"github.com/fullstack-lang/gongtree/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Button:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, playground.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Buttons"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Node),
					"Buttons",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Node, *models.Button](
					nil,
					"Buttons",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Node:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("BackgroundColor", instanceWithInferedType.BackgroundColor, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("HasCheckboxButton", instanceWithInferedType.HasCheckboxButton, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsCheckboxDisabled", instanceWithInferedType.IsCheckboxDisabled, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsInEditMode", instanceWithInferedType.IsInEditMode, instanceWithInferedType, playground.formStage, formGroup, false)
		BasicFieldtoForm("IsNodeClickable", instanceWithInferedType.IsNodeClickable, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("Children", instanceWithInferedType, &instanceWithInferedType.Children, formGroup, playground)
		AssociationSliceToForm("Buttons", instanceWithInferedType, &instanceWithInferedType.Buttons, formGroup, playground)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Children"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Node),
					"Children",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Node, *models.Node](
					nil,
					"Children",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Tree"
			rf.Fieldname = "RootNodes"
			reverseFieldOwner := orm.GetReverseFieldOwner(playground.stageOfInterest, playground.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Tree),
					"RootNodes",
					instanceWithInferedType,
					formGroup,
					playground)
			} else {
				AssociationReverseFieldToForm[*models.Tree, *models.Node](
					nil,
					"RootNodes",
					instanceWithInferedType,
					formGroup,
					playground)
			}	
		}

	case *models.Tree:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup, false)
		AssociationSliceToForm("RootNodes", instanceWithInferedType, &instanceWithInferedType.RootNodes, formGroup, playground)

	default:
		_ = instanceWithInferedType
	}
}
