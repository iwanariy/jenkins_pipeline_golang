#!/usr/bin/env groovy

node {
    // FIXME: workaround because Go Plugin is not available in pipeline currently
    withEnv(['GOROOT=/var/lib/jenkins/tools/org.jenkinsci.plugins.golang.GolangInstallation/Go_1.6.2/', 'PATH=/var/lib/jenkins/tools/org.jenkinsci.plugins.golang.GolangInstallation/Go_1.6.2/bin:$PATH']) {
        stage 'Setup'
        env.GOPATH='${JENKINS_HOME}/workspace/${JOB_NAME}/'
        env.GO15VENDOREXPERIMENT=1


        stage 'Checkout'
            git 'https://github.com/narikin/jenkins_pipeline_golang.git'


        stage 'Build'


        stage 'Test'
  }
}
