# Day 2

## Lab - Managing Ubuntu and Rocky ansible node containers with a single ansible playbook
```
cd ~/terraform-2428march-2025
git pull
cd Day2/ansible/nginx
ansible-playbook install-nginx-playbook.yml

curl http://localhost:8001
curl http://localhost:8002
curl http://localhost:8003
curl http://localhost:8004
```

Expected output
![image](https://github.com/user-attachments/assets/e463a901-1d9f-4e92-8aed-0e83656e8c9a)
![image](https://github.com/user-attachments/assets/b077e3f1-cf8b-41c6-93bf-66f4cb1dc08f)
![image](https://github.com/user-attachments/assets/ac73d86b-78f4-48f7-8653-4ed11e1c8927)
![image](https://github.com/user-attachments/assets/318b3867-4ab9-4865-86ee-ebc289e9b803)
![image](https://github.com/user-attachments/assets/6c06f802-717d-47da-befb-86519bee3fc5)
![image](https://github.com/user-attachments/assets/3ee36eb0-f896-4572-8da0-88b9c1ee7d09)
![image](https://github.com/user-attachments/assets/74b30240-d3f1-488d-82b1-756f0395a424)

## Info - SOLID Design Principles
<pre>
S - Single Responsibility Principle (SRP)
O - Open Closed Principle (OCP)
    - a good design should be open for extension without modifying existing code
    - we can add new code in new files
L - Liskov Substitution Principle ( LSP)
I - Interface Seggregation Principle 
D - Dependency Injection or Dependency Inversion or Inversion of Control(IOC)
</pre>

## Lab - Running the refactored install nginx playbook 
```
cd ~/terraform-2428march-2025
git pull
cd Day2/ansible/nginx/after-refactoring
ansible-playbook install-nginx-playbook.yml

curl http://localhost:8001
curl http://localhost:8002
curl http://localhost:8003
curl http://localhost:8004
```

Expected output
![image](https://github.com/user-attachments/assets/967b81d0-0bdb-4525-bfec-7fc870cd2878)
![image](https://github.com/user-attachments/assets/a14f1736-2115-44ff-b5bf-c1edbab367b8)
![image](https://github.com/user-attachments/assets/2b478808-c145-4915-8647-dc90e555dd16)
![image](https://github.com/user-attachments/assets/238ee9dc-524c-4098-9c19-382a34946bda)
![image](https://github.com/user-attachments/assets/aac589d1-839e-4300-b694-2da3a7b8cbb0)
![image](https://github.com/user-attachments/assets/0cab7886-a442-4950-82f6-f7f8a7dc32e1)
![image](https://github.com/user-attachments/assets/c2f31b7d-098f-4217-93aa-5ca1f53e6ddd)
![image](https://github.com/user-attachments/assets/c2d4dc4e-ef6e-4c21-b3a8-a4ea587a9d37)
![image](https://github.com/user-attachments/assets/3aaf4f70-ec25-4802-a809-3ed908bbfe4b)
![image](https://github.com/user-attachments/assets/ec02f55a-64b3-495a-81e7-00116405d60c)

