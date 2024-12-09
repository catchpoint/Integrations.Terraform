package catchpoint

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceManageProduct() *schema.Resource {
	return &schema.Resource{
		Create: resourceProductCreate,
		Read:   resourceProductRead,
		Update: resourceProductUpdate,
		Delete: resourceProductDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"division_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The Division where the Product will be created",
			},
			"product_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the product",
			},
			"status": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Optional. Product status: active or inactive",
				ValidateFunc: validation.StringInSlice([]string{"active", "inactive"}, false),
			},
			"test_data_webhook_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Optional. test data webhook id",
			},
			"alert_group_id": {
				Type:        schema.TypeInt,
				Optional:    true,
				Computed:    true,
				Description: "Optional. alert group id",
			},
			"request_settings": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Optional. Used for overriding authentication and HTTP request headers",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:     schema.TypeSet,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_type": {
										Type:         schema.TypeString,
										Required:     true,
										Description:  "Type of authentication to use 'basic', 'ntlm', 'digest', 'login'",
										ValidateFunc: validation.StringInSlice([]string{"basic", "ntlm", "digest", "login"}, false),
									},
									"password_ids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Optional. Password ids in a list",
										Sensitive:   true,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
								},
							},
						},
						"token_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Optional. Token ids in a list",
							Sensitive:   true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"library_certificate_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Optional. Library certificate ids in a list",
							Sensitive:   true,
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"http_request_headers": {
							Type:     schema.TypeSet,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"user_agent": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the user agent header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"accept": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the accept header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"accept_encoding": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the user accept encoding header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"accept_language": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the accept language header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"accept_charset": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the accept charset header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"cookie": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the cookie header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"cache_control": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the cache control header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"pragma": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the pragma header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"referer": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the referer header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"host": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the host header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"dns_override": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the dns override header for the given child_host_pattern",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},
									"request_override": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the request override header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"request_block": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the request block header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"request_delay": {
										Type:        schema.TypeSet,
										MaxItems:    1,
										Optional:    true,
										Description: "Optional. Sets the request delay header for product url if child_host_pattern attribute is omitted",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": {
													Type:     schema.TypeString,
													Required: true,
												},
												"child_host_pattern": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"insights": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Optional. Used for overriding the insights section",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"tracepoint_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Optional. Tracepoint ids in a list",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"indicator_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Optional. Indicator ids in a list",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
					},
				},
			},
			"schedule_settings": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Optional. Used for overriding the schedule section",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"run_schedule_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Optional. The run schedule id to utilize for the test",
						},
						"maintenance_schedule_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Optional. The maintenance schedule id to utilize for the test",
						},
						"frequency": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "5 minutes",
							Description:  "Sets the scheduling frequency: '1 minute', '5 minutes', '10 minutes', '15 minutes', '20 minutes', '30 minutes', '60 minutes', '2 hours', '3 hours', '4 hours', '6 hours', '8 hours', '12 hours', '24 hours', '4 minutes', '2 minutes'",
							ValidateFunc: validation.StringInSlice([]string{"1 minute", "5 minutes", "10 minutes", "15 minutes", "20 minutes", "30 minutes", "60 minutes", "2 hours", "3 hours", "4 hours", "6 hours", "8 hours", "12 hours", "24 hours", "4 minutes", "2 minutes"}, false),
						},
						"node_distribution": {
							Type:         schema.TypeString,
							Optional:     true,
							Default:      "random",
							Description:  "Node distribution type: 'random' or 'concurrent'",
							ValidateFunc: validation.StringInSlice([]string{"random", "concurrent"}, false),
						},
						"node_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Node ids in a list",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"node_group_ids": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "Node group ids in a list",
							Elem: &schema.Schema{
								Type: schema.TypeInt,
							},
						},
						"no_of_subset_nodes": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Optional. Number of subset nodes",
						},
					},
				},
			},
			"alert_settings": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				MaxItems:    1,
				Description: "Optional. Used for overriding the alert section",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_rule": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Optional. Sets the alert rule with attributes such as threshold, trigger type, warning, critical trigger and more",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"node_threshold_type": {
										Type:         schema.TypeString,
										Required:     true,
										Description:  "Sets the node threshold type for alert: 'runs', 'average across node' or 'node'",
										ValidateFunc: validation.StringInSlice([]string{"runs", "average across nodes", "node"}, false),
									},
									"threshold_number_of_runs": {
										Type:        schema.TypeInt,
										Description: "Optional. Sets the threshold for the number of runs or nodes the alert should trigger",
										Optional:    true,
									},
									"threshold_percentage_of_runs": {
										Type:        schema.TypeFloat,
										Description: "Optional. Sets the threshold for the percentage of runs the alert should trigger",
										Optional:    true,
									},
									"number_of_failing_nodes": {
										Type:        schema.TypeInt,
										Description: "Optional. Sets the number of failed nodes the alert should trigger if node_threshold_type is 'average across nodes'",
										Optional:    true,
									},
									"trigger_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the trigger type: 'specific value', 'trailing value', 'trendshift'",
										ValidateFunc: validation.StringInSlice([]string{"specific value", "trailing value", "trendshift"}, false),
									},
									"operation_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the operation type: 'not equals', 'greater than', 'greater than or equals', 'less than', 'less than or equals'",
										ValidateFunc: validation.StringInSlice([]string{"equals", "not equals", "greater than", "greater than or equals", "less than", "less than or equals"}, false),
									},
									"statistical_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the statistical type for 'trailing value' trigger type. Supports only 'average' for now",
										ValidateFunc: validation.StringInSlice([]string{"average"}, false),
									},
									"historical_interval": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the historical interval for 'trailing value' trigger type: '5 minutes', '10 minutes', '15 minutes', '30 minutes', '1 hour', '2 hours', '6 hours', '12 hours', '1 day', '1 week'",
										ValidateFunc: validation.StringInSlice([]string{"5 minutes", "10 minutes", "15 minutes", "30 minutes", "1 hour", "2 hours", "6 hours", "12 hours", "1 day", "1 week"}, false),
									},
									"warning_trigger": {
										Type:        schema.TypeFloat,
										Description: "Optional. Warning trigger value for 'specific value' and 'trailing value' trigger types.",
										Optional:    true,
									},
									"critical_trigger": {
										Type:        schema.TypeFloat,
										Description: "Optional. Critical trigger value for 'specific value' and 'trailing value' trigger types.",
										Optional:    true,
									},
									"enable_consecutive": {
										Type:        schema.TypeBool,
										Description: "Optional. Checks consecutive number of runs or nodes for triggering alerts.",
										Optional:    true,
										Default:     false,
									},
									"consecutive_number_of_runs": {
										Type:        schema.TypeInt,
										Description: "Optional. Sets the number of consecutive runs only if enable_consecutive field is true and node_threshold_type is node",
										Optional:    true,
									},
									"expression": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Optional. Sets trigger expression for content match alert type ",
									},
									"warning_reminder": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets alert warning reminder interval: 'none', '1 minute', '5 minutes', '10 minutes', '15 minutes', '30 minutes', '1 hour', 'daily'",
										ValidateFunc: validation.StringInSlice([]string{"none", "1 minute", "5 minutes", "10 minutes", "15 minutes", "30 minutes", "1 hour", "daily"}, false),
									},
									"critical_reminder": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets alert critical reminder interval: 'none', '1 minute', '5 minutes', '10 minutes', '15 minutes', '30 minutes', '1 hour', 'daily'",
										ValidateFunc: validation.StringInSlice([]string{"none", "1 minute", "5 minutes", "10 minutes", "15 minutes", "30 minutes", "1 hour", "daily"}, false),
									},
									"threshold_interval": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the alert time threshold: 'default', '5 minutes', '10 minutes', '15 minutes', '30 minutes', '1 hour', '2 hours', '6 hours', '12 hours'",
										ValidateFunc: validation.StringInSlice([]string{"default", "5 minutes", "10 minutes", "15 minutes", "30 minutes", "1 hour", "2 hours", "6 hours", "12 hours"}, false),
									},
									"use_rolling_window": {
										Type:        schema.TypeBool,
										Description: "Optional. Set to true for using rolling window instead of schedule time threshold",
										Optional:    true,
										Default:     false,
									},
									"notification_type": {
										Type:         schema.TypeString,
										Description:  "Optional. Notification group type to alert. Supports only default contacts for now.",
										Optional:     true,
										ValidateFunc: validation.StringInSlice([]string{"default contacts"}, false),
									},
									"alert_type": {
										Type:         schema.TypeString,
										Description:  "Sets the alert type: 'test failure', 'timing', 'availability'",
										Required:     true,
										ValidateFunc: validation.StringInSlice([]string{"test failure", "timing", "availability", "host failure", "requests", "content match", "byte length"}, false),
									},
									"alert_sub_type": {
										Type:         schema.TypeString,
										Optional:     true,
										Description:  "Optional. Sets the sub alert type: 'dns', 'connect', 'send', 'wait', 'load', 'ttfb', 'content load', 'response', 'test time', 'dom load', 'test time with suspect', 'server response', 'document complete', 'redirect', 'test', 'content', '% downtime'",
										ValidateFunc: validation.StringInSlice([]string{"dns", "connect", "send", "wait", "load", "ttfb", "content load", "response", "test time", "dom load", "test time with suspect", "server response", "document complete", "redirect", "test", "content", "% downtime", "# requests", "# hosts", "# connections", "# redirects", "# other", "# images", "# scripts", "# html", "# css", "# xml", "# flash", "# media", "regular expression", "response code", "response headers", "byte length", "page", "file size"}, false),
									},
									"enforce_test_failure": {
										Type:        schema.TypeBool,
										Description: "Optional. Sets enforce test failure property for an alert",
										Optional:    true,
										Default:     false,
									},
									"omit_scatterplot": {
										Type:        schema.TypeBool,
										Description: "Optional. Omits scatterplot image from alert emails if set to true",
										Optional:    true,
										Default:     false,
									},
									"notification_group": {
										Type:        schema.TypeSet,
										Required:    true,
										MaxItems:    5,
										Description: "List of Notification groups for configuring alert notifications, including recipients' email addresses and alert settings. To ensure either recipient_email_ids or contact_groups is provided",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"notify_on_warning": {
													Type:        schema.TypeBool,
													Description: "Optional. Set to true to include warning alerts in notifications. Default is false.",
													Optional:    true,
													Default:     false,
												},
												"notify_on_critical": {
													Type:        schema.TypeBool,
													Description: "Optional. Set to true to include critical alerts in notifications. Default is false.",
													Optional:    true,
													Default:     false,
												},
												"notify_on_improved": {
													Type:        schema.TypeBool,
													Description: "Optional. Set to true to include improved alerts in notifications. Default is false.",
													Optional:    true,
													Default:     false,
												},
												"subject": {
													Type:        schema.TypeString,
													Description: "Email subject for the alert notifications. Required field.",
													Required:    true,
												},
												"recipient_email_ids": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Optional. List of email addresses to receive alert notifications. To ensure either recipient_email_ids or contact_groups is provided",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"contact_groups": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Optional. List of contact groups to receive alert notifications. To ensure either recipient_email_ids or contact_groups is provided",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"notification_group": {
							Type:        schema.TypeSet,
							Required:    true,
							MaxItems:    1,
							Description: "Notification group for setting up alert recipients, adding alert webhook ids. To ensure either recipient_email_ids or contact_groups is provided",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"subject": {
										Type:        schema.TypeString,
										Description: "Email subject for the alert notifications. Required field.",
										Required:    true,
									},
									"alert_webhook_ids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Optional. Alert webhook ids for the webhook endpoints to associate with this alert setting.",
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"recipient_email_ids": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Optional. List of emails to alert. To ensure either recipient_email_ids or contact_groups is provided",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"contact_groups": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Optional. List of contact groups to receive alert notifications. To ensure either recipient_email_ids or contact_groups is provided",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"advanced_settings": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Optional. Used for overriding the advanced settings",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_path_mtu_discovery": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables Path MTU Discovery",
							Optional:    true,
							Default:     false,
						},
						"verify_test_on_failure": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables verify on test failure setting",
							Optional:    true,
							Default:     false,
						},
						"debug_primary_host_on_failure": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables debug primary host on failure setting",
							Optional:    true,
							Default:     false,
						},
						"enable_http2": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables enable http/2 setting",
							Optional:    true,
							Default:     false,
						},
						"debug_referenced_hosts_on_failure": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables debug referenced hosts on failure setting",
							Optional:    true,
							Default:     false,
						},
						"capture_http_headers": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables capture http headers setting for all runs",
							Optional:    true,
							Default:     false,
						},
						"capture_response_content": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables capture response content setting for all runs",
							Optional:    true,
							Default:     false,
						},
						"ignore_ssl_failures": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables ignore SSL failures setting",
							Optional:    true,
							Default:     false,
						},
						"host_data_collection_enabled": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables host data collection setting",
							Optional:    true,
							Default:     false,
						},
						"zone_data_collection_enabled": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables zone data collection setting",
							Optional:    true,
							Default:     false,
						},
						"f40x_or_50x_http_mark_successful": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables 40x or 50x error mark successful setting",
							Optional:    true,
							Default:     false,
						},
						"t30x_redirects_do_not_follow": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables 30x redirects do not follow setting",
							Optional:    true,
							Default:     false,
						},
						"enable_self_versus_third_party_zones": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables self versus third party zones setting and matches self zone by test URL",
							Optional:    true,
							Default:     false,
						},
						"allow_test_download_limit_override": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables test download limit override setting",
							Optional:    true,
							Default:     false,
						},
						"capture_filmstrip": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables capture filmstrip setting",
							Optional:    true,
							Default:     false,
						},
						"capture_screenshot": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables capture screenshot setting for all runs",
							Optional:    true,
							Default:     false,
						},
						"stop_test_on_document_complete": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables stop test on document complete setting",
							Optional:    true,
							Default:     false,
						},
						"disable_cross_origin_iframe_access": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables disable cross origin iframe access setting for chrome monitor",
							Optional:    true,
							Default:     false,
						},
						"stop_test_on_dom_content_load": {
							Type:        schema.TypeBool,
							Description: "Optional. True enables stop test on DOM content load setting",
							Optional:    true,
							Default:     false,
						},
						"enforce_test_failure_if_runs_longer_than": {
							Type:         schema.TypeInt,
							Description:  "Optional. Set the time value in seconds post which the test will be marked as failure.",
							ValidateFunc: validation.IntInSlice([]int{5, 10, 15, 20, 30, 60, 90, 120}),
							Optional:     true,
						},
						"wait_for_no_activity": {
							Type:         schema.TypeInt,
							Description:  "Optional. Set the time value in ms to stop the test after no network activity on document complete. Use with stop_test_on_document_complete flag",
							ValidateFunc: validation.IntInSlice([]int{0, 500, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000}),
							Optional:     true,
						},
						"viewport_height": {
							Type:        schema.TypeInt,
							Description: "Optional. Set the viewport height. Use with viewport_width attribute",
							Optional:    true,
						},
						"viewport_width": {
							Type:        schema.TypeInt,
							Description: "Optional. Set the viewport width. Use with viewport_height attribute",
							Optional:    true,
						},
						"additional_monitor": {
							Type:         schema.TypeString,
							Description:  "Optional. Set the additional monitor to run along with the test monitor: 'ping icmp', 'ping tcp', 'ping udp','traceroute icmp','traceroute udp','traceroute tcp'",
							ValidateFunc: validation.StringInSlice([]string{"ping icmp", "ping tcp", "ping udp", "traceroute icmp", "traceroute udp", "traceroute tcp"}, false),
							Optional:     true,
						},
						"bandwidth_throttling": {
							Type:         schema.TypeString,
							Description:  "Optional. Set the bandwidth throttling for chrome: 'gprs','regular 2g','good 2g','regular 3g','good 3g','regular 4g','dsl','wifi'",
							ValidateFunc: validation.StringInSlice([]string{"gprs", "regular 2g", "good 2g", "regular 3g", "good 3g", "regular 4g", "dsl", "wifi"}, false),
							Optional:     true,
						},
					},
				},
			},
		},
	}
}
func resourceProductRead(d *schema.ResourceData, m interface{}) error {
	productId := d.Id()
	api_token := m.(*Config).ApiToken

	log.Printf("[DEBUG] Fetching Product: %v", productId)

	product, respStatus, err := getProduct(api_token, productId)
	if err != nil {
		return err
	}
	if respStatus != "200 ok" {
		log.Printf("[ERROR] Error while reading product: %v", productId)
		return errors.New(respStatus)
	}
	if product == nil {
		d.SetId("")
		log.Printf("[DEBUG] Product not found %v", productId)
		return nil
	}
	log.Printf("[DEBUG] Response Code from Catchpoint API: " + respStatus)

	productNew := flattenProduct(product)
	d.Set("division_id", productNew["division_id"])
	d.Set("product_name", productNew["product_name"])
	d.Set("status", productNew["status"])
	d.Set("alert_group_id", productNew["alert_group_id"])
	d.Set("test_data_webhook_id", productNew["test_data_webhook_id"])
	d.Set("insights", productNew["insights"])
	d.Set("schedule_settings", productNew["schedule_settings"])
	d.Set("advanced_settings", productNew["advanced_settings"])
	d.Set("request_settings", productNew["request_settings"])
	d.Set("alert_settings", productNew["alert_settings"])

	return nil
}

