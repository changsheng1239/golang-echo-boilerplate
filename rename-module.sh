NEW_MODULE_NAME=new-name

go mod edit -module $NEW_MODULE_NAME

find . -type f -name '*.go' \
    -exec sed -i -e "s,echo-boilerplate,$NEW_MODULE_NAME,g" {} \;
