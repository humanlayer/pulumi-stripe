package shim

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/stripe/terraform-provider-stripe/internal/provider"
)

// NewProvider returns a new Stripe Terraform provider.
func NewProvider() *schema.Provider {
	return provider.New("")()
}
