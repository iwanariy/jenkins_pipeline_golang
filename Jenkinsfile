#!/usr/bin/env groovy

node {
  // FIXME: workaround because Go Plugin is not available in pipeline currently
  withEnv(['GOROOT=/var/lib/jenkins/tools/org.jenkinsci.plugins.golang.GolangInstallation/Go_1.6.2/', 'PATH=/var/lib/jenkins/tools/org.jenkinsci.plugins.golang.GolangInstallation/Go_1.6.2/bin:$PATH']) {
    env.GOPATH='${JENKINS_HOME}/workspace/${JOB_NAME}/'
    env.GO15VENDOREXPERIMENT=1

    stage 'Setup'


    stage 'Build'


    stage 'Test'
  }
}