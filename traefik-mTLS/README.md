# Traefik mTLS setup


## Prerequisites
- some knowledge of Traefik & Smallstep (certs)
- knowledge of TLS & mTLS
- access to hosts file for creating DNS entries

## Setting up certificates

```bash

## inside the hosts file, locate for the 127.0.0.1 entry, append DNS separated by spaces.

- install step cli: `brew install step`

127.0.0.1 localhost mo.test server.test ca.test

The localhost is very important, standard, default entry.
ca.test = DNS for the CA, the certificate authority.
mo.test = DNS for the traefik dashboard.
mo.ab = an application running on local to access a POC API.


# Initialise the step ca
step ca init --profile=test.ca --context=test.ca 

? What deployment type would you like to configure?: 
Use the arrow keys to navigate: ↓ ↑ → ← 
Use the arrow keys to navigate: ↓ ↑ → ← 
Use the arrow keys to navigate: ↓ ↑ → ← 
Use the arrow keys to navigate: ↓ ↑ → ← 
✔ Deployment Type: Standalone
What would you like to name your new PKI?
✔ (e.g. Smallstep): Test
What DNS names or IP addresses will clients use to reach your CA?
✔ (e.g. ca.example.com[,10.1.2.3,etc.]): localhost,ca.test
What IP and port will your new CA bind to? (:443 will bind to 0.0.0.0:443)
✔ (e.g. :443 or 127.0.0.1:443): :54321
What would you like to name the CA's first provisioner?
✔ (e.g. you@smallstep.com): mo@example.com
Choose a password for your CA keys and first provisioner.
✔ Password: <...>
✔ Password: <...>
Generating root certificate... done!
Generating intermediate certificate... done!


####


## Add a provisioner for ACME protocol (required by Traefik)

step ca provisioner add acme --type ACME --x509-min-dur 1h --x509-default-dur 9490h1m0s --x509-max-dur 9490h1m0s

## Install root certificate on the local machine.

step certificate install /Users/mohameda/.step/authorities/test.ca/certs/root_ca.crt

## add system password ^^^

## Start smallstep CA server

step-ca --context=test.ca


https://mo.test

```

## Start Traefik & test

```bash

chmod 600 acme.json

traefik --configfile=static.yml

http://localhost:8080/dashboard/#/

https://mo.test

```


## Running Traefik method (works)

### Local (config file)

Direct Traefik:
- Install Traefik: `brew install traefik`
- `sudo rm -rf /etc/traefik/conf/dynamic.yml`

- `sudo ln -s $(pwd)/dynamic.yml /etc/traefik/conf/dynamic.yml`
- `traefik --configfile=./static.yml`
- `http://localhost:8080/dashboard`

via Containers:
- `docker build . -t traefik`
- `docker build . -t traefik --build-arg ARCH=linux_arm64` // for M1 Macs
- `docker run -p 8080:8080 -p 80:80 -v ./static.yml:/etc/traefik/traefik.yml -it traefik`
- `http://localhost:8080/dashboard`


## mTLS setup

```bash

## Creating certificate for mTLS connection
- step ca certificate client client.crt client.key --set emailAddresses=example@test.in --context=test.ca --size=4096 --kty=RSA

## if you run the above, make sure the CA server is running!! (step-ca --context=test.ca)

- use the provisioner created previously to create a certificate for the client. (Provisioner: mo@example.com (JWK))

step certificate p12 client.p12 client.crt client.key


```
