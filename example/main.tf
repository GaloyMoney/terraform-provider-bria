terraform {
  required_providers {
    briaadmin = {
      source  = "galoymoney/briaadmin"
      version = "0.0.6"
    }
    bria = {
      source  = "galoymoney/bria"
      version = "0.0.6"
    }
  }
}

resource "random_string" "postfix" {
  length  = 6
  special = false
  upper   = false
  numeric  = false
}

resource "briaadmin_account" "example" {
  name = "tf-example-${random_string.postfix.result}"
}

provider "bria" {
  api_key = briaadmin_account.example.api_key
}

resource "bria_profile" "example" {
  name = "profile-${random_string.postfix.result}"
}

resource "bria_api_key" "example" {
  profile = bria_profile.example.name
}

output "api_key" {
  value = bria_api_key.example.key
  sensitive = true
}

resource "bria_xpub" "lnd" {
  name       = "lnd"
  xpub       = "tpubDDEGUyCLufbxAfQruPHkhUcu55UdhXy7otfcEQG4wqYNnMfq9DbHPxWCqpEQQAJUDi8Bq45DjcukdDAXasKJ2G27iLsvpdoEL5nTRy5TJ2B"
  derivation = "m/64h/1h/0"
}

resource "bria_wallet" "example" {
  name  = "example"
  keychain {
    wpkh {
      xpub = bria_xpub.lnd.id
    }
  }
}

resource "bria_signer_config" "lnd" {
  xpub = bria_xpub.lnd.id
  lnd {
    endpoint = "localhost:10009"
    macaroon_base64 = "AgEDbG5kAvgBAwoQB1FdhGa9xoewc1LEXmnURRIBMBoWCgdhZGRyZXNzEgRyZWFkEgV3cml0ZRoTCgRpbmZvEgRyZWFkEgV3cml0ZRoXCghpbnZvaWNlcxIEcmVhZBIFd3JpdGUaIQoIbWFjYXJvb24SCGdlbmVyYXRlEgRyZWFkEgV3cml0ZRoWCgdtZXNzYWdlEgRyZWFkEgV3cml0ZRoXCghvZmZjaGFpbhIEcmVhZBIFd3JpdGUaFgoHb25jaGFpbhIEcmVhZBIFd3JpdGUaFAoFcGVlcnMSBHJlYWQSBXdyaXRlGhgKBnNpZ25lchIIZ2VuZXJhdGUSBHJlYWQAAAYgqHDdwGCqx0aQL1/Z3uUfzCpeBhfapGf9s/AZPOVwf6s="
    cert = <<EOT
-----BEGIN CERTIFICATE-----
MIICTzCCAfagAwIBAgIRAN7zELSxwC0+P97mtkLTDeMwCgYIKoZIzj0EAwIwODEf
MB0GA1UEChMWbG5kIGF1dG9nZW5lcmF0ZWQgY2VydDEVMBMGA1UEAxMMYWI4NDIz
NGJlMTEzMB4XDTIyMDkyMjA4MjQ0NloXDTM0MDMyNDA4MjQ0NlowODEfMB0GA1UE
ChMWbG5kIGF1dG9nZW5lcmF0ZWQgY2VydDEVMBMGA1UEAxMMYWI4NDIzNGJlMTEz
MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE+fYHYw5VuOsVn7kCqy4dvK99y2OP
0A//zHN52G5Nm6apoyQlvbjeCyVVmz63uit3yIprAXAmv9ca8RPC77XZ+qOB4DCB
3TAOBgNVHQ8BAf8EBAMCAqQwEwYDVR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/
BAUwAwEB/zAdBgNVHQ4EFgQUU1Kg1ObzxEmK1rWwkgvQ+lObIiUwgYUGA1UdEQR+
MHyCDGFiODQyMzRiZTExM4IJbG9jYWxob3N0gg1sbmQtb3V0c2lkZS0xgg1sbmQt
b3V0c2lkZS0yggRsbmQxggRsbmQyggR1bml4ggp1bml4cGFja2V0ggdidWZjb25u
hwR/AAABhxAAAAAAAAAAAAAAAAAAAAABhwSsGQAMMAoGCCqGSM49BAMCA0cAMEQC
IHJfxAKakV11FTi0qlZ8/5Z7zn4Rize4UnFEKlHrwcp4AiA59nQlOBqb8RtpCkqd
FvV1D2W0uFGiZLTHgfh4VaHMyA==
-----END CERTIFICATE-----
EOT
  }
}

resource "bria_payout_queue" "interval" {
  name = "interval-${random_string.postfix.result}"
  description = "An example Bria batch group"

  config {
    tx_priority = "ECONOMY"
    consolidate_deprecated_keychains = true
    interval_secs = 3600
  }
}

