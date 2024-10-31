# Backend em Golang
Repositório padrão para projetos Backend usando Golang.


#### 📁 Estrutura

- [ ] Pastas
- [ ] Docker
- [ ] Logs
- [ ] Versionamento

#### 🛢️ Banco de Dados

- [ ] MongoDB
- [ ] Migrations
- [ ] Seeders

#### 🔐 Autenticação

- [ ] JWT
- [ ] SSO Google

#### 📧 SMTP

- [ ] Configuração
- [ ] Recuperar senha

#### 🛠️ CRUD

- [ ] Perfil
- [ ] Usuários

#### 🛒 Compras

- [ ] Stripe

#### ☁️ AWS

- [ ] S3
- [ ] RDS
- [ ] SQS

## Variáveis

.env
```shell
#Configurações para o Docker (local)
DB_URL=mongodb://mongo:27017
DB_NAME=backend-golang
DB_USERNAME=root
DB_PASSWORD=root

#SMTP
SMTP_HOST=
SMTP_PORT=465
SMTP_USERNAME=contato@overall.cloud
SMTP_PASSWORD=
```


## Ambiente de Desenvolvimento

### Configurando Docker

Pré configuração:

```shell
go mod init backend-golang
go mod tidy
```

Criando a imagem:

```shell
docker-compose up
```

Para verificar (listagem containers):

```shell
docker ps
```

Para remover:

```shell
docker-compose down
```

Para logs:

```shell
docker logs -f backend-golang mongo mongo-express
```

## Ambiente de Produção
    
 ### Instalanndo e Configurando no Servidor

Instalar o Go no Ubuntu:

 ```shell
sudo apt install golang-go git
 ```

### Clone o repositório

```shell
git clone git@github.com:gustavokennedy/backend-golang.git
cd backend-golang
```

### Instalando Dependências

```shell
go build main.go
```

<a href="https://www.digitalocean.com/community/tutorials/how-to-install-nginx-on-ubuntu-22-04" target="_Blank">Instalar o Nginx no Ubuntu.</a>

<a href="https://www.digitalocean.com/community/tutorials/how-to-secure-nginx-with-let-s-encrypt-on-ubuntu-22-04" target="_Blank">Instalar SSL com Nginx no Ubuntu.</a>

Primeiro, crie um novo arquivo no /lib/systemd/system/ chamado backend-golang.service:

 ```shell
 sudo nano /lib/systemd/system/backend-golang.service
 ```
 
 ```shell
[Unit]
Description=backend-golang

[Service]
Type=simple
Restart=always
RestartSec=5s
ExecStart=/home/ubuntu/go/backend-golang/main

[Install]
WantedBy=multi-user.target
```

Agora que você escreveu o arquivo da unidade de serviço, inicie seu serviço da web Go executando:

```shell
 sudo service backend-golang start
 ```

Para confirmar se o serviço está em execução, use o seguinte comando:

```shell
 sudo service backend-golang status
 ```

Para verificar no Log no Service:

  ```shell
 sudo journalctl -u backend-golang -b
 ```

 Para reiniciar configurações de Service:

  ```shell
 sudo systemctl daemon-reload
 ```

 ### Configurando Nginx

 Primeiro, altere seu diretório de trabalho para o sites-enabled do Nginx:

```shell
sudo nano /etc/nginx/sites-enabled/default
 ```

Adicione as seguintes linhas ao arquivo para estabelecer as configurações:

```shell
server {
    server_name _;

    location / {
        proxy_pass http://localhost:8080;
    }
}
 ```

Em seguida, recarregue suas configurações do Nginx executando o comando reload:

```shell
sudo nginx -s reload
 ```
