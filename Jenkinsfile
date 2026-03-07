pipeline{
    agent any {
        environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-creds')
        IMAGE_NAME = 'vimala92/devops-pipeline-demo'
    }
        stages{
            stage ('checkout') {
                steps{
                    git 'https://github.com/Vimala2020/devops-go-lang-demo.git'
                }
            }
            stage('buid'){
                steps{
                    sh 'go build -o main .'
                }
            }
            stage('Test'){
                steps{
                    sh 'go test ./..'
                }
            }
            stage('Docker build'){
                steps{
                    sh 'docker build -t $IMAGE_NAME:v1 .'
                }
            }
            stage('Docker Push') {
            steps {
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

        stage('SonarQube Scan') {
            steps {
                withSonarQubeEnv('sonarqube') { // 'sonarqube' = configured server in Jenkins
                sh 'sonar-scanner -Dsonar.projectKey=devops-demo -Dsonar.sources=.'
        }
    }
}

    }
}