pipeline {
    agent any

    environment {
        IMAGE_NAME = "vimala92/devops-pipeline-demo"
    }

    stages {

        stage('Checkout') {
            steps {
                git branch: 'main', url: 'https://github.com/Vimala2020/devops-go-lang-demo.git'
            }
        }

        stage('Build Go App') {
            agent {
                docker {
                    image 'golang:1.21'
                }
            }
            steps {
                sh 'go version'
                sh 'go build -o main .'
            }
        }

        stage('Run Tests') {
            agent {
                docker {
                    image 'golang:1.21'
                }
            }
            steps {
                sh 'go test ./...'
            }
        }

        stage('Docker Build') {
            steps {
                sh 'docker build -t $IMAGE_NAME:v1 .'
            }
        }

        stage('Push Image') {
            steps {
                withDockerRegistry([credentialsId: 'dockerhub-creds', url: '']) {
                    sh 'docker push $IMAGE_NAME:v1'
                }
            }
        }

        stage('Deploy') {
            steps {
                sh 'kubectl apply -f k8s/deployment.yaml'
                sh 'kubectl apply -f k8s/service.yaml'
            }
        }
    }
}