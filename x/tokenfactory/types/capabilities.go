package types

const (
	EnableSetMetadata          = "enable_metadata"
	EnableForceTransfer        = "enable_force_transfer"
	EnableBurnFrom             = "enable_burn_from"
	EnableGasConsumeCreateCost = "enable_gas_consume_create_cost"
)

func IsCapabilityEnabled(enabledCapabilities []string, capability string) bool {
	if len(enabledCapabilities) == 0 {
		return true
	}

	for _, v := range enabledCapabilities {
		if v == capability {
			return true
		}
	}

	return false
}
