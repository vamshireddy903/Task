pipeline {
    agent any

    environment {
        IMAGE_NAME = "vamshi589/go-app"
        TAG = "${BUILD_NUMBER}"
        FULL_IMAGE = "${IMAGE_NAME}:${TAG}"
    }

    stages {
        stage('Clone') {
            steps {
                echo "Cloning repository..."
                git branch: 'main', url: 'https://github.com/vamshireddy903/Task.git'
            }
        }

        stage('Build') {
            steps {
                echo "Building Go application..."
                sh 'go build -o go-app'
            }
        }

        stage('Test') {
            steps {
                echo "Running unit tests..."
                sh 'go test ./...'
            }
        }

        stage('Docker Build') {
            steps {
                echo "Building Docker image: ${FULL_IMAGE}"
                sh "docker build -t ${FULL_IMAGE} ."
            }
        }

        stage('Deploy') {
            steps {
                echo "Deploying container locally..."
                sh """
                    docker stop go-app || true
                    docker rm go-app || true
                    docker run -d --name go-app -p 5000:5000 ${FULL_IMAGE}
                """
            }
        }
    }

    post {
        success {
            echo "Pipeline completed successfully!"
        }
        failure {
            echo "Pipeline failed!"
        }
    }
}
