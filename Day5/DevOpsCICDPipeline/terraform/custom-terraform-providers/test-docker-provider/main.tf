terraform {
  required_providers {
     docker = {
       source = "tektutor/docker"
     }
     localfile = {
       source = "tektutor/file"
     }
  }
}

resource "docker_image" "nginx" {
  image_name = "bitnami/nginx:latest"  
}

resource "docker_container" "containers" {
   count = 5
   container_name = "container_${count.index+1}"
   host_name = "container_${count.index+1}"
   image_name = "tektutor/ubuntu-ansible-node:latest"
}

resource "localfile" "myfile" {
  file_name = "./sample.txt"
  file_content = "Sample file content"

  provisioner "local-exec" {
     command = "ansible-playbook install-nginx-playbook.yml"
  }
}
