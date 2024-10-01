#!/bin/bash

case "$OSTYPE" in
    linux*|darwin*)
        DOCKER=$(which docker)
        DOCKER_COMPOSE=$(which docker-compose)
        ;;
    msys*|cygwin*)
        DOCKER="docker"
        DOCKER_COMPOSE="docker-compose"
        ;;
    *)
        echo "Your OS is not supported"
        ;;
esac

PROJECT_DIR=$(pwd)
PROJECT_DOCKER_DIR=$PROJECT_DIR/docker
INFRASTRUCTURE_DIR=$PROJECT_DOCKER_DIR/infrastructure
SERVICES_DIR=$PROJECT_DOCKER_DIR/services
PROJECT_NAME=goquorum-parser
PREFIX=goquorum-parser

isSeed=true

SERVICES_LIST=(backend)
MIGRATIONS_LIST=(backend)
SEEDS_LIST=(seedup-backend)

: '
Checks if docker is installed
'
if [[ ! -x ${DOCKER} ]]; then
    echo "docker is not installed. Please install docker and try again. More info: https://docs.docker.com/engine/install/"
    exit 1
fi

: '
Checks which version of docker compose is installed
'
if [[ ! -x ${DOCKER_COMPOSE} ]]; then
    DOCKER_COMPOSE="${DOCKER} compose"
fi

: '
Checks on existing --noseed key in arguments
You can use this key if you do not want to make seeds
'
if [[ "$@" == *"--noseed"* ]]; then
    isSeed=false
    echo "Deploy without seeds"
fi

: '
startInfrastructure is a function that starting infrastructure containers
Such as:
    - Postgres
    - Nats
'
startInfrastructure() {
        echo "Starting infrastructure containers"
        cd $INFRASTRUCTURE_DIR
        $DOCKER_COMPOSE up -d
}

: '
stopInfrastructure is a function that stopping infrastructure containers
'
stopInfrastructure() {
        echo "Stopping infrastructure containers"
        cd $INFRASTRUCTURE_DIR
        $DOCKER_COMPOSE down
}

: '
stopServices is a function that stopping backend services containers
Such as:
    - Cryptogate
'
stopServices() {
        echo "Stopping services containers"
        cd $SERVICES_DIR
        $DOCKER_COMPOSE down
}

: '
Information table after docker deploying that showing services URLs to connect them
'
servicesInfo() {
        printf "
Backend is running.
"
}

: '
loopServices is a function that running docker commands to up or down backend services
List of services defined in SERVICES_LIST variable
Usage arguments:
    -d      - detach mode for docker containers
    compose - use docker compose instead of docker, if you need to run from compose file
'
loopServices() {
    cd $SERVICES_DIR
    binDocker=$DOCKER
    if [[ "$@" == *"compose"* ]]; then
        binDocker=$DOCKER_COMPOSE
    fi

    if [[ "$@" == *"-d"* ]]; then
        detach="-d"
    fi

    for i in "${SERVICES_LIST[@]}"; do
        echo "$1 $i"
        $binDocker $1 $detach $PREFIX-$i
    done
}

: '
makeMigrations is a function that running migrations
List of migrations defined in MIGRATIONS_LIST variable
'
makeMigrations() {
    for i in "${MIGRATIONS_LIST[@]}"; do
        echo "Doing $i migrations"
        make mup-$i
    done
}

: '
makeSeed is a function that running seeds
List of seeds defined in SEEDS_LIST variable
'
makeSeed() {
    cd $PROJECT_DIR
    echo "Init seeds"
    #cp docker.env .env

    for i in "${SEEDS_LIST[@]}"; do
        echo "Do $i"
        make $i
    done
}

case "$1" in
    start)
        startInfrastructure

        cd $PROJECT_DIR
        echo "Building docker image"
        $DOCKER buildx build -f Dockerfile --tag $PREFIX-services:latest .

        echo "Starting services containers"
        loopServices up -d compose

        if $isSeed; then
            makeSeed
        fi

        servicesInfo
        ;;
    stop)
        stopInfrastructure
        stopServices
        ;;
    restart)
        loopServices stop
        loopServices start

        servicesInfo
        ;;
    reset)
        $0 clear

        # echo "Remove docker volumes"
        # docker volume rm -f $(docker volume ls --filter label=project=$PROJECT_NAME --quiet)
        # docker volume prune -f --filter label=project=$PROJECT_NAME

        startInfrastructure

        sleep 5

        cd $PROJECT_DIR
        export PATH=$PATH:$HOME/go/bin

        if [[ -z $(which migrate) ]]; then
            echo "Install dependencies"
            make install-deps
        fi

        makeMigrations

        if $isSeed; then
            makeSeed
        fi

        ;;
    up)
        if [[ -z $2 ]]; then
            echo "Unknown service. Please use \"$0 $1 {backend}\""
            exit 1
        fi

        if $isSeed; then
            makeSeed
        fi

        echo "Starting services containers"
        cd $SERVICES_DIR
        $DOCKER_COMPOSE up -d $PREFIX-$2

        ;;
    down)
        if [[ -z $2 ]]; then
            echo "Unknown service. Please use \"$0 $1 {backend}\""
            exit 1
        fi

        echo "Stopping services containers"
        cd $SERVICES_DIR
        $DOCKER_COMPOSE down $PREFIX-$2
        ;;
    clear)
        stopInfrastructure

        stopServices

        echo "Remove docker image"
        $DOCKER rmi -f $PREFIX-services:latest
        ;;
    refresh)

        stopServices

        echo "Remove docker image"
        $DOCKER rmi -f $PREFIX-services:latest

        cd $PROJECT_DIR
        $0 start --noseed
        ;;
    seed)
        makeSeed
        ;;
    help)
        printf "Usage: $0 {start|stop|restart|reset|up <service name>|down <service name>|clear|seed|help} <option>

Commands:

    start               - run all services
    stop                - stop and remove all services
    restart             - restart only backend services
    reset               - stop all services, run infrastructure containers and make migrations
    up <service name>   - start specific service
    down <service name> - stop specific service
    refresh             - stop and remove all services, run infrastructure containers and make migrations
    clear               - remove all backend docker images
    seed                - make seeds for databases
    help                - show this info

Options:
    --noseed            - do not make seeds
"
        ;;
    *)
        echo "Unknown operation, please, use \"$0 help\" for more info."
        ;;
esac
