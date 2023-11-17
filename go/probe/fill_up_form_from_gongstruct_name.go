// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtree/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
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
				probe,
			),
		}).Stage(formStage)
		button := new(models.Button)
		FillUpForm(button, formGroup, probe)
	case "Node":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Node Form",
			OnSave: __gong__New__NodeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		node := new(models.Node)
		FillUpForm(node, formGroup, probe)
	case "SVGIcon":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " SVGIcon Form",
			OnSave: __gong__New__SVGIconFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		svgicon := new(models.SVGIcon)
		FillUpForm(svgicon, formGroup, probe)
	case "Tree":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Tree Form",
			OnSave: __gong__New__TreeFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		tree := new(models.Tree)
		FillUpForm(tree, formGroup, probe)
	}
	formStage.Commit()
}
