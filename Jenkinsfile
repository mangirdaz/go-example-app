node('maven') {
  stage('Build FE Bin') {
    git url: "https://github.com/bobbydeveaux/go-example-app.git"
    sh "cd fe/ && CGO_ENABLED=0 GOOS=linux go build -o fe . "
  }
  stage('Build Image') {
    unstash name:"jar"
    sh "oc start-build fe --from-file=fe/ --follow"
  }
  stage('Deploy') {
    openshiftDeploy depCfg: 'fe'
    openshiftVerifyDeployment depCfg: 'fe', replicaCount: 1, verifyReplicaCount: true
  }
  stage('System Test') {
    sh "curl -s -X GET http://api:8080/api/v1/img"
    sh "curl -s -X GET http://fe:8080/ | grep 'UKCloud'"
  }
}