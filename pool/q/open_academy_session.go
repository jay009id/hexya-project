// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package q

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/q/open_academy_session"
)

type OpenAcademySessionCondition = open_academy_session.Condition

// OpenAcademySession returns a open_academy_session.ConditionStart for OpenAcademySessionModel
func OpenAcademySession() open_academy_session.ConditionStart {
	return open_academy_session.ConditionStart{
		ConditionStart: &models.ConditionStart{},
	}
}
