@Library('custom-shared-library') _

def config = [
    projectName: 'backend-helloworld',
    dockerImage: 'backend-helloworld:latest',
    containerName: 'backend-container',
    containerPort: 8080,
    hostPort: 8000
]

buildDockerImage(config)
deployWithComposeOnHost(config)