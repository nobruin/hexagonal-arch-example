package app

const (
	DISABLED                     = "disabled"
	ENABLED                      = "enabled"
	ENABLED_ERROR_MESSAGE        = "the price must be greater than zero to enable the product"
	DISABLED_ERROR_MESSAGE       = "the price must be zero in order to have the product disabled"
	STATUS_INVALID_ERROR_MESSAGE = "the status must be enabled or disabled"
	ISVALID_PRICE_ERROR_MESSAGE  = "the price must be greater than or equal zero"
	CLI_RESULT_MESSAGE_CREATED   = "Product ID %s with name %s has ben created with price %f and status: %s"
)
