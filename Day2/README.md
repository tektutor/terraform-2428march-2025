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
cd Day2/ansible/ansible-custom-rolkubectl get pods -n ansible-awx
e
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
kubectl get nodes
```

Expected output
![image](https://github.com/user-attachments/assets/67770ac8-263b-4444-a9f3-3fef880f3efc)

Accessing Ansible Tower Dashboard from chrome web browser on the Ubuntu lab machine
```
minikube service awx-demo-service --url -n ansible-awx
```
You can launch the AWS webconsole using the url shown in the above command
```
http://192.168.49.2:31225
```
![image](https://github.com/user-attachments/assets/ca51256e-8f4d-449f-9d35-7929f5dd9e50)

Ansible Tower Login Credentials , save the login credentials in a text file to avoid typing the below command each time to get password
<pre>
username - admin   
password - 
</pre>

To get the password, you need to type the below command
```
kubectl get secret awx-demo-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```
![image](https://github.com/user-attachments/assets/2cfcf13a-f4d5-49d9-8a36-ff1f261e73c7)
![image](https://github.com/user-attachments/assets/7b0c56d5-b3ba-4b7e-87b5-0a8cc208c2c6)
![image](https://github.com/user-attachments/assets/05362dfd-4f54-4651-bd27-53962ae3a3e6)

## Troubleshooting - In case your ansible tower is not working, you will have to reinstall AWX
```
minikube delete
minikube start
```

Check if minikube is installed properly
```
kubectl get nodes
```

Expected output

Let's install AWX operator in minikube
```
cd ~
cd awx-operator/
git checkout 2.19.0
export NAMESPACE=ansible-awx
make deploy

