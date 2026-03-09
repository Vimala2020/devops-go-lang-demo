pipeline {
    agent any

    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-creds')
        IMAGE_NAME = 'vimala92/devops-pipeline-demo'
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/Vimala2020/devops-go-lang-demo.git'
            }
        }

        stage('Build Go App') {
            steps {
                sh 'go version'        // verify Go is installed
                sh 'go build -o main .' 
            }
        }

        stage('Test') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Docker Build & Push') {
            steps {
                sh 'docker version'    // verify Docker access
                sh 'docker build -t $IMAGE_NAME:v1 .'
                withDockerRegistry([credentialsId: 'dockerhub-creds', url: '']) {
                    sh 'docker push $IMAGE_NAME:v1'
                }
            }
        }

        stage('Deploy to Kubernetes') {
            steps {
                sh 'kubectl apply -f k8s/deployment.yaml'
                sh 'kubectl apply -f k8s/service.yaml'
            }
        }
    }
}