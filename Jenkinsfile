pipeline {
  agent {
    kubernetes {
      inheritFrom 'default'
      defaultContainer "go"
      slaveConnectTimeout 700
      yaml """
apiVersion: v1
kind: Pod
spec:
  containers:
  - name: go
    image: ghcr.io/gythialy/golang-cross-builder:v1.23.2-0-bullseye
    tty: true
    env:
    - name: GOPROXY
      value: "https://go.mirror.douban,direct"
    - name: GOPRIVATE
      value: "*.intra.douban.com"
    resources:
      limits:
        cpu: 3
        memory: 3Gi
      requests:
        cpu: 500m
        memory: 512M
    volumeMounts:
    - mountPath: "/etc/ssl/certs"
      name: "volume-0"
      readOnly: true
"""
    }
  }
  stages {
    stage("test") {
      steps {
        sh "go mod tidy"
        sh "go install gotest.tools/gotestsum@latest"
        sh "gotestsum --junitfile unit-tests.xml -- -p 1 ./..."
      }
      post {
        always {
          // 扫描测试报告
          junit 'unit-tests.xml'
        }
      }
    }
    stage("build") {
      matrix {
        axes {
          axis {
            name 'GOOS'
            values 'windows','linux','darwin'
          }
          axis {
            name 'GOARCH'
            values 'amd64'
          }
        }
        stages {
          stage("build-each-platform") {
            environment {
              FULL_NAME = "archery-cli-${env.GOOS}-${GOARCH}"
            }
            steps {
              script {
                sh "git config --global --add safe.directory '*'"
                if (env.GOOS == "darwin") {
                  sh "CGO_ENABLED=1 CC=x86_64h-apple-darwin20.2-cc SDKROOT=`xcrun --sdk macosx --show-sdk-path` go build -o ./build/${FULL_NAME}"
                } else {
                  sh "go build -o ./build/${FULL_NAME}"
                }
                if (env.GOOS == "windows") {
                  sh "mv build/${FULL_NAME} build/${FULL_NAME}.exe"
                }
              }
            }
          }
        }
      }
      post {
        success {
          archiveArtifacts artifacts: 'build/*', followSymlinks: false
        }
      }
    }
  }
}