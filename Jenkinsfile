pipeline {
  agent any
    stages {
      stage("Build Image") {
        steps {
          sh '''
            source bin/env
            arm build
            '''
        }
      }

      stage("Push Image") {
        when { branch 'master' }
        steps {
          sh "arm login"
          sh "DOCKER_REGISTRY=docker.io/armory arm push"
          archiveArtifacts artifacts: 'build.properties'
        }
      }
    }
}
