// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package q

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/q/image_mixin"
)

type ImageMixinCondition = image_mixin.Condition

// ImageMixin returns a image_mixin.ConditionStart for ImageMixinModel
func ImageMixin() image_mixin.ConditionStart {
	return image_mixin.ConditionStart{
		ConditionStart: &models.ConditionStart{},
	}
}