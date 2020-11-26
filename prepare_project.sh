#!/bin/bash
set -e

print_usage () {
    echo "
usage:
   ./prepare_project.sh [-p <sdk-project-name>] [-d <project-description>] [-g <git-repo-url>] [-s <service-category-description>] [-c <service-category-name>]
where:
   -p: specify project name (e.g. platform-services-go-sdk)
   -d: specify project description string (e.g. \"IBM Cloud Platform Services Go SDK\")
   -g: specify the git url (e.g. https://github.com/IBM/ibm-platform-services-go-sdk)
   -s: specify sdk name string (e.g. \"Platform Services\")
   -c: specify the service category (e.g. platform-services)
   -h: view usage instructions
"
}

# Parse flags and arguments
while getopts 'p:d:g:s:c:h' flag; do
  case "${flag}" in
    p) PROJECT_NAME=${OPTARG} ;;
    d) PROJECT_DESCRIPTION=${OPTARG} ;;
    g) PROJECT_GIT_URL=${OPTARG} ;;
    s) SDK_NAME=${OPTARG} ;;
    c) SERVICE_CATEGORY=${OPTARG} ;;
    *) print_usage
        exit 1 ;;
  esac
done

if [[ -z "$PROJECT_NAME" ]]; then
    PROJECT_NAME=$(basename $PWD)
fi

if [[ -z "$SERVICE_CATEGORY" ]]; then
    SERVICE_CATEGORY=$(echo $PROJECT_NAME | sed 's/-go-sdk//')
fi

if [[ -z "$SDK_NAME" ]]; then
    SDK_NAME=$(echo $SERVICE_CATEGORY | tr '-' ' ' |  awk '{for (i=1;i<=NF;i++) $i=toupper(substr($i,1,1)) substr($i,2)} 1')
fi

if [[ -z "$PROJECT_DESCRIPTION" ]]; then
    PROJECT_DESCRIPTION="Go SDK for IBM Cloud ${SDK_NAME} services"
fi

if [[ -z "$PROJECT_GIT_URL" ]]; then
    url=$(git config --get remote.origin.url | sed 's/git@//' | sed 's/com:/com\//' | sed 's/.git$//')
    PROJECT_GIT_URL=${url}
fi

    IMPORT_PATH="$( sed 's~.*://~~' <<< "$PROJECT_GIT_URL" )"

    printf "\n>>>>> Project Initialization In Progress...\n\t IMPORT_PATH: ${IMPORT_PATH}\n\t PROJECT_NAME: ${PROJECT_NAME}\n\t PROJECT_DESCRIPTION: ${PROJECT_DESCRIPTION}\n\t PROJECT_GIT_URL: ${PROJECT_GIT_URL}\n\t SDK_NAME: ${SDK_NAME}\n"

    # Remove sample files
    rm -r exampleservicev1
    printf "\n>>>>> Example Service files removed.\n"

    # Update common Go files
    sed -i.bak 's/my-go-sdk/'${PROJECT_NAME}'/' common/headers.go
    rm go.mod go.sum common/headers.go.bak
    go mod init ${IMPORT_PATH}

    printf "\n>>>>> common Go files updated."

    # Update .travis.yml
    sed -i.bak '/After creating your SDK/,/setup_and_generate.sh/d' .travis.yml
    rm .travis.yml.bak
    printf "\n>>>>> .travis.yml updated."

    # Update documentation
    sed -i.bak "s/^# .*/# ${PROJECT_DESCRIPTION} 0.0.1/" README.md
    sed -i.bak "s/travis.ibm.com/travis-ci.com/" README.md
    sed -i.bak "s/MySDK Service/${SDK_NAME}/" README.md
    sed -i.bak "s/MySDK/${SDK_NAME}/" README.md
    sed -i.bak "s~mysdk~${IMPORT_PATH}~" README.md
    sed -i.bak "s~github.ibm.com/CloudEngineering/go-sdk-template~${IMPORT_PATH}~" README.md
    sed -i.bak "s~<github-repo-url>~${PROJECT_GIT_URL}~" README.md
    sed -i.bak "s/<service-category>/${SERVICE_CATEGORY}/" README.md
    GH_SLUG="$( sed 's~.*.com/~~' <<< "$PROJECT_GIT_URL" )"
    sed -i.bak "s~CloudEngineering/go-sdk-template~${GH_SLUG}~g" README.md
    sed -i.bak "s~^\[Example Service\].*~<!-- [Example Service](https://cloud.ibm.com/apidocs/example-service) | exampleservicev1 -->~" README.md

    rm README.md.bak
    printf "\n>>>>> README.md updated."

    sed -i.bak "s~<github-repo-url>~${PROJECT_GIT_URL}~" CONTRIBUTING.md
    rm CONTRIBUTING.md.bak
    printf "\n>>>>> CONTRIBUTING.md updated."

    printf "\n>>>>> Project Initialized Successfully!\n"
