/*
 * Author: Samsul Ma'arif <samsulma828@gmail.com>
 * Copyright (c) 2023.
 */

package lib

const (
	ROLE_SUPER_ADMIN_ID = iota + 1
	ROLE_ADMIN_ID
)

const (
	SheetFilenameImportProduct = "product"

	InvoiceStatusPaid     = "PAID"
	InvoiceStatusAwaiting = "AWAITING"

	SubscriptionPlanActive   = "ACTIVE"
	SubscriptionPlanAwaiting = "AWAITING"
	SubscriptionPlanSuspend  = "SUSPEND"

	SuperAdminPermission = "*"
)

var (
	Roles = map[int]string{
		ROLE_SUPER_ADMIN_ID: "SUPER ADMIN",
		ROLE_ADMIN_ID:       "ADMIN",
	}
)
