package provider

import (
	"fmt"

	"github.com/GaloyMoney/terraform-provider-bria/bria"
	briav1 "github.com/GaloyMoney/terraform-provider-bria/bria/proto/api"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourcePayoutQueue() *schema.Resource {
	return &schema.Resource{
		Create: resourcePayoutQueueCreate,
		Read:   resourcePayoutQueueRead,
		Update: resourcePayoutQueueUpdate,
		Delete: resourcePayoutQueueDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"config": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tx_priority": {
							Type:     schema.TypeString,
							Required: true,
						},
						"consolidate_deprecated_keychains": {
							Type:     schema.TypeBool,
							Required: true,
						},
						"interval_secs": {
							Type:     schema.TypeInt,
							Required: true,
						},
						"cpfp_payouts_after_mins": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  -1,
						},
						"cpfp_payouts_after_blocks": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  -1,
						},
						"force_min_change_sats": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  -1,
						},
					},
				},
			},
		},
	}
}

func resourcePayoutQueueCreate(d *schema.ResourceData, m interface{}) error {
	client := m.(*bria.AccountClient)

	name := d.Get("name").(string)
	description := d.Get("description").(string)
	configData := d.Get("config").([]interface{})[0].(map[string]interface{})

	config := &briav1.PayoutQueueConfig{
		TxPriority:                     briav1.TxPriority(briav1.TxPriority_value[configData["tx_priority"].(string)]),
		ConsolidateDeprecatedKeychains: configData["consolidate_deprecated_keychains"].(bool),
	}

	config.Trigger = &briav1.PayoutQueueConfig_IntervalSecs{IntervalSecs: uint32(configData["interval_secs"].(int))}

	if val, ok := configData["cpfp_payouts_after_mins"]; ok {
		if val.(int) >= 0 {
			tempVal := uint32(val.(int))
			config.CpfpPayoutsAfterMins = &tempVal
		}
	}
	if val, ok := configData["cpfp_payouts_after_blocks"]; ok {
		if val.(int) >= 0 {
			tempVal := uint32(val.(int))
			config.CpfpPayoutsAfterBlocks = &tempVal
		}
	}
	if val, ok := configData["force_min_change_sats"]; ok {
		if val.(int) >= 0 {
			tempVal := uint64(val.(int))
			config.ForceMinChangeSats = &tempVal
		}
	}

	res, err := client.CreatePayoutQueue(name, description, config)
	if err != nil {
		return fmt.Errorf("error creating Bria batch group: %w", err)
	}

	d.SetId(res.Id)

	return resourcePayoutQueueRead(d, m)
}

func resourcePayoutQueueRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bria.AccountClient)

	queueId := d.Id()

	queue, err := client.ReadPayoutQueue(queueId)
	if err != nil {
		return fmt.Errorf("error reading Bria queue: %w", err)
	}

	if queue == nil {
		// queue was deleted
		d.SetId("")
		return nil
	}

	d.Set("id", queue.Id)
	d.Set("name", queue.Name)
	d.Set("description", queue.Description)

	if queue.Config != nil {
		config := map[string]interface{}{
			"tx_priority":                      queue.Config.TxPriority.String(),
			"consolidate_deprecated_keychains": queue.Config.ConsolidateDeprecatedKeychains,
		}

		if intervalSecs, ok := queue.Config.Trigger.(*briav1.PayoutQueueConfig_IntervalSecs); ok {
			config["interval_secs"] = intervalSecs.IntervalSecs
		}

		if queue.Config.CpfpPayoutsAfterMins != nil {
			config["cpfp_payouts_after_mins"] = *queue.Config.CpfpPayoutsAfterMins
		} else {
			config["cpfp_payouts_after_mins"] = -1
		}
		if queue.Config.CpfpPayoutsAfterBlocks != nil {
			config["cpfp_payouts_after_blocks"] = *queue.Config.CpfpPayoutsAfterBlocks
		} else {
			config["cpfp_payouts_after_blocks"] = -1
		}
		if queue.Config.ForceMinChangeSats != nil {
			config["force_min_change_sats"] = *queue.Config.ForceMinChangeSats
		} else {
			config["force_min_change_sats"] = -1
		}

		if err := d.Set("config", []interface{}{config}); err != nil {
			return fmt.Errorf("error setting config: %w", err)
		}
	}

	return nil
}

func resourcePayoutQueueUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*bria.AccountClient)

	queueId := d.Id()
	description := d.Get("description").(string)
	configData := d.Get("config").([]interface{})[0].(map[string]interface{})

	config := &briav1.PayoutQueueConfig{
		TxPriority:                     briav1.TxPriority(briav1.TxPriority_value[configData["tx_priority"].(string)]),
		ConsolidateDeprecatedKeychains: configData["consolidate_deprecated_keychains"].(bool),
	}

	config.Trigger = &briav1.PayoutQueueConfig_IntervalSecs{IntervalSecs: uint32(configData["interval_secs"].(int))}

	if val, ok := configData["cpfp_payouts_after_mins"]; ok {
		if val.(int) >= 0 {
			tempVal := uint32(val.(int))
			config.CpfpPayoutsAfterMins = &tempVal
		}
	}
	if val, ok := configData["cpfp_payouts_after_blocks"]; ok {
		if val.(int) >= 0 {
			tempVal := uint32(val.(int))
			config.CpfpPayoutsAfterBlocks = &tempVal
		}
	}
	if val, ok := configData["force_min_change_sats"]; ok {
		if val.(int) >= 0 {
			tempVal := uint64(val.(int))
			config.ForceMinChangeSats = &tempVal
		}
	}

	_, err := client.UpdatePayoutQueue(queueId, description, config)
	if err != nil {
		return fmt.Errorf("error updating Bria payout queue: %w", err)
	}

	return resourcePayoutQueueRead(d, m)
}

func resourcePayoutQueueDelete(d *schema.ResourceData, meta interface{}) error {
	// Implement the delete function for the bria_account_signer_config resource
	// If the API does not provide a delete functionality, you can set the ID to an empty string
	d.SetId("")
	return nil
}
