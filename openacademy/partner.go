package openacademy

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/pool/h"
)

func init() {
    h.Partner().AddFields(map[string]models.FieldDefinition{
        "Instructor":       fields.Boolean{},
        "AttendedSessions": fields.Many2Many{RelationModel: h.OpenAcademySession()},
    })
}