// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package q

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/q/user_change_password_wizard_line"
)

type UserChangePasswordWizardLineCondition = user_change_password_wizard_line.Condition

// UserChangePasswordWizardLine returns a user_change_password_wizard_line.ConditionStart for UserChangePasswordWizardLineModel
func UserChangePasswordWizardLine() user_change_password_wizard_line.ConditionStart {
	return user_change_password_wizard_line.ConditionStart{
		ConditionStart: &models.ConditionStart{},
	}
}