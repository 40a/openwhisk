# Setting up OpenWhisk on Ubuntu server

The following are verified to work on Ubuntu 14.04.3 LTS. You may need `sudo` or root access to install required software depending on your system setup.

  ```
  # Install git if it is not installed
  apt-get install git

  # Clone openwhisk
  git clone https://github.com/openwhisk/openwhisk.git
  
  # Change current directory to openwhisk
  cd openwhisk

  # Install all required software
  (cd tools/ubuntu-setup && source all.sh)
  ```

### Select one type of data store when creating vm
Follow instructions [tools/db/README.md](../db/README.md) on how to configure a data store for OpenWhisk.

## Build

  ```
  cd <home_openwhisk>
  ./gradlew distDocker
  ```

## Deploy

Follow the instructions in [ansible/README.md](../../ansible/README.md) to deploy and teardown.

Once deployed, several Docker containers will be running in your linux machine.
You can check that containers are running by using the docker cli with the command  `docker ps`.

### Configure the CLI
Follow instructions in [Configure CLI](../../docs/README.md#setting-up-the-openwhisk-cli). The API host
should be `172.17.0.1` or more formally, the IP of the `edge` host from the
[ansible environment file](../../ansible/environments/local/hosts).

### Use the wsk CLI
```
bin/wsk action invoke /whisk.system/samples/echo -p message hello --blocking --result
{
    "message": "hello"
}
```

