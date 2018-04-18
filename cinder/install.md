### How to install cinder on bare metal.


### standalone install

#### container

step 0：

git clone https://github.com/openstack/cinder

cd cinder/contrib/block-box && make

step 1:

sudo curl -L https://github.com/docker/compose/releases/download/1.21.0/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose

docker-compose up

#### binary
