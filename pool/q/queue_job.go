// This file is autogenerated by hexya-generate
// DO NOT MODIFY THIS FILE - ANY CHANGES WILL BE OVERWRITTEN

package q

import (
	"github.com/hexya-erp/hexya/src/models"
	"github.com/hexya-erp/pool/q/queue_job"
)

type QueueJobCondition = queue_job.Condition

// QueueJob returns a queue_job.ConditionStart for QueueJobModel
func QueueJob() queue_job.ConditionStart {
	return queue_job.ConditionStart{
		ConditionStart: &models.ConditionStart{},
	}
}