pipeline {
  agent any
  stages {
    stage('Deploy') {
      steps {
        sh '''docker-compose build --parallel
docker-compose up'''
      }
    }

  }
}