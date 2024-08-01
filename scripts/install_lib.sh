#!/bin/bash

MODULE_NAME="go-backend-crm-gin"

echo "Initializing Go module..."
go mod init $MODULE_NAME

echo "Installing necessary dependencies..."

# Install Gin and other necessary packages
go get github.com/gin-gonic/gin
go get github.com/google/uuid
go get github.com/spf13/viper
go get github.com/stretchr/testify

echo "Installation complete."
