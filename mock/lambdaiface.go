// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/aws/aws-sdk-go/service/lambda/lambdaiface (interfaces: LambdaAPI)

package mock_lambdaiface

import (
	request "github.com/aws/aws-sdk-go/aws/request"
	lambda "github.com/aws/aws-sdk-go/service/lambda"
	gomock "github.com/golang/mock/gomock"
)

// Mock of LambdaAPI interface
type MockLambdaAPI struct {
	ctrl     *gomock.Controller
	recorder *_MockLambdaAPIRecorder
}

// Recorder for MockLambdaAPI (not exported)
type _MockLambdaAPIRecorder struct {
	mock *MockLambdaAPI
}

func NewMockLambdaAPI(ctrl *gomock.Controller) *MockLambdaAPI {
	mock := &MockLambdaAPI{ctrl: ctrl}
	mock.recorder = &_MockLambdaAPIRecorder{mock}
	return mock
}

func (_m *MockLambdaAPI) EXPECT() *_MockLambdaAPIRecorder {
	return _m.recorder
}

func (_m *MockLambdaAPI) AddPermission(_param0 *lambda.AddPermissionInput) (*lambda.AddPermissionOutput, error) {
	ret := _m.ctrl.Call(_m, "AddPermission", _param0)
	ret0, _ := ret[0].(*lambda.AddPermissionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) AddPermission(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddPermission", arg0)
}

func (_m *MockLambdaAPI) AddPermissionRequest(_param0 *lambda.AddPermissionInput) (*request.Request, *lambda.AddPermissionOutput) {
	ret := _m.ctrl.Call(_m, "AddPermissionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.AddPermissionOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) AddPermissionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddPermissionRequest", arg0)
}

func (_m *MockLambdaAPI) CreateAlias(_param0 *lambda.CreateAliasInput) (*lambda.AliasConfiguration, error) {
	ret := _m.ctrl.Call(_m, "CreateAlias", _param0)
	ret0, _ := ret[0].(*lambda.AliasConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) CreateAlias(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAlias", arg0)
}

func (_m *MockLambdaAPI) CreateAliasRequest(_param0 *lambda.CreateAliasInput) (*request.Request, *lambda.AliasConfiguration) {
	ret := _m.ctrl.Call(_m, "CreateAliasRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.AliasConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) CreateAliasRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateAliasRequest", arg0)
}

func (_m *MockLambdaAPI) CreateEventSourceMapping(_param0 *lambda.CreateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	ret := _m.ctrl.Call(_m, "CreateEventSourceMapping", _param0)
	ret0, _ := ret[0].(*lambda.EventSourceMappingConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) CreateEventSourceMapping(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateEventSourceMapping", arg0)
}

func (_m *MockLambdaAPI) CreateEventSourceMappingRequest(_param0 *lambda.CreateEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration) {
	ret := _m.ctrl.Call(_m, "CreateEventSourceMappingRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.EventSourceMappingConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) CreateEventSourceMappingRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateEventSourceMappingRequest", arg0)
}

func (_m *MockLambdaAPI) CreateFunction(_param0 *lambda.CreateFunctionInput) (*lambda.FunctionConfiguration, error) {
	ret := _m.ctrl.Call(_m, "CreateFunction", _param0)
	ret0, _ := ret[0].(*lambda.FunctionConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) CreateFunction(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateFunction", arg0)
}

func (_m *MockLambdaAPI) CreateFunctionRequest(_param0 *lambda.CreateFunctionInput) (*request.Request, *lambda.FunctionConfiguration) {
	ret := _m.ctrl.Call(_m, "CreateFunctionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.FunctionConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) CreateFunctionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateFunctionRequest", arg0)
}

func (_m *MockLambdaAPI) DeleteAlias(_param0 *lambda.DeleteAliasInput) (*lambda.DeleteAliasOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteAlias", _param0)
	ret0, _ := ret[0].(*lambda.DeleteAliasOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) DeleteAlias(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAlias", arg0)
}

func (_m *MockLambdaAPI) DeleteAliasRequest(_param0 *lambda.DeleteAliasInput) (*request.Request, *lambda.DeleteAliasOutput) {
	ret := _m.ctrl.Call(_m, "DeleteAliasRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.DeleteAliasOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) DeleteAliasRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteAliasRequest", arg0)
}

func (_m *MockLambdaAPI) DeleteEventSourceMapping(_param0 *lambda.DeleteEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	ret := _m.ctrl.Call(_m, "DeleteEventSourceMapping", _param0)
	ret0, _ := ret[0].(*lambda.EventSourceMappingConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) DeleteEventSourceMapping(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteEventSourceMapping", arg0)
}

func (_m *MockLambdaAPI) DeleteEventSourceMappingRequest(_param0 *lambda.DeleteEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration) {
	ret := _m.ctrl.Call(_m, "DeleteEventSourceMappingRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.EventSourceMappingConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) DeleteEventSourceMappingRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteEventSourceMappingRequest", arg0)
}

func (_m *MockLambdaAPI) DeleteFunction(_param0 *lambda.DeleteFunctionInput) (*lambda.DeleteFunctionOutput, error) {
	ret := _m.ctrl.Call(_m, "DeleteFunction", _param0)
	ret0, _ := ret[0].(*lambda.DeleteFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) DeleteFunction(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteFunction", arg0)
}

func (_m *MockLambdaAPI) DeleteFunctionRequest(_param0 *lambda.DeleteFunctionInput) (*request.Request, *lambda.DeleteFunctionOutput) {
	ret := _m.ctrl.Call(_m, "DeleteFunctionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.DeleteFunctionOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) DeleteFunctionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteFunctionRequest", arg0)
}

func (_m *MockLambdaAPI) GetAlias(_param0 *lambda.GetAliasInput) (*lambda.AliasConfiguration, error) {
	ret := _m.ctrl.Call(_m, "GetAlias", _param0)
	ret0, _ := ret[0].(*lambda.AliasConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) GetAlias(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAlias", arg0)
}

func (_m *MockLambdaAPI) GetAliasRequest(_param0 *lambda.GetAliasInput) (*request.Request, *lambda.AliasConfiguration) {
	ret := _m.ctrl.Call(_m, "GetAliasRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.AliasConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) GetAliasRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAliasRequest", arg0)
}

func (_m *MockLambdaAPI) GetEventSourceMapping(_param0 *lambda.GetEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	ret := _m.ctrl.Call(_m, "GetEventSourceMapping", _param0)
	ret0, _ := ret[0].(*lambda.EventSourceMappingConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) GetEventSourceMapping(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEventSourceMapping", arg0)
}

func (_m *MockLambdaAPI) GetEventSourceMappingRequest(_param0 *lambda.GetEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration) {
	ret := _m.ctrl.Call(_m, "GetEventSourceMappingRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.EventSourceMappingConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) GetEventSourceMappingRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEventSourceMappingRequest", arg0)
}

func (_m *MockLambdaAPI) GetFunction(_param0 *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error) {
	ret := _m.ctrl.Call(_m, "GetFunction", _param0)
	ret0, _ := ret[0].(*lambda.GetFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) GetFunction(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFunction", arg0)
}

func (_m *MockLambdaAPI) GetFunctionConfiguration(_param0 *lambda.GetFunctionConfigurationInput) (*lambda.FunctionConfiguration, error) {
	ret := _m.ctrl.Call(_m, "GetFunctionConfiguration", _param0)
	ret0, _ := ret[0].(*lambda.FunctionConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) GetFunctionConfiguration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFunctionConfiguration", arg0)
}

func (_m *MockLambdaAPI) GetFunctionConfigurationRequest(_param0 *lambda.GetFunctionConfigurationInput) (*request.Request, *lambda.FunctionConfiguration) {
	ret := _m.ctrl.Call(_m, "GetFunctionConfigurationRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.FunctionConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) GetFunctionConfigurationRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFunctionConfigurationRequest", arg0)
}

func (_m *MockLambdaAPI) GetFunctionRequest(_param0 *lambda.GetFunctionInput) (*request.Request, *lambda.GetFunctionOutput) {
	ret := _m.ctrl.Call(_m, "GetFunctionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.GetFunctionOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) GetFunctionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetFunctionRequest", arg0)
}

func (_m *MockLambdaAPI) GetPolicy(_param0 *lambda.GetPolicyInput) (*lambda.GetPolicyOutput, error) {
	ret := _m.ctrl.Call(_m, "GetPolicy", _param0)
	ret0, _ := ret[0].(*lambda.GetPolicyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) GetPolicy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPolicy", arg0)
}

func (_m *MockLambdaAPI) GetPolicyRequest(_param0 *lambda.GetPolicyInput) (*request.Request, *lambda.GetPolicyOutput) {
	ret := _m.ctrl.Call(_m, "GetPolicyRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.GetPolicyOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) GetPolicyRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetPolicyRequest", arg0)
}

func (_m *MockLambdaAPI) Invoke(_param0 *lambda.InvokeInput) (*lambda.InvokeOutput, error) {
	ret := _m.ctrl.Call(_m, "Invoke", _param0)
	ret0, _ := ret[0].(*lambda.InvokeOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) Invoke(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Invoke", arg0)
}

func (_m *MockLambdaAPI) InvokeAsync(_param0 *lambda.InvokeAsyncInput) (*lambda.InvokeAsyncOutput, error) {
	ret := _m.ctrl.Call(_m, "InvokeAsync", _param0)
	ret0, _ := ret[0].(*lambda.InvokeAsyncOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) InvokeAsync(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InvokeAsync", arg0)
}

func (_m *MockLambdaAPI) InvokeAsyncRequest(_param0 *lambda.InvokeAsyncInput) (*request.Request, *lambda.InvokeAsyncOutput) {
	ret := _m.ctrl.Call(_m, "InvokeAsyncRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.InvokeAsyncOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) InvokeAsyncRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InvokeAsyncRequest", arg0)
}

func (_m *MockLambdaAPI) InvokeRequest(_param0 *lambda.InvokeInput) (*request.Request, *lambda.InvokeOutput) {
	ret := _m.ctrl.Call(_m, "InvokeRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.InvokeOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) InvokeRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "InvokeRequest", arg0)
}

func (_m *MockLambdaAPI) ListAliases(_param0 *lambda.ListAliasesInput) (*lambda.ListAliasesOutput, error) {
	ret := _m.ctrl.Call(_m, "ListAliases", _param0)
	ret0, _ := ret[0].(*lambda.ListAliasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) ListAliases(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAliases", arg0)
}

func (_m *MockLambdaAPI) ListAliasesRequest(_param0 *lambda.ListAliasesInput) (*request.Request, *lambda.ListAliasesOutput) {
	ret := _m.ctrl.Call(_m, "ListAliasesRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.ListAliasesOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) ListAliasesRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAliasesRequest", arg0)
}

func (_m *MockLambdaAPI) ListEventSourceMappings(_param0 *lambda.ListEventSourceMappingsInput) (*lambda.ListEventSourceMappingsOutput, error) {
	ret := _m.ctrl.Call(_m, "ListEventSourceMappings", _param0)
	ret0, _ := ret[0].(*lambda.ListEventSourceMappingsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) ListEventSourceMappings(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListEventSourceMappings", arg0)
}

func (_m *MockLambdaAPI) ListEventSourceMappingsPages(_param0 *lambda.ListEventSourceMappingsInput, _param1 func(*lambda.ListEventSourceMappingsOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListEventSourceMappingsPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLambdaAPIRecorder) ListEventSourceMappingsPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListEventSourceMappingsPages", arg0, arg1)
}

func (_m *MockLambdaAPI) ListEventSourceMappingsRequest(_param0 *lambda.ListEventSourceMappingsInput) (*request.Request, *lambda.ListEventSourceMappingsOutput) {
	ret := _m.ctrl.Call(_m, "ListEventSourceMappingsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.ListEventSourceMappingsOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) ListEventSourceMappingsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListEventSourceMappingsRequest", arg0)
}

func (_m *MockLambdaAPI) ListFunctions(_param0 *lambda.ListFunctionsInput) (*lambda.ListFunctionsOutput, error) {
	ret := _m.ctrl.Call(_m, "ListFunctions", _param0)
	ret0, _ := ret[0].(*lambda.ListFunctionsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) ListFunctions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListFunctions", arg0)
}

func (_m *MockLambdaAPI) ListFunctionsPages(_param0 *lambda.ListFunctionsInput, _param1 func(*lambda.ListFunctionsOutput, bool) bool) error {
	ret := _m.ctrl.Call(_m, "ListFunctionsPages", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockLambdaAPIRecorder) ListFunctionsPages(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListFunctionsPages", arg0, arg1)
}

func (_m *MockLambdaAPI) ListFunctionsRequest(_param0 *lambda.ListFunctionsInput) (*request.Request, *lambda.ListFunctionsOutput) {
	ret := _m.ctrl.Call(_m, "ListFunctionsRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.ListFunctionsOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) ListFunctionsRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListFunctionsRequest", arg0)
}

func (_m *MockLambdaAPI) ListVersionsByFunction(_param0 *lambda.ListVersionsByFunctionInput) (*lambda.ListVersionsByFunctionOutput, error) {
	ret := _m.ctrl.Call(_m, "ListVersionsByFunction", _param0)
	ret0, _ := ret[0].(*lambda.ListVersionsByFunctionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) ListVersionsByFunction(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListVersionsByFunction", arg0)
}

func (_m *MockLambdaAPI) ListVersionsByFunctionRequest(_param0 *lambda.ListVersionsByFunctionInput) (*request.Request, *lambda.ListVersionsByFunctionOutput) {
	ret := _m.ctrl.Call(_m, "ListVersionsByFunctionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.ListVersionsByFunctionOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) ListVersionsByFunctionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListVersionsByFunctionRequest", arg0)
}

func (_m *MockLambdaAPI) PublishVersion(_param0 *lambda.PublishVersionInput) (*lambda.FunctionConfiguration, error) {
	ret := _m.ctrl.Call(_m, "PublishVersion", _param0)
	ret0, _ := ret[0].(*lambda.FunctionConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) PublishVersion(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PublishVersion", arg0)
}

func (_m *MockLambdaAPI) PublishVersionRequest(_param0 *lambda.PublishVersionInput) (*request.Request, *lambda.FunctionConfiguration) {
	ret := _m.ctrl.Call(_m, "PublishVersionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.FunctionConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) PublishVersionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PublishVersionRequest", arg0)
}

func (_m *MockLambdaAPI) RemovePermission(_param0 *lambda.RemovePermissionInput) (*lambda.RemovePermissionOutput, error) {
	ret := _m.ctrl.Call(_m, "RemovePermission", _param0)
	ret0, _ := ret[0].(*lambda.RemovePermissionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) RemovePermission(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePermission", arg0)
}

func (_m *MockLambdaAPI) RemovePermissionRequest(_param0 *lambda.RemovePermissionInput) (*request.Request, *lambda.RemovePermissionOutput) {
	ret := _m.ctrl.Call(_m, "RemovePermissionRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.RemovePermissionOutput)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) RemovePermissionRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemovePermissionRequest", arg0)
}

func (_m *MockLambdaAPI) UpdateAlias(_param0 *lambda.UpdateAliasInput) (*lambda.AliasConfiguration, error) {
	ret := _m.ctrl.Call(_m, "UpdateAlias", _param0)
	ret0, _ := ret[0].(*lambda.AliasConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) UpdateAlias(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateAlias", arg0)
}

func (_m *MockLambdaAPI) UpdateAliasRequest(_param0 *lambda.UpdateAliasInput) (*request.Request, *lambda.AliasConfiguration) {
	ret := _m.ctrl.Call(_m, "UpdateAliasRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.AliasConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) UpdateAliasRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateAliasRequest", arg0)
}

func (_m *MockLambdaAPI) UpdateEventSourceMapping(_param0 *lambda.UpdateEventSourceMappingInput) (*lambda.EventSourceMappingConfiguration, error) {
	ret := _m.ctrl.Call(_m, "UpdateEventSourceMapping", _param0)
	ret0, _ := ret[0].(*lambda.EventSourceMappingConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) UpdateEventSourceMapping(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateEventSourceMapping", arg0)
}

func (_m *MockLambdaAPI) UpdateEventSourceMappingRequest(_param0 *lambda.UpdateEventSourceMappingInput) (*request.Request, *lambda.EventSourceMappingConfiguration) {
	ret := _m.ctrl.Call(_m, "UpdateEventSourceMappingRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.EventSourceMappingConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) UpdateEventSourceMappingRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateEventSourceMappingRequest", arg0)
}

func (_m *MockLambdaAPI) UpdateFunctionCode(_param0 *lambda.UpdateFunctionCodeInput) (*lambda.FunctionConfiguration, error) {
	ret := _m.ctrl.Call(_m, "UpdateFunctionCode", _param0)
	ret0, _ := ret[0].(*lambda.FunctionConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) UpdateFunctionCode(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateFunctionCode", arg0)
}

func (_m *MockLambdaAPI) UpdateFunctionCodeRequest(_param0 *lambda.UpdateFunctionCodeInput) (*request.Request, *lambda.FunctionConfiguration) {
	ret := _m.ctrl.Call(_m, "UpdateFunctionCodeRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.FunctionConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) UpdateFunctionCodeRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateFunctionCodeRequest", arg0)
}

func (_m *MockLambdaAPI) UpdateFunctionConfiguration(_param0 *lambda.UpdateFunctionConfigurationInput) (*lambda.FunctionConfiguration, error) {
	ret := _m.ctrl.Call(_m, "UpdateFunctionConfiguration", _param0)
	ret0, _ := ret[0].(*lambda.FunctionConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) UpdateFunctionConfiguration(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateFunctionConfiguration", arg0)
}

func (_m *MockLambdaAPI) UpdateFunctionConfigurationRequest(_param0 *lambda.UpdateFunctionConfigurationInput) (*request.Request, *lambda.FunctionConfiguration) {
	ret := _m.ctrl.Call(_m, "UpdateFunctionConfigurationRequest", _param0)
	ret0, _ := ret[0].(*request.Request)
	ret1, _ := ret[1].(*lambda.FunctionConfiguration)
	return ret0, ret1
}

func (_mr *_MockLambdaAPIRecorder) UpdateFunctionConfigurationRequest(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateFunctionConfigurationRequest", arg0)
}