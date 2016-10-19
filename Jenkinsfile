#!/usr/bin/env groovy

node {
  stage('Setup') {
    // $GOROOT is required when installing to a custom location
    // (default location is /usr/local/go/bin)
    env.GOPATH='${WORKSPACE}'
      env.GO15VENDOREXPERIMENT=1
      sh 'go get github.com/stretchr/testify/assert'
      sh 'go get github.com/jstemmer/go-junit-report'
      sh 'go install github.com/jstemmer/go-junit-report'
  }

  stage('Checkout') {
    checkout([$class: 'GitSCM', branches: [[name: '**']], doGenerateSubmoduleConfigurations: false, extensions: [], submoduleCfg: [], userRemoteConfigs: [[url: 'https://github.com/narikin/jenkins_pipeline_golang.git']]])
  }

  stage('Build') {
    sh 'go build -o fizzbuzz'
    archiveArtifacts 'fizzbuzz'
  }

  stage('Test') {
    sh 'go test -v | ${GOPATH}/bin/go-junit-report > report.xml'
    junit 'report.xml'
  }
}
