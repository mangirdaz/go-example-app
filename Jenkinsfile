node('master') {
  stage('Build FE Bin') {
    git url: "https://github.com/bobbydeveaux/go-example-app.git"
    sh "make get-deps"
    sh "make go-build-fe"
  }
  stage('Build Image') {
    sh "oc start-build fe --from-file=fe/ --follow"
  }
  stage('Deploy') {
    openshiftDeploy depCfg: 'fe', namespace: 'ci'
    openshiftVerifyDeployment depCfg: 'fe', replicaCount: 1, verifyReplicaCount: true
  }
  stage('System Test') {
    sh "curl -s -X GET http://fe:8080/ | grep 'UKCloud Image gallery'"
    sh "exit 0"
  }
}