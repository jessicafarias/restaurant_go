pipeline {
  agent any
  parameters {
    string(name: 'name_container', defaultValue: 'restaurant-go-api-rest', description: 'nombre de la imagen')
    string(name: 'name_imagen', defaultValue: 'restaurant-go-api-rest', description: 'nombre de la imagen')
    string(name: 'tag_imagen', defaultValue: 'latest', description: 'etiqueta de la imagen')
    string(name: 'puerto_imagen', defaultValue: '8080', description: 'etiqueta de la imagen')
  }
  environment {
    name_final = "${name_container}${tag_imagen}${puerto_imagen}"        
  }
  stages {
    stage('stop/rm') {
      when {
        expression { 
          DOCKER_EXIST = sh(returnStdout: true, script: 'echo "$(docker ps -q --filter name=${name_final})"').trim()
          return  DOCKER_EXIST != '' 
        }
      }
      steps {
        script{
          sh ''' 
            echo *****REMOVE AND STOP CONTAINER*****
            sudo docker stop ${name_final}
            sudo docker rmi ${name_final}
            sudo docker rm -vf ${name_final}
            sudo docker rm ${name_container}
          '''
          }
        }                                   
      }
  stage('testing'){
    steps{
      script{
        sh '''
          echo *****TESTING DUMMY TEST*****
          go test ./... # for all
          # go test for test only main
        '''
      }
    }
  }
  stage('build') {
      steps {
        script{
          sh ''' 
          echo *****BUILDING IMAGE*****
          docker build -t ${name_imagen}:${tag_imagen} .
          '''
          }
        }                                       
      }
      stage('run') {
      steps {
        script{
          sh ''' 
                          echo *****RUN DOCKER*****
              docker run  -dtp ${puerto_imagen}:8080 --name ${name_final} ${name_imagen}:${tag_imagen}
              
          '''
          }
        }                                  
      }
    }
    post {
        always {
            echo 'STATUS'
        }
        success {
            echo 'SUCCESS THE PROJECT SHOULD BE RUN ON LOCALHOST 8080'
        }
        failure {
            echo 'Somethines goes wrong'
        }
        unstable {
            echo 'This will run only if the run was marked as unstable'
        }
        changed {
            echo 'Somethin on Pipeline has changed'
        }
    }
  }