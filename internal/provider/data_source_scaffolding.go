package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSubscription() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample data source in the Terraform provider scaffolding.",

		ReadContext: dataSourceSubscriptionRead,

		Schema: map[string]*schema.Schema{
			"organisation_id": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"name": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"start_date": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"end_date": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeString,
				Required:    true,
			},
			"total_seats": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeInt,
				Required:    true,
			},
			"seats_in_use": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeInt,
				Required:    true,
			},
		},
	}
}

func dataSourceSubscriptionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)

	subscription, err := client.client.GetSubscription()
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("organisation_id", subscription.OrganisationId); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", subscription.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("start_date", subscription.StartDate); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("end_date", subscription.EndDate); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("total_seats", subscription.TotalSeats); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("seats_in_use", subscription.SeatsInUse); err != nil {
		return diag.FromErr(err)
	}

	return diag.Errorf("not implemented")
}
