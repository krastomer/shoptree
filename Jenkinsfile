pipeline {
  agent any
  stages {
    stage('Deploy') {
      steps {
        sh '''docker-compose build backend database
docker-compose up backend database'''
      }
    }

  }
}