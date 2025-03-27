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