kubectl get pods -n ansible-awx
```

Expected output
![image](https://github.com/user-attachments/assets/df848ac8-be68-4e5e-8f43-092fffcd5d4f)

Track the progress of awx installation
```
kubectl logs -f deployments/awx-operator-controller-manager -c awx-manager -n ansible-awx
```
Expected output


Access the AWX dashboard
```
kubectl create -f awx-ubuntu.yml -n ansible-awx
kubectl get pods -n ansible-awx
kubectl get svc -n ansible-awx
minikube service awx-demo-service --url -n ansible-awx
```
Expected output
AWX Login Credentials
```
username - admin
```

To retrieve password
```
kubectl get secret awx-demo-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```

If everything went smooth, you are expected to see similar page

To access the awx dashboard from other machines, you need to do port-forwarding
```
kubectl port-forward service/awx-ubuntu-service -n ansible-awx --address 0.0.0.0 10445:80 &> /dev/null &
```

You may now access the dashboard from other machines as
```
http://10.0.1.72:10445
```

## Lab - Installing SSH Server on our lab machine
```
sudo apt update && sudo apt install -y net-tools openssh-server
```
Expected output
![image](https://github.com/user-attachments/assets/791a0168-4695-4c82-b0f3-21f699dd499a)

## Lab - Creating an inventory in Ansible AWX
Navigate to your Ansible AWX Dashboard
![image](https://github.com/user-attachments/assets/6cc9489c-66f8-4737-a4c5-1ee5bccb5926)

Under Resources menu, select "Inventories"
![image](https://github.com/user-attachments/assets/a761021a-b527-4f81-89ce-1ba0b4932a6c)
Click "Add"
![image](https://github.com/user-attachments/assets/70b1469c-453d-46cb-8542-f1804c7c9121)
Select "Add inventory"
![image](https://github.com/user-attachments/assets/e38e321c-9c38-4b8c-a486-670ff38e98da)
![image](https://github.com/user-attachments/assets/e517a213-0bcf-4112-a386-373d81db471a)
Scroll down and click "Save" button
![image](https://github.com/user-attachments/assets/1120060e-ec9e-4525-a28a-ee12ade05c3b)
![image](https://github.com/user-attachments/assets/123b4030-0ee8-4cd1-848e-9c2f2ae2b041)

Click on "Hosts" within the MyInventory
![image](https://github.com/user-attachments/assets/e7037654-29cd-40af-9aa7-8c1f5ae82db5)
Click "Add"
![image](https://github.com/user-attachments/assets/bd8dcbc5-ed0c-41aa-92d1-95c7f2a30b76)
![image](https://github.com/user-attachments/assets/c937c0f0-3c55-4d22-ad46-e63ec8973201)
Click "Save"
![image](https://github.com/user-attachments/assets/06bf0dcb-2b11-4de8-a323-c420795109e7)
![image](https://github.com/user-attachments/assets/299c7d6b-ecff-4b88-8e93-a7266939b723)

Repeat this to add Ubuntu2, Rocky1 and Rocky2 anisble nodes
![image](https://github.com/user-attachments/assets/052f90c8-0850-43bd-9b1a-c45373c3bada)
Click "Save"
![image](https://github.com/user-attachments/assets/93674ebb-c329-4abc-b66d-315c0b98a573)
![image](https://github.com/user-attachments/assets/9c064212-c377-417c-bbdb-2ee0e656c9ab)
Click "Save"
![image](https://github.com/user-attachments/assets/23db861e-f314-4d40-8654-aae065657edd)
![image](https://github.com/user-attachments/assets/11ad4b03-02eb-42e9-a9f8-6000b308ceef)
![image](https://github.com/user-attachments/assets/3a70080e-c2a5-424d-bab4-2b3a29c2ea50)
Click "Save"
![image](https://github.com/user-attachments/assets/3a569d8d-efd7-480e-bd09-2d64f98a07a7)

ansible_host: localhost must be replaced with your lab machine IP address as shown below
![image](https://github.com/user-attachments/assets/821b3aec-2fb7-4b34-9930-89399b1aa7e6)
![image](https://github.com/user-attachments/assets/d11879ad-add4-41cb-8eda-1e6b17040e61)
![image](https://github.com/user-attachments/assets/aa0cfcff-7f7f-409e-b95c-fb4dd503a35d)
![image](https://github.com/user-attachments/assets/6e75177b-ee50-4f90-bdb8-06bbc8019317)


## Lab - Creating Credentials in Ansible AWX to provide private key details 
Navigate to your Ansible AWX Dashboard
![image](https://github.com/user-attachments/assets/62fba8a8-0fde-44b0-9188-4ae3efddbb43)

Under Resource menu on the left side, click on "Credentials"
![image](https://github.com/user-attachments/assets/463e422e-920d-4ffb-8216-232a346063cc)
Click "Add"
![image](https://github.com/user-attachments/assets/1f68e537-2b59-4d76-b321-7fe4325383df)
![image](https://github.com/user-attachments/assets/8248ea32-3121-4591-be87-6e93fb589061)
Under "Credential Type" search for Machine
![image](https://github.com/user-attachments/assets/88f0ff5f-95a8-4fbf-a469-119decef5468)
From terminal, cat your private key file
```
cat ~/.ssh/id_ed25519
```
![image](https://github.com/user-attachments/assets/542515d8-0b71-4c56-88d5-07ab572d2c56)
Copy as shown below
![image](https://github.com/user-attachments/assets/6ada258f-7a9c-4477-98bd-dea92d53ede4)
Under the SSH Private Key, you need to paste the private key you copied from your terminal
![image](https://github.com/user-attachments/assets/3b48b6b7-0a04-42f5-8317-41258f1412eb)
![image](https://github.com/user-attachments/assets/832dd2e8-e215-4b97-a4f3-9bef44e33bb7)
Scroll down
![image](https://github.com/user-attachments/assets/7a42b84f-96f6-4e87-91aa-eb715e8a91b3)
Click "Save"
![image](https://github.com/user-attachments/assets/cb36db0d-9156-443b-af44-40b8568eb7d4)


## Lab - Creating Project in Ansible AWX
Navigate to your Ansible AWX Dashboard
![image](https://github.com/user-attachments/assets/5e969802-c70b-4543-9550-552aee07fa26)

Under the Resources menu on the left side, click "Projects"
![image](https://github.com/user-attachments/assets/e9ec3f7f-8ea1-4cb9-a6b7-87ba9d3beba3)
Click "Add"
![image](https://github.com/user-attachments/assets/5d483efb-cb05-4841-8f9f-599287b6a789)
![image](https://github.com/user-attachments/assets/ac2de99f-6b51-4139-8129-136facb7a746)
Select "Git" under "Source Code Type"
![image](https://github.com/user-attachments/assets/faa0e713-22a6-46d0-bba3-d2026e09fea6)
Under Source Code URL type "https://github.com/tektutor/terraform-2428march-2025.git"
![image](https://github.com/user-attachments/assets/c102150f-7c30-4e1c-a059-e2e05b0f6e3e)
Under the Options, select the checkbox "Update Revision on Launch"
![image](https://github.com/user-attachments/assets/ea958775-f594-4e3c-a09d-4a7b04042218)
Scroll down and click "Save"
![image](https://github.com/user-attachments/assets/1f21ba6f-cb68-4ce6-b0cc-1729ce888030)
![image](https://github.com/user-attachments/assets/1d1f31f5-e4a9-4d39-8f3e-bd268cbea06b)

Progress
![image](https://github.com/user-attachments/assets/28562f02-8943-4657-af58-3ea8ab99fb07)
![image](https://github.com/user-attachments/assets/b2da1ad2-ff47-41cc-bf1d-df093cceba8b)
![image](https://github.com/user-attachments/assets/60c51de6-626d-460b-a8d3-175e72d54852)
![image](https://github.com/user-attachments/assets/861619bc-66f3-4d7f-b95d-899685b0c317)
![image](https://github.com/user-attachments/assets/24cf2363-3966-45e1-a3c7-274368b476d6)

## Lab - Creating a Job Template to run Ansible Playbook in Ansible Automation Platform
Navigate to your Ansible AWX Dashboard
![image](https://github.com/user-attachments/assets/66e18ae8-f133-4db4-8d66-b210759f348a)

Under Resources menu in the left side, select "Templates"
![image](https://github.com/user-attachments/assets/be2ef159-cd72-466f-8ac8-183da9e02596)
Click "Add"
![image](https://github.com/user-attachments/assets/69dd5b1b-dc37-438d-bfeb-e5c86ac5e90f)
Select "Add Job Template"
![image](https://github.com/user-attachments/assets/f150c744-feed-4af4-9988-03cf84627337)
Click the Search option under Inventory and Select "MyInventory" and click "Select" button
![image](https://github.com/user-attachments/assets/ba737779-9413-4ff0-844a-57705d5b658e)
Click the Search option under Project and Select "Install Nginx Playbook"
![image](https://github.com/user-attachments/assets/e72df0ab-18ed-4368-8475-8b1dd77cc692)
Click "Select" button
![image](https://github.com/user-attachments/assets/ae413b5a-94b1-4593-b426-da0f3ca7531d)
Click "Search" button under Credentials and Select "PrivateKey"
![image](https://github.com/user-attachments/assets/4dfd5fdb-4717-4fad-b626-c1ea54d70aa5)
Click "Select" button
![image](https://github.com/user-attachments/assets/2f4e41eb-69b8-4565-ac3e-9493dd010d36)

Under Playbook, select "Day2/ansible/ansible-custom-role/install-nginx-playbook"
![image](https://github.com/user-attachments/assets/d035085c-497d-44ed-990b-b6d1d2a7fd18)
![image](https://github.com/user-attachments/assets/838e61a6-a077-4c62-a559-278d773d30a1)
Scroll down
![image](https://github.com/user-attachments/assets/f4d25f66-9904-4cb1-8836-48eb68b31745)
Click "Save"
![image](https://github.com/user-attachments/assets/9db0a550-a4ae-4ae1-ab76-efad72a27038)

Click "Launch" button to run the playbook
![image](https://github.com/user-attachments/assets/1241d7d0-85ff-45e2-880d-649ca8ea6e11)
![image](https://github.com/user-attachments/assets/e04a5760-95c5-4eec-8b74-63a5282790a1)
![image](https://github.com/user-attachments/assets/cb1ec154-5793-4de8-905e-d96d66f2c88c)
![image](https://github.com/user-attachments/assets/c9d6e740-e86f-49d1-af5f-3edea6c78010)
![image](https://github.com/user-attachments/assets/907037ec-30ce-4b1d-84c7-4c98d3ea1a38)
![image](https://github.com/user-attachments/assets/4591687c-2588-4d97-a78c-9e1b4fd7c7d4)
![image](https://github.com/user-attachments/assets/a691273a-ae5c-4647-9dba-28d061066dc9)
![image](https://github.com/user-attachments/assets/c4ff6e0d-1412-4255-a999-9237f5b60c72)
![image](https://github.com/user-attachments/assets/d2961f28-6211-4af8-bbc9-dbb7377c8de4)

## Lab - Creating a non-non admin user in AWX
Navigate to AWX Dashboard
![image](https://github.com/user-attachments/assets/989f0e73-ec6c-4db0-a377-ffc418e7d871)

Under Access menu in the left side, select "Users"
![image](https://github.com/user-attachments/assets/4966273e-b13e-4c97-b454-acb3c72ee2dc)
Click "Add"
![image](https://github.com/user-attachments/assets/763cbbaa-2552-42fc-8219-dece80640343)
![image](https://github.com/user-attachments/assets/9391e91a-c015-404d-a5d2-e2d365d04079)
Click "Save"
![image](https://github.com/user-attachments/assets/ec7c7125-15e3-4d5d-be39-0dc32351d0b5)
Logout from AWX

Login as the new user you created
![image](https://github.com/user-attachments/assets/f564e272-23cd-4531-95a0-3c94e41d6f28)
![image](https://github.com/user-attachments/assets/93915477-5bfb-41fa-952e-36a79b471bc7)
![image](https://github.com/user-attachments/assets/88a4e543-3139-43ea-b926-d1cb33dda9b4)
![image](https://github.com/user-attachments/assets/f7756cee-b1b2-46f3-8e60-86212c8756b7)
![image](https://github.com/user-attachments/assets/86227194-fc51-40c6-ad0d-8f3bc88a3228)
![image](https://github.com/user-attachments/assets/605bacdc-06b2-43ac-aa56-c2638a066f71)

Logout from new user and relogin as admin user to provide access to the new user you created.
![image](https://github.com/user-attachments/assets/306e4384-18a0-42db-9de7-0a74495e2f61)
![image](https://github.com/user-attachments/assets/6e8677b6-e71a-4146-adaf-192ac01325dd)
Click "Projects"
![image](https://github.com/user-attachments/assets/e22f2be1-bf9f-405b-862b-ef2ec20f4139)
Select "Install Nginx Playbook"
![image](https://github.com/user-attachments/assets/aafe47d4-fe74-48be-a6cf-78df9503e606)
Switch to "Access" tab( I'm not referring the menu on left side )
![image](https://github.com/user-attachments/assets/c36b524b-3c4c-43e5-9134-31ef7c428b6c)
Click "Add"
![image](https://github.com/user-attachments/assets/aea129f4-0af9-4bde-bdc8-28fa92f966af)
Select "Users"
![image](https://github.com/user-attachments/assets/7a3b9e7d-594e-4b0c-bc0a-d920fa7e9d7a)
Click "Next"
![image](https://github.com/user-attachments/assets/d839e9c2-17a6-485f-82f0-9fe883363249)
Select your new user ( in my case 'jegan' )
![image](https://github.com/user-attachments/assets/7246bf89-494f-4fca-a4cf-51d270528b61)
Click "Next" button
![image](https://github.com/user-attachments/assets/3e5e5156-d8f5-4812-96ab-9ea9fc68fe5a)
Select the checkbox "Read" and "Use"
![image](https://github.com/user-attachments/assets/57330b5f-6918-4c6d-801a-31fc14d5caf6)
Click "Save"
![image](https://github.com/user-attachments/assets/0d390a01-1abd-4468-9b24-00a53aa76e3c)
![image](https://github.com/user-attachments/assets/ecc8d4db-efb6-4614-958e-ee4d46e19d35)
