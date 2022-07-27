echo yourpassword | sudo -S docker stop passwd
sudo docker container rm passwd
sudo docker image rm kiririx/passwd:latest
sudo docker pull kiririx/passwd:latest
sudo docker run -p 10011:8080 -d --name passwd -e db=mysql -e db_username=root -e db_password=xxx -e db_host=xxx -e db_port=3306 -e db_database=passwd kiririx/passwd:latest