pipeline {

	agent {label 'master'}

	environment{
		WORKSPACE = '/home/jenkins'
	}

	stages {
		stage ('Install Golang') {
			steps {
				sh 'sudo rm -rf /usr/local/go'
 				sh "printenv | sort"
				sh 'wget https://dl.google.com/go/go1.12.4.linux-amd64.tar.gz'
				sh 'tar -xvf go1.12.4.linux-amd64.tar.gz'
				sh 'sudo mv go /usr/local'
				sh 'echo "GOROOT=/usr/local/go" >> ~/.profile'
				sh 'echo "GOPATH=/var/lib/jenkins/go" >> ~/.profile'
				sh 'echo "PATH=$GOPATH/bin:$GOROOT/bin:$PATH" >> ~/.profile'
				sh '. ~/.profile'
				sh 'echo $GOPATH'
				sh '/usr/local/go/bin/go version'
				sh '/usr/local/go/bin/go get -v github.com/YaleOpenLab/openx'
				sh '/usr/local/go/bin/go get -v github.com/YaleOpenLab/opensolar'
				sh 'cd ~/go/src/github.com/YaleOpenLab/openx'
				sh '/usr/local/go/bin/go build -v ./...'
			}
		}
	}
}
