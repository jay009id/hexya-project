// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package q

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/q/partner_category"
)

type PartnerCategoryCondition = partner_category.Condition

// PartnerCategory returns a partner_category.ConditionStart for PartnerCategoryModel
func PartnerCategory() partner_category.ConditionStart {
	return partner_category.ConditionStart{
		ConditionStart: &models.ConditionStart{},
	}
}