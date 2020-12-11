#Commands

sudo docker-compose up -d
sudo docker run --name NAME --network net -e VIRTUAL_HOST="DOMAIN_NAME" -e LETSENCRYPT_HOST="DOMAIN_NAME" -d DOCKERIMAGE:TAG
