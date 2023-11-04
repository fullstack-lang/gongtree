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
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Button:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Buttons"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Node),
					"Buttons",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Node, *models.Button](
					nil,
					"Buttons",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	case *models.Node:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("BackgroundColor", instanceWithInferedType.BackgroundColor, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("HasCheckboxButton", instanceWithInferedType.HasCheckboxButton, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("IsCheckboxDisabled", instanceWithInferedType.IsCheckboxDisabled, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("IsInEditMode", instanceWithInferedType.IsInEditMode, instanceWithInferedType, probe.formStage, formGroup, false)
		BasicFieldtoForm("IsNodeClickable", instanceWithInferedType.IsNodeClickable, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationSliceToForm("Children", instanceWithInferedType, &instanceWithInferedType.Children, formGroup, probe)
		AssociationSliceToForm("Buttons", instanceWithInferedType, &instanceWithInferedType.Buttons, formGroup, probe)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Node"
			rf.Fieldname = "Children"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Node),
					"Children",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Node, *models.Node](
					nil,
					"Children",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Tree"
			rf.Fieldname = "RootNodes"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Tree),
					"RootNodes",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Tree, *models.Node](
					nil,
					"RootNodes",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	case *models.Tree:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationSliceToForm("RootNodes", instanceWithInferedType, &instanceWithInferedType.RootNodes, formGroup, probe)

	default:
		_ = instanceWithInferedType
	}
}
