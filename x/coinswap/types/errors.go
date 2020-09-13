package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace sdk.CodespaceType = ModuleName

	CodeInvalidInput   CodeType = 101
	CodeInvalidAddress CodeType = sdk.CodeInvalidAddress
	CodeUnauthorized   CodeType = sdk.CodeUnauthorized
	CodeInternal       CodeType = sdk.CodeInternal
	CodeUnknownRequest CodeType = sdk.CodeUnknownRequest
)
