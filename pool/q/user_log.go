// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package q

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/q/user_log"
)

type UserLogCondition = user_log.Condition

// UserLog returns a user_log.ConditionStart for UserLogModel
func UserLog() user_log.ConditionStart {
	return user_log.ConditionStart{
		ConditionStart: &models.ConditionStart{},
	}
}
