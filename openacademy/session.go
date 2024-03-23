package openacademy

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/hexya/src/models/fields"
	"github.com/hexya-erp/hexya/src/models/types"
	"github.com/hexya-erp/hexya/src/models/types/dates"
	"github.com/hexya-erp/hexya/src/tools/nbutils"
	"github.com/hexya-erp/pool/h"
)

func init() {
    models.NewModel("OpenAcademySession")
    h.OpenAcademySession().AddFields(map[string]models.FieldDefinition{
        "Name":      fields.Char{Required: true},
        "StartDate": fields.Date{Default: func(env models.Environment) interface{} {
			return dates.Today()
		}},
		"Active":    fields.Boolean{Default: models.DefaultValue(true)},
        "Duration":  fields.Float{Digits: nbutils.Digits{Precision: 6, Scale: 2},
            Help: "Duration in days"},
        "Seats":     fields.Integer{String: "Number of seats"},
        "State":     fields.Selection{
            Selection: types.Selection{"planned": "Planned", "in_progress": "In Progress", "done": "Finished"}},
		"Instructor": fields.Many2One{RelationModel: h.Partner()},
		"Course":     fields.Many2One{RelationModel: h.OpenAcademyCourse(),
				Required: true, OnDelete: models.Cascade},
		"Attendees":  fields.Many2Many{RelationModel: h.Partner()},
		"TakenSeats": fields.Float{
			// Compute: h.OpenAcademySession().Methods().ComputeTakenSeats(),
			Depends: []string{"Seats", "Attendees"}},
    })
	// h.OpenAcademySession().Methods().ComputeTakenSeats().DeclareMethod(
    //     `ComputeTakenSeats returns the percentage of taken seats in this session`,
    //     func(rs m.OpenAcademySessionSet) m.OpenAcademySessionData {
    //         res h.OpenAcademySessionData().NewData()
    //         if rs.Seats() != 0 {
    //             res.SetTakenSeats(100.0 * float64(rs.Attendees().Len()) / float64(rs.Seats()))
    //         }
    //         return res
    //     })
}