#!/bin/bash
# get latest update for OS
ensureLatest() {

	sudo apt-get update  # To get the latest package lists
    echo "...done"
	return 1
}

installGit() {
	sudo apt-get install git-core

	echo "git config --global user.name \"Michael Lawrence\"" >>~/.gitconfig
	echo "git config --global user.email mlawrence@aurlaw.com" >>~/.gitconfig

}

installGoLang() {

	cd ~ 
	wget https://storage.googleapis.com/golang/go1.3.1.linux-amd64.tar.gz
	tar -zxvf go1.3.1.linux-amd64.tar.gz
	rm go1.3.1.linux-amd64.tar.gz
	# binary distributions assume they will be installed at /usr/local/go Otherwise, you must set the GOROOT environment variable
	sudo mv ~/go /usr/local
	echo "export PATH=$PATH:/usr/local/go/bin" >> /etc/profile
	source /etc/profile
	echo 'updated profile'

	# make goproj directory for cloning go projects into
	sudo mkdir /tmp/goproj

	echo "export GOPATH=/tmp/goproj" >> /etc/profile
	source /etc/profile
	echo 'updated GOPATH'


    echo "Golang done"
	return 1


}

# Provision commands
provision() {
	
   ensureLatest

   installGit

   installGoLang

    return 0
}

provision