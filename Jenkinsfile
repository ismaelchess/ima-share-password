pipeline {
    agent any

    environment {
        XDG_CACHE_HOME="/tmp/.cache"
        GOPATH="$HOME/go"
        GOBIN="$GOPATH/bin"
        PATH="$PATH:$GOPATH/bin:/usr/local/go/bin"
    }
    stages {
        stage('unit tests') {
            steps {
                sh '''
                    make test
                '''
            }
        }
        stage('golang-lint'){
            steps{
                sh '''
                    make lint
                '''
            }
        }
    }
}