package openacademy

import (
    "github.com/hexya-erp/hexya/src/models"
    "github.com/hexya-erp/hexya/src/models/fields"
    "github.com/hexya-erp/hexya/src/models/types"
    "github.com/hexya-erp/pool/h"
)

func init() {
    models.NewModel("OpenAcademyCourse")
    h.OpenAcademyCourse().AddFields(map[string]models.FieldDefinition{
        "Name":        fields.Char{},
        "Description": fields.Text{},
		"Responsible": fields.Many2One{
            RelationModel: h.User(), OnDelete: models.SetNull, Index: true},
        "Sessions":    fields.One2Many{
            RelationModel: h.OpenAcademySession(), ReverseFK: "Course"},
        "State":       fields.Selection{
            Selection: types.Selection{"draft": "Draft", "confirm": "Confirm", "done": "Done"},
            String: "Status",
            Default: models.DefaultValue("draft"),
        },
    })
}