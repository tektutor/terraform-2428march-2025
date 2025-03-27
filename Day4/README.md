# Day 4

## Lab - Using remote-exec provisioner block in Terraform to run commands on remote machine
```
cd ~/terraform-2428march-2025
git pull
cd Day4/terraform/remote-exec
terraform init
terraform apply --auto-approve
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/7204143d-b1b3-4720-bbbe-6a2af74402cb)
![image](https://github.com/user-attachments/assets/634da5d7-a096-4ca5-bdf5-460808dc6e11)
![image](https://github.com/user-attachments/assets/7868b483-8d5b-466a-9a43-0c5b75c4bb82)
![image](https://github.com/user-attachments/assets/4fd5973e-443a-4abb-a3e5-791f1cfe4af3)
![image](https://github.com/user-attachments/assets/265450eb-0395-41ff-bb3d-888d4cbb43f1)
![image](https://github.com/user-attachments/assets/9f1784d2-217f-4dcd-b40e-5506d05d13ea)

Once you are done with this exercise, let's dispose the resource we provisioned via terraform
```
terraform destroy --auto-approve
```
![image](https://github.com/user-attachments/assets/89d3b63f-4ea0-43f5-84a3-6623015e9a5a)
![image](https://github.com/user-attachments/assets/a4dd1fe8-08a9-4906-aa01-80bc88d4793c)
![image](https://github.com/user-attachments/assets/08a995a4-5227-43e5-acea-ce73c01a20bd)
![image](https://github.com/user-attachments/assets/b417bc76-6650-4e08-813d-ecce2eaf5412)

## Lab - Terraform module
```
cd ~/terraform-2428march-2025
git pull
cd Day4/terraform/module
terraform init
terraform apply --auto-approve
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/bf1529a9-65fd-4e47-8c80-b9ba9473f157)
![image](https://github.com/user-attachments/assets/1491d8ad-e053-4923-a26b-85a00a04d26e)
![image](https://github.com/user-attachments/assets/66578f99-5597-4e80-b7e4-e5970f7f3496)
![image](https://github.com/user-attachments/assets/fce01f59-a0be-4976-aeb3-b792e730be92)
![image](https://github.com/user-attachments/assets/1b6e1fb3-e980-41c7-9dd7-08ea9a258a72)

Once you are done with this exercise, you may dispose the resources provisioned by terraform
```
terraform destroy --auto-approve
```
![image](https://github.com/user-attachments/assets/cbd2e341-309f-45c6-81f8-39e41d37c3cb)
![image](https://github.com/user-attachments/assets/af11fe35-ace5-41e5-b37c-3448038f9e91)

## Lab - Deleting a specific resource while retaining other resources
```
cd ~/terraform-2428march-2025
git pull
cd Day4/terraform/module

terraform init
terraform apply --auto-approve
docker ps

terraform destroy --target module.create-docker-containers.docker_container.my_container[4]

terraform destroy --target module.create-docker-containers.docker_container.my_container[3] --target module.create-docker-containers.docker_container.my_container[2]

docker ps

terraform destroy --target module.create-docker-containers.docker_container.my_container[0] --target module.create-docker-containers.docker_container.my_container[1] 

docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/480990c0-6c80-4241-a1ed-55832968af66)
![image](https://github.com/user-attachments/assets/558e1f6c-035d-470d-9e1f-d11e7b5df042)
![image](https://github.com/user-attachments/assets/8a26a02c-78e0-43d6-8bb7-5a59aa5b4207)
![image](https://github.com/user-attachments/assets/51957767-ac3b-43d1-aca6-02c75197a48c)
![image](https://github.com/user-attachments/assets/b18756aa-fafe-4b4f-a467-1dd8ad89cf6b)
![image](https://github.com/user-attachments/assets/037f431f-f3dd-423f-aca1-11190e0a40a8)
![image](https://github.com/user-attachments/assets/954d769a-2d8f-4d74-b9df-2938d19e81ec)
![image](https://github.com/user-attachments/assets/4b8c63fb-3c77-49bd-adc0-944f0cc59fe1)
![image](https://github.com/user-attachments/assets/918e28fb-f184-4a71-9c03-9396be6d0a95)

## Lab - Develop a custom terraform file provider in golang using Terraform Provider plugin SDK

You need to create a folder 
```
mkdir -p /home/rps/go/bin
touch ~/.terraformrc
```

Create a file named .terraformrc under your home directory i.e /home/rps/.terraformrc
<pre>
provider_installation {
  dev_overrides {
     "registry.terraform.io/tektutor/docker" = "/home/rps/go/bin",
     "registry.terraform.io/tektutor/file" = "/home/rps/go/bin",
  }
  direct {}
}  
</pre>

Then you may proceed with the below
```
cd ~/terraform-2428march-2025
git pull
cd Day4/terraform/custom-terraform-providers/file
go mod init github.com/tektutor/terraform-provider-file
go mod tidy
ls -l
go build
ls -l
go install
ls -l /home/rps/go/bin
```

Expected output
![image](https://github.com/user-attachments/assets/8717c29c-b134-43ac-8895-319f0adffa4f)
![image](https://github.com/user-attachments/assets/ee8e2d3a-0ce6-4aff-b15c-9bf362ecac87)
![image](https://github.com/user-attachments/assets/b4dc8086-7e2f-4ccb-92e7-b5b4b6d94c86)

## Lab - Let's use our custom terraform file provider in Terraform HCL manifest script
```
cd ~/terraform-2428march-2025
git pull
cd Day4/terraform/custom-terraform-providers/test-file-provider
cat main.tf
terraform plan
ls
terraform apply --auto-approve
ls
cat myfile.txt
cat terraform.tfstate
```

Expected ouptut
![image](https://github.com/user-attachments/assets/346b2468-32e9-40c9-b7d4-c201f3a12315)
![image](https://github.com/user-attachments/assets/a8cd4418-aa41-454f-a733-dae6b4d3a5ed)
![image](https://github.com/user-attachments/assets/edfe9a4f-b37a-4602-8082-9acf2763a010)
![image](https://github.com/user-attachments/assets/817c75dc-8385-4d87-a708-b04bf26897a8)
![image](https://github.com/user-attachments/assets/1da25021-3fd5-4e74-8fca-26b464019bcd)
![image](https://github.com/user-attachments/assets/1636f558-5055-412b-b0fa-5b0c00df4aba)
![image](https://github.com/user-attachments/assets/9f5c6983-72b3-4b43-be47-b4202464052b)
![image](https://github.com/user-attachments/assets/1714e521-14c5-4754-b28d-d1c8c6ce3554)


Let's see if the terraform destroy is supported by our file provider
```
ls
cat myfile.txt
terraform destroy --auto-approve
ls
```

Expected output
![image](https://github.com/user-attachments/assets/00a3de73-c7eb-48ee-b5fa-863d9220eec2)
![image](https://github.com/user-attachments/assets/1986e376-47a1-43c5-a62e-a068e7447af2)
![image](https://github.com/user-attachments/assets/46f3251b-2acb-4e36-9a94-3002bd9a26bd)

## Lab - Developing a Custom Terraform Docker Provider in Go language
```
cd ~/terraform-2428march-2025
git pull
cd Day4/terraform/custom-terraform-providers/docker
go mod init github.com/tektutor/terraform-provider-docker
go mod tidy
ls -l
go build
ls -l
go install
ls -l /home/rps/go/bin
```

Expected output
![image](https://github.com/user-attachments/assets/a7bc99b2-6f7f-41a1-8bd8-d2b5b116e183)
![image](https://github.com/user-attachments/assets/2ba9c16d-8f05-4063-93c9-6c8776f329a7)
![image](https://github.com/user-attachments/assets/c13cac60-162e-4677-9234-ccff27bba093)

## Lab - Using our custom terraform docker provider in Terraform scripts
```
cd ~/terraform-2428march-2025
git pull
cd Day4/terraform/custom-terraform-providers/test-docker-provider
cat main.tf
terraform plan
terraform apply --auto-approve
docker ps
```

Expected output
![image](https://github.com/user-attachments/assets/31eb035a-8a5e-4edc-8425-45a9e29cdf74)
![image](https://github.com/user-attachments/assets/e5bbb09f-8f00-4054-aebc-196e8b391318)
![image](https://github.com/user-attachments/assets/3c029425-1a47-430d-b249-1271896748c3)
![image](https://github.com/user-attachments/assets/f9667cc1-b098-4773-b4fe-dcf50ff07cb2)

