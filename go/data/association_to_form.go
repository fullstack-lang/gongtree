// generated code - do not edit
package data

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gongtree/go/models"
)

// AssociationFieldToForm will append a div to the form
// with the values of options equal to the name of each possible instances in the association type
func AssociationFieldToForm[FieldType models.PointerToGongstruct](
	fieldName string, field FieldType, stageOfInterest *models.StageStruct, formStage *form.StageStruct, formGroup *form.FormGroup,
) {

	formDiv := (&form.FormDiv{
		Name: fieldName,
	}).Stage(formStage)
	formGroup.FormDivs = append(formGroup.FormDivs, formDiv)
	formField := (&form.FormField{
		Name:        fieldName,
		Label:       fieldName,
		Placeholder: "",
	}).Stage(formStage)
	formDiv.FormFields = append(formDiv.FormFields, formField)

	formFieldSelect := (&form.FormFieldSelect{
		Name:       "association",
		CanBeEmpty: true,
	}).Stage(formStage)
	formField.FormFieldSelect = formFieldSelect

	// generate one option per possible instances for the field
	formField.FormFieldSelect.Options = make([]*form.Option, 0)
	for instance := range *models.GetGongstructInstancesSetFromPointerType[FieldType](stageOfInterest) {
		option := (&form.Option{
			Name: instance.GetName(),
		}).Stage(formStage)

		// set up select value if field matches the instance
		if instance == field {
			formFieldSelect.Value = option
		}

		formField.FormFieldSelect.Options =
			append(formField.FormFieldSelect.Options, option)
	}
}