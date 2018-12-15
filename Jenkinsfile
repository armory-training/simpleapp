node {
    agent any

    stage("Build Image") {
        sh '''
          source bin/env
          arm build
        '''
    }

    stage("Push Image") {
        if (env.BRANCH_NAME == "master") {
            sh "arm login"
            sh "DOCKER_REGISTRY=docker.io/armory arm push"
            archiveArtifacts artifacts: 'build.properties'
        }
    }
}
