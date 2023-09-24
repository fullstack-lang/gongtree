// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtree/go/models"
)

func FillUpFormFromGongstruct[T models.Gongstruct](instance *T, playground *Playground) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	switch instancesTyped := any(instance).(type) {
	// insertion point
	case *models.Button:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Button Form",
			OnSave: __gong__New__ButtonFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Node:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Node Form",
			OnSave: __gong__New__NodeFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	case *models.Tree:
		formGroup := (&gongtable.FormGroup{
			Name:  gongtable.FormGroupDefaultName.ToString(),
			Label: "Update Tree Form",
			OnSave: __gong__New__TreeFormCallback(
				instancesTyped,
				playground,
			),
		}).Stage(formStage)
		FillUpForm(instancesTyped, formGroup, playground)
	default:
		_ = instancesTyped
	}
	formStage.Commit()
}