func resourceProductCreate(d *schema.ResourceData, m interface{}) error {
	api_token := m.(*Config).ApiToken
	division_id := d.Get("division_id").(int)
	product_name := d.Get("product_name").(string)
	status := d.Get("status").(string)
	test_data_webhook_id := d.Get("test_data_webhook_id").(int)
	alert_group_id := d.Get("alert_group_id").(int)
	status_id := getTestStatusTypeId(status)
	var productConfig = ProductConfig{}
	productConfig = ProductConfig{
		DivisionId:        division_id,
		ProductName:       product_name,
		ProductStatus:     status_id,
		TestDataWebhookId: test_data_webhook_id,
		AlertGroupId:      alert_group_id,
	}
	schedule_settings, schedule_settingsOk := d.GetOk("schedule_settings")
	if schedule_settingsOk {
		schedule_setting_list := schedule_settings.(*schema.Set).List()
		schedule_setting := schedule_setting_list[0].(map[string]interface{})

		err := configureProductScheduleSettings(schedule_setting, &productConfig)
		if err != nil {
			return err
		}

	}
	alert_settings, alert_settingsOk := d.GetOk("alert_settings")
	if alert_settingsOk {
		alert_setting_list := alert_settings.(*schema.Set).List()
		alert_setting := alert_setting_list[0].(map[string]interface{})

		err := configureProductAlertSettings(alert_setting, &productConfig)
		if err != nil {
			return err
		}
	}

	advanced_settings, advanced_settingsOk := d.GetOk("advanced_settings")
	if advanced_settingsOk {
		advanced_setting_list := advanced_settings.(*schema.Set).List()
		advanced_setting := advanced_setting_list[0].(map[string]interface{})

		configureProductAdvancedSettings(advanced_setting, &productConfig)
	}
	request_settings, request_settingsOk := d.GetOk("request_settings")
	if request_settingsOk {
		request_settings_list := request_settings.(*schema.Set).List()

		request_setting := request_settings_list[0].(map[string]interface{})

		err := configureProductRequestSettings(request_setting, &productConfig)
		if err != nil {
			return err
		}
	}

	insight_settings, insight_settingsOk := d.GetOk("insights")
	if insight_settingsOk {
		insight_setting_list := insight_settings.(*schema.Set).List()
		insight_setting := insight_setting_list[0].(map[string]interface{})

		configureProductInsightSettings(insight_setting, &productConfig)
	}

	jsonStr := createProductJson(productConfig)
	if m.(*Config).LogJson {
		log.Printf("[Product JSON] \n" + jsonStr)
	}

	log.Printf("[DEBUG] Creating product: " + product_name)
	respBody, respStatus, productId, err := createProduct(api_token, jsonStr)
	if err != nil {
		log.Fatal(err)
	}
	if respStatus != "200 ok" {
		log.Printf("[ERROR] Error while creating product: " + product_name)
		log.Printf("[ERROR] Error description: " + respBody)
		return errors.New(respStatus)
	}

	log.Printf("[DEBUG] Response Code from Catchpoint API: " + respStatus)
	log.Print(respBody)
	d.SetId(productId)

	return resourceProductRead(d, m)
}

