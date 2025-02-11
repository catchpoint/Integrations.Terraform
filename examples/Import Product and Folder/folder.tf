terraform {
  required_providers {
    catchpoint = {
      source  = "catchpoint/catchpoint"
      version = "1.5.0"
    }
  }
}

provider "catchpoint" {
api_token="5618ABF44CA1117B4286C9572XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
}

resource "manage_folder" "folder"{
    provider = catchpoint
    id="42632"
}

# =========================================================
# Command to import product details:
# terraform import manage_product.product 42632