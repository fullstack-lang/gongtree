// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtree/go/models"
)

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	playground *Playground,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Button:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, playground.formStage, formGroup)

	case *models.Node:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("BackgroundColor", instanceWithInferedType.BackgroundColor, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsExpanded", instanceWithInferedType.IsExpanded, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("HasCheckboxButton", instanceWithInferedType.HasCheckboxButton, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsCheckboxDisabled", instanceWithInferedType.IsCheckboxDisabled, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsInEditMode", instanceWithInferedType.IsInEditMode, instanceWithInferedType, playground.formStage, formGroup)
		BasicFieldtoForm("IsNodeClickable", instanceWithInferedType.IsNodeClickable, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("Children", instanceWithInferedType, &instanceWithInferedType.Children, formGroup, playground)
		AssociationSliceToForm("Buttons", instanceWithInferedType, &instanceWithInferedType.Buttons, formGroup, playground)

	case *models.Tree:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, playground.formStage, formGroup)
		AssociationSliceToForm("RootNodes", instanceWithInferedType, &instanceWithInferedType.RootNodes, formGroup, playground)

	default:
		_ = instanceWithInferedType
	}
}
