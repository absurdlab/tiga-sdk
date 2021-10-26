#!/bin/bash

go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen \
  --package=commonv1 \
  --generate types,skip-prune \
  -o ./apiv1/common/types.gen.go \
  ../shared/api/v1/common.yaml

go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen \
  --package=loginv1 \
  --generate types \
  -o ./apiv1/login/types.gen.go \
  --import-mapping=./common.yaml:github.com/absurdlab/tiga-sdk/tiga-go-sdk/apiv1/common \
  ../shared/api/v1/login.yaml

go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen \
  --package=consentv1 \
  --generate types \
  -o ./apiv1/consent/types.gen.go \
  --import-mapping=./common.yaml:github.com/absurdlab/tiga-sdk/tiga-go-sdk/apiv1/common \
  ../shared/api/v1/consent.yaml

go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen \
  --package=logoutv1 \
  --generate types \
  -o ./apiv1/logout/types.gen.go \
  --import-mapping=./common.yaml:github.com/absurdlab/tiga-sdk/tiga-go-sdk/apiv1/common \
  ../shared/api/v1/logout.yaml

go mod tidy