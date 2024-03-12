resource "playwright_test" "test" {
  test_name  = "playwright_TF_12"
    provider=catchpoint
    monitor="chrome"
    test_script="await page.goto('https://www.amazon.in/');"
    division_id=2633
    product_id=25232
    test_description="An object test for catchpoint.com"
    test_script_type = "playwright"
    alerts_paused=true
    enable_test_data_webhook=true
    end_time="2024-12-30T04:59:00Z"

    label {
        key="label1"
        values=["v1","v2"]
    }

    advanced_settings {
        verify_test_on_failure = false
        additional_monitor="ping icmp"
    }

    thresholds {
      test_time_warning = 47
      test_time_critical = 56
      availability_critical = 30
      availability_warning = 67
    }

    schedule_settings{
      frequency="6 hours"
      node_distribution ="random"
      no_of_subset_nodes = 5
      node_ids =[6388]
      node_group_ids =[9922,9848]
    }

    request_settings {
      authentication {
        authentication_type = "basic"
        password_ids = [2332]
      }
      token_ids = [1096]
      http_request_headers {
        user_agent {
          value = "vikash"
        }
      }
    }

    alert_settings {
        alert_rule {
            alert_type="timing"
            alert_sub_type="response"
            node_threshold_type="node"
            threshold_number_of_runs=5
            threshold_interval="30 minutes"
            trigger_type="specific value"
            warning_trigger=51
            critical_trigger=70.0
            operation_type = "less than or equals"
            use_rolling_window=true
            enable_consecutive = true
            consecutive_number_of_runs = 5
            notification_group {
              notify_on_critical = true
              subject = "contact group testing"
              recipient_email_ids = ["vkumar@catchpoint.com"]
              contact_groups = ["stage test"]
            }
        }
        notification_group {
          subject = "testing"
            recipient_email_ids=["vikash@catchpoint.com"]
        }

    }
}