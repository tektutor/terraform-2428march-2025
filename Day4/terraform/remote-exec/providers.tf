terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "3.0.2"
    }
    local = {
      source = "hashicorp/local"
      version = "2.5.2"
    }
  }
}

# This is an instance of the provider with optional configurations
provider "docker" {
  # Configuration options
}

provider "local" {
  # Configuration options
}
