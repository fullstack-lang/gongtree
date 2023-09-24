// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtree/go/models"
)

func FillUpFormFromGongstructName(
	playground *Playground,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := playground.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Button":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Button Form",
			OnSave: __gong__New__ButtonFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		button := new(models.Button)
		FillUpForm(button, formGroup, playground)
	case "Node":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Node Form",
			OnSave: __gong__New__NodeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		node := new(models.Node)
		FillUpForm(node, formGroup, playground)
	case "Tree":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Tree Form",
			OnSave: __gong__New__TreeFormCallback(
				nil,
				playground,
			),
		}).Stage(formStage)
		tree := new(models.Tree)
		FillUpForm(tree, formGroup, playground)
	}
	formStage.Commit()
}
