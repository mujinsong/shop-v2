// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"shop-v2/internal/dao/internal"
)

// internalConsigneeInfoDao is internal type for wrapping internal DAO implements.
type internalConsigneeInfoDao = *internal.ConsigneeInfoDao

// consigneeInfoDao is the data access object for table consignee_info.
// You can define custom methods on it to extend its functionality as you wish.
type consigneeInfoDao struct {
	internalConsigneeInfoDao
}

var (
	// ConsigneeInfo is globally public accessible object for table consignee_info operations.
	ConsigneeInfo = consigneeInfoDao{
		internal.NewConsigneeInfoDao(),
	}
)

// Fill with you ideas below.
