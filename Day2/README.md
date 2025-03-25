![image](https://github.com/user-attachments/assets/7716c29d-6667-4649-afb4-4f8cd910b42c)# Day 2

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

## Info - Ansible tools
<pre>
ansible - used to run ansible ad-hoc command ( invoking a single ansible module without writing an ansible playbook )
ansible-playbook - used to run ansible playbooks ( end to end automation that invokes multiple ansible modules )
ansible-doc - used to get help about any ansible module(s)
ansible-galaxy - used to download/install/develop custom ansible roles
ansible-vault - used to encrypt/decrypt sensitive data like login credentials, certs, etc.,
</pre>

## Info - Ansible Role
<pre>
- ansible roles helps write reusable code
- ansible roles are like Dynamic Link Library with reusable functions
- just like we can't run the DLL directory, we can't run the ansible roles directly
- just like DLL are loaded in application, after dynamically loading & linking functions from DLL application is able to invoke the functions defined in the DLL, same way ansible playbooks can invoke the ansible roles
- For what all purpose we could develop Ansible Role
  - let's say you wish to install nginx in multiple Linux distributions, you could develop a custom nginx role
  - let's say you wish to install weblogic in multiple OS ( Linux, Windows & Mac ) you could develop a custom weblogic role
  - let's say you wish to install Kubernetes in multiple OS ( Many Linux Distributions ), this could developed as a kubernetes role
- Ansible roles would look like an Ansible Playbook, but it can't be run directly, can only be invoked from Ansible Playbooks
- it follows specific recommended folder structures and many Ansible best practices
</pre>

## Demo - Developing a custom ansible role to install nginx in Rocky and Ubuntu Linux Distributions
```
cd ~/terraform-2428march-2025
git pull
cd Day2/ansible/ansible-custom-role
ansible-galaxy init nginx
tree nginx
```
![image](https://github.com/user-attachments/assets/5da024d5-bfbd-407c-84b1-551b90b8ac6d)
![image](https://github.com/user-attachments/assets/c494c820-7931-4fa8-b7fb-5fd541524231)

Note
<pre>
defaults - has read-only(static - which is not going change while playbook runs) user-defined variables 
vars - has regular variables 
files - copy ansible modules picks the files we mention under src attribute from this folder
handlers 
- based on notification, certain tasks can be executed on demand at run time
- in case, you wish to restart a service once you made some config changes the config task can notify the restart handler
tasks - will contain all the tasks we normally write in an ansible playbook
templates - template module picks the files we mention under src attribute from this folder
tests - contains a test inventory and test playbook to demonstrate how one could invoke ansible role ( we are going to delete this)
</pre>

Running the ansible playbook that invokes our custom nginx ansible role
```
cd ~/terraform-2428march-2025
git pull
cd Day2/ansible/ansible-custom-role
cat install-nginx-playbook.yml
ansible-playbook install-nginx-playbook.yml
```
Expected output
![image](https://github.com/user-attachments/assets/ab43eec8-85ac-4d53-810b-7cb19cb80ef3)
![image](https://github.com/user-attachments/assets/d91cd1b7-7c79-4dbb-973f-6298d403c22f)
![image](https://github.com/user-attachments/assets/7040a06a-5877-4ab6-ba93-e11619565c09)
![image](https://github.com/user-attachments/assets/006de241-26cc-4fd1-b90d-63986fc82691)
![image](https://github.com/user-attachments/assets/df24a62f-c71d-403c-bf3d-a21f7b5fb959)
![image](https://github.com/user-attachments/assets/2905a51e-55b8-4a83-b577-aedb23f925a4)
![image](https://github.com/user-attachments/assets/bf6ad431-2a0f-4f29-98c8-860b216a7244)

## Lab - Ansible Dynamic Inventory
```
cd ~/terraform-2428march-2025
git pull
cd Day2/ansible/dynamic-inventory
cat ansible.cfg
./inventory
ansible all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/d664c231-67e9-47f2-afe6-996bbe3eb28e)
![image](https://github.com/user-attachments/assets/b091c14f-8a63-4148-8dc2-293e1093541d)

## Lab - Securing sensitive data using ansible vault and using them in ansible playbook

When ansible-playbook prompts for password, type root as the password.

You need to create the ~/.my-vault-password as shown below
```
echo root > ~/.my-vault-password
```

Now you may proceed with the below commands
```
cd ~/terraform-2428march-2025
git pull
cd Day2/ansible/vault
ansible-vault create mysql-login-credentials.yml
cat mysql-login-credentials.yml

ansible-vault view mysql-login-credentials.yml
ansible-vault edit mysql-login-credentials.yml
ansible-vault decrypt mysql-login-credentials.yml
cat mysql-login-credentials.yml
ansible-vault encrypt mysql-login-credentials.yml
cat playbook.yml
ansible-playbook playbook.yml --ask-vault-pass
cat ansible.cfg
cat ~/.my-vault-password
ansible-playbook playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/45af7b6b-e06a-48b8-8879-3d40739d572a)
![image](https://github.com/user-attachments/assets/d39afc1a-87d7-4e03-9ede-c58c78ff37bc)
![image](https://github.com/user-attachments/assets/81647a19-297e-4bbb-afd1-a58fa1658af3)
![image](https://github.com/user-attachments/assets/916653c0-053e-4c0d-b00b-e859d0526964)
![image](https://github.com/user-attachments/assets/c0f4d957-5426-450e-85b7-8cc69c710ebb)
![image](https://github.com/user-attachments/assets/aed65dcf-c424-47d4-bd3b-65b9746adcc4)

## Info - Red Hat Ansible Automation Platform ( a.k.a Ansible Tower in the past )
<pre>
- this is an enterprise product developed by Red Hat on top of opensource AWX
- AWX is an opensource product that is developed on top of opensource Ansible core
- as AWX supports webconsole, Red Hat Ansibe Automation Platform also supports webconsole
- can invoke existing ansible playbooks
- supports RBAC(Role Based Access Control - User Management)
- Ansible AWX or Ansible Automation Platform can be installed on any Linux Distribution or within Kubernetes/Openshift
</pre>

## Lab - Launching Ansible Tower
```
minikube status
minikube start
minikube status
```

Expected output
![image](https://github.com/user-attachments/assets/67770ac8-263b-4444-a9f3-3fef880f3efc)

## Lab - Installing SSH Server on our lab machine
```
sudo apt update && sudo apt install -y net-tools openssh-server
```
Expected output
![image](https://github.com/user-attachments/assets/791a0168-4695-4c82-b0f3-21f699dd499a)
