pipeline {
    agent any
    tools {
        go '1.20'
    }
    environment {
        GO111MODULE = 'on' //enable go module
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
        stage('Build') {
            steps {
                sh 'echo "Building go program"'
                sh 'go build'
                sh 'go run main.go'
            }
        }

        stage('Build docker-compose') {
            steps {
                sh 'docker-compose up --build'
            }
        }
    }
}