func resourceProductUpdate(d *schema.ResourceData, m interface{}) error {
	productId := d.Id()
	api_token := m.(*Config).ApiToken
	var productConfig = ProductConfig{}
	var jsonPatchDocs = []string{}

	if d.HasChange("product_name") {
		productConfigUpdate := ProductConfigUpdate{
			UpdatedFieldValue: d.Get("product_name").(string),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonProductPatchDocument(productConfigUpdate, "/name", true))
	}
	if d.HasChange("status") {
		updated_status_id := getTestStatusTypeId(d.Get("status").(string))
		productConfigUpdate := ProductConfigUpdate{
			UpdatedFieldValue: strconv.Itoa(updated_status_id),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonProductPatchDocument(productConfigUpdate, "/status", true))
	}

	if d.HasChange("test_data_webhook_id") {
		productConfigUpdate := ProductConfigUpdate{
			UpdatedFieldValue: strconv.Itoa(d.Get("test_data_webhook_id").(int)),
		}
		jsonPatchDocs = append(jsonPatchDocs, createJsonProductPatchDocument(productConfigUpdate, "/testDataWebhookId", true))

	}
	if d.HasChange("advanced_settings") {
		advanced_settings, advanced_settingsOk := d.GetOk("advanced_settings")
		if advanced_settingsOk {
			advanced_setting_list := advanced_settings.(*schema.Set).List()
			advanced_setting := advanced_setting_list[0].(map[string]interface{})

			configureProductAdvancedSettings(advanced_setting, &productConfig)

			productConfigUpdate := ProductConfigUpdate{
				UpdatedAdvancedSettingsSection: setProductAdvancedSettings(&productConfig),
				SectionToUpdate:                "/advancedSettingsModel",
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonProductPatchDocument(productConfigUpdate, productConfigUpdate.SectionToUpdate, false))
		}
	}

	if d.HasChange("request_settings") {
		request_settings, request_settingsOk := d.GetOk("request_settings")
		if request_settingsOk {
			request_settings_list := request_settings.(*schema.Set).List()
			request_setting := request_settings_list[0].(map[string]interface{})

			err := configureProductRequestSettings(request_setting, &productConfig)
			if err != nil {
				return err
			}
			productConfigUpdate := ProductConfigUpdate{
				UpdatedRequestSettingsSection: setProductRequestSettings(&productConfig),
				SectionToUpdate:               "/requestSettings",
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonProductPatchDocument(productConfigUpdate, productConfigUpdate.SectionToUpdate, false))
		}
	}
	if d.HasChange("insights") {
		insight_settings, insight_settingsOk := d.GetOk("insights")
		if insight_settingsOk {
			insight_setting_list := insight_settings.(*schema.Set).List()
			insight_setting := insight_setting_list[0].(map[string]interface{})
			updateProductInsightSettings(insight_setting, &jsonPatchDocs)

		}
	}
	if d.HasChange("schedule_settings") {
		schedule_settings, schedule_settingsOk := d.GetOk("schedule_settings")
		if schedule_settingsOk {
			schedule_setting_list := schedule_settings.(*schema.Set).List()
			schedule_setting := schedule_setting_list[0].(map[string]interface{})
			updateProductScheduleSettings(schedule_setting, &jsonPatchDocs)
		}
	}

	if d.HasChange("alert_settings") {
		alert_settings, alert_settingsOk := d.GetOk("alert_settings")
		if alert_settingsOk {
			alert_setting_list := alert_settings.(*schema.Set).List()
			alert_setting := alert_setting_list[0].(map[string]interface{})

			err := configureProductAlertSettings(alert_setting, &productConfig)
			if err != nil {
				return err
			}

			productConfigUpdate := ProductConfigUpdate{
				UpdatedAlertSettingsSection: setProductAlertSettings(&productConfig),
				SectionToUpdate:             "/alertGroup",
			}
			jsonPatchDocs = append(jsonPatchDocs, createJsonProductPatchDocument(productConfigUpdate, productConfigUpdate.SectionToUpdate, false))
		}
	}
	jsonPatchDoc := "[" + strings.Join(jsonPatchDocs, ",") + "]"

	if jsonPatchDoc != "[]" {
		log.Printf("[DEBUG] Updating product: %v", productId)
		if m.(*Config).LogJson {
			log.Printf("[DEBUG] Updating product with JSON PATCH: %v", jsonPatchDoc)
		}
		respBody, respStatus, completed, err := updateProduct(api_token, productId, jsonPatchDoc)
		if err != nil {
			log.Fatal(err)
		}
		if !completed {
			log.Printf("[ERROR] Error while Updating product: %v", productId)
			log.Printf("[ERROR] Error description: " + respBody)
			return errors.New(respBody)
		}
		log.Printf("[DEBUG] Response Code from Catchpoint API: " + respStatus)
		log.Print(respBody)
		return resourceProductRead(d, m)
	} else {
		return errors.New("no changes. Your infrastructure matches the configuration")
	}

}

func resourceProductDelete(d *schema.ResourceData, m interface{}) error {
	return errors.New("delete operation is not supported for the product")

}
