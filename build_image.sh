#!/usr/bin/env bash

printf "\n\nStep1\n\n"
printf "=== Remove previous build static execution file ===\n\n"

# file="./main"

# if [ -f "$file" ]
# then
#     echo "file \"$file\" found."
#     rm main > /dev/null 2>&1
#     if [ $? -eq 1 ]; then
#         echo " process error"
#             exit 1
#     else
#         echo " process success !!"
#     fi
# else    
#     echo " \"file\" $file not found skip to Step2."
# fi

# printf "\n\nStep2-1\n\n"
# printf "=== get all dependecies files ===\n\n"

# go get -d ./...

# if [ $? -eq 1 ]; then    
#     echo " get all dependecies fail"
#     exit 1
# else    
#     echo " get all dependecies success!!"
# fi

# printf "\n\nStep2-2\n\n"
# printf "=== build static execution file ===\n\n"

# CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . 

# if [ $? -eq 1 ]; then    
#     echo " build static execution file fail"
#     exit 1
# else    
#     echo " build static execution file success!!"
# fi

printf "\n\nStep3\n\n"
printf "=== build docker image ===\n\n"
read -p "Enter the Docker ID on DockerHub " username
token=`git log -1 | grep commit`
version=($token)
echo "version number for this build: " ${version[1]} 

docker build -t $username/column_grouping:${version[1]} . 2> /dev/null
build_docker_result=`docker build -t $username/column_grouping:${version[1]} .`
echo $building_docker_result

read -p "Do you want to push image to docker hub registry? (y/n) " yn
case $yn in
    [Yy]* )      
        docker login 
        if [ $? -eq 0 ]; then
            printf "\n\nStep4\n\n"
            printf "=== push docker image to DockerHub/Registry ===\n\n"
            docker push $username/column_grouping:${version[1]}
            if [ $? -eq 0 ]; then
                read -p "Do you want to delete image? (y/n) " yn
                case $yn in
                    [Yy]* ) 
                        printf "\n\nStep5\n\n"
                        printf "=== remove Docker image in local ===\n\n"
                        sudo docker rmi -f $username/column_grouping:${version[1]}
                    ;;
                esac
            fi     
        fi
    ;;
    [Nn]* ) 
        printf "\n\n"
    ;;
esac

printf "\n\n"
echo " _______________________________________________________"
echo "|                                                       |"
echo "|  Whole Process of Docker Images Building is finished. |"
echo "|_______________________________________________________|"
printf "\n\n"
