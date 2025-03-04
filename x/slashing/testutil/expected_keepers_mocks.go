// Code generated by MockGen. DO NOT EDIT.
// Source: x/slashing/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	reflect "reflect"

	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types1 "github.com/cosmos/cosmos-sdk/x/params/types"
	types2 "github.com/cosmos/cosmos-sdk/x/staking/types"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// GetAccount mocks base method.
func (m *MockAccountKeeper) GetAccount(ctx types.Context, addr types.AccAddress) types0.AccountI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", ctx, addr)
	ret0, _ := ret[0].(types0.AccountI)
	return ret0
}

// GetAccount indicates an expected call of GetAccount.
func (mr *MockAccountKeeperMockRecorder) GetAccount(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).GetAccount), ctx, addr)
}

// IterateAccounts mocks base method.
func (m *MockAccountKeeper) IterateAccounts(ctx types.Context, process func(types0.AccountI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateAccounts", ctx, process)
}

// IterateAccounts indicates an expected call of IterateAccounts.
func (mr *MockAccountKeeperMockRecorder) IterateAccounts(ctx, process interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateAccounts", reflect.TypeOf((*MockAccountKeeper)(nil).IterateAccounts), ctx, process)
}

// MockBankKeeper is a mock of BankKeeper interface.
type MockBankKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockBankKeeperMockRecorder
}

// MockBankKeeperMockRecorder is the mock recorder for MockBankKeeper.
type MockBankKeeperMockRecorder struct {
	mock *MockBankKeeper
}

// NewMockBankKeeper creates a new mock instance.
func NewMockBankKeeper(ctrl *gomock.Controller) *MockBankKeeper {
	mock := &MockBankKeeper{ctrl: ctrl}
	mock.recorder = &MockBankKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBankKeeper) EXPECT() *MockBankKeeperMockRecorder {
	return m.recorder
}

// GetAllBalances mocks base method.
func (m *MockBankKeeper) GetAllBalances(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllBalances", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// GetAllBalances indicates an expected call of GetAllBalances.
func (mr *MockBankKeeperMockRecorder) GetAllBalances(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllBalances", reflect.TypeOf((*MockBankKeeper)(nil).GetAllBalances), ctx, addr)
}

// GetBalance mocks base method.
func (m *MockBankKeeper) GetBalance(ctx types.Context, addr types.AccAddress, denom string) types.Coin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalance", ctx, addr, denom)
	ret0, _ := ret[0].(types.Coin)
	return ret0
}

// GetBalance indicates an expected call of GetBalance.
func (mr *MockBankKeeperMockRecorder) GetBalance(ctx, addr, denom interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalance", reflect.TypeOf((*MockBankKeeper)(nil).GetBalance), ctx, addr, denom)
}

// LockedCoins mocks base method.
func (m *MockBankKeeper) LockedCoins(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockedCoins", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// LockedCoins indicates an expected call of LockedCoins.
func (mr *MockBankKeeperMockRecorder) LockedCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockedCoins", reflect.TypeOf((*MockBankKeeper)(nil).LockedCoins), ctx, addr)
}

// SpendableCoins mocks base method.
func (m *MockBankKeeper) SpendableCoins(ctx types.Context, addr types.AccAddress) types.Coins {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpendableCoins", ctx, addr)
	ret0, _ := ret[0].(types.Coins)
	return ret0
}

// SpendableCoins indicates an expected call of SpendableCoins.
func (mr *MockBankKeeperMockRecorder) SpendableCoins(ctx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpendableCoins", reflect.TypeOf((*MockBankKeeper)(nil).SpendableCoins), ctx, addr)
}

// MockParamSubspace is a mock of ParamSubspace interface.
type MockParamSubspace struct {
	ctrl     *gomock.Controller
	recorder *MockParamSubspaceMockRecorder
}

// MockParamSubspaceMockRecorder is the mock recorder for MockParamSubspace.
type MockParamSubspaceMockRecorder struct {
	mock *MockParamSubspace
}

// NewMockParamSubspace creates a new mock instance.
func NewMockParamSubspace(ctrl *gomock.Controller) *MockParamSubspace {
	mock := &MockParamSubspace{ctrl: ctrl}
	mock.recorder = &MockParamSubspaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockParamSubspace) EXPECT() *MockParamSubspaceMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockParamSubspace) Get(ctx types.Context, key []byte, ptr interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Get", ctx, key, ptr)
}

// Get indicates an expected call of Get.
func (mr *MockParamSubspaceMockRecorder) Get(ctx, key, ptr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockParamSubspace)(nil).Get), ctx, key, ptr)
}

// GetParamSet mocks base method.
func (m *MockParamSubspace) GetParamSet(ctx types.Context, ps types1.ParamSet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetParamSet", ctx, ps)
}

// GetParamSet indicates an expected call of GetParamSet.
func (mr *MockParamSubspaceMockRecorder) GetParamSet(ctx, ps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParamSet", reflect.TypeOf((*MockParamSubspace)(nil).GetParamSet), ctx, ps)
}

// HasKeyTable mocks base method.
func (m *MockParamSubspace) HasKeyTable() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasKeyTable")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasKeyTable indicates an expected call of HasKeyTable.
func (mr *MockParamSubspaceMockRecorder) HasKeyTable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasKeyTable", reflect.TypeOf((*MockParamSubspace)(nil).HasKeyTable))
}

// SetParamSet mocks base method.
func (m *MockParamSubspace) SetParamSet(ctx types.Context, ps types1.ParamSet) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetParamSet", ctx, ps)
}

// SetParamSet indicates an expected call of SetParamSet.
func (mr *MockParamSubspaceMockRecorder) SetParamSet(ctx, ps interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetParamSet", reflect.TypeOf((*MockParamSubspace)(nil).SetParamSet), ctx, ps)
}

// WithKeyTable mocks base method.
func (m *MockParamSubspace) WithKeyTable(table types1.KeyTable) types1.Subspace {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithKeyTable", table)
	ret0, _ := ret[0].(types1.Subspace)
	return ret0
}

// WithKeyTable indicates an expected call of WithKeyTable.
func (mr *MockParamSubspaceMockRecorder) WithKeyTable(table interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithKeyTable", reflect.TypeOf((*MockParamSubspace)(nil).WithKeyTable), table)
}

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// Delegation mocks base method.
func (m *MockStakingKeeper) Delegation(arg0 types.Context, arg1 types.AccAddress, arg2 types.ValAddress) types2.DelegationI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delegation", arg0, arg1, arg2)
	ret0, _ := ret[0].(types2.DelegationI)
	return ret0
}

// Delegation indicates an expected call of Delegation.
func (mr *MockStakingKeeperMockRecorder) Delegation(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delegation", reflect.TypeOf((*MockStakingKeeper)(nil).Delegation), arg0, arg1, arg2)
}

// IterateValidators mocks base method.
func (m *MockStakingKeeper) IterateValidators(arg0 types.Context, arg1 func(int64, types2.ValidatorI) bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "IterateValidators", arg0, arg1)
}

// IterateValidators indicates an expected call of IterateValidators.
func (mr *MockStakingKeeperMockRecorder) IterateValidators(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IterateValidators", reflect.TypeOf((*MockStakingKeeper)(nil).IterateValidators), arg0, arg1)
}

// Jail mocks base method.
func (m *MockStakingKeeper) Jail(arg0 types.Context, arg1 types.ConsAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Jail", arg0, arg1)
}

// Jail indicates an expected call of Jail.
func (mr *MockStakingKeeperMockRecorder) Jail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jail", reflect.TypeOf((*MockStakingKeeper)(nil).Jail), arg0, arg1)
}

// MaxValidators mocks base method.
func (m *MockStakingKeeper) MaxValidators(arg0 types.Context) uint32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxValidators", arg0)
	ret0, _ := ret[0].(uint32)
	return ret0
}

// MaxValidators indicates an expected call of MaxValidators.
func (mr *MockStakingKeeperMockRecorder) MaxValidators(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxValidators", reflect.TypeOf((*MockStakingKeeper)(nil).MaxValidators), arg0)
}

// Slash mocks base method.
func (m *MockStakingKeeper) Slash(arg0 types.Context, arg1 types.ConsAddress, arg2, arg3 int64, arg4 types.Dec) math.Int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Slash", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(math.Int)
	return ret0
}

// Slash indicates an expected call of Slash.
func (mr *MockStakingKeeperMockRecorder) Slash(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Slash", reflect.TypeOf((*MockStakingKeeper)(nil).Slash), arg0, arg1, arg2, arg3, arg4)
}

// Unjail mocks base method.
func (m *MockStakingKeeper) Unjail(arg0 types.Context, arg1 types.ConsAddress) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Unjail", arg0, arg1)
}

// Unjail indicates an expected call of Unjail.
func (mr *MockStakingKeeperMockRecorder) Unjail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unjail", reflect.TypeOf((*MockStakingKeeper)(nil).Unjail), arg0, arg1)
}

// Validator mocks base method.
func (m *MockStakingKeeper) Validator(arg0 types.Context, arg1 types.ValAddress) types2.ValidatorI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Validator", arg0, arg1)
	ret0, _ := ret[0].(types2.ValidatorI)
	return ret0
}

// Validator indicates an expected call of Validator.
func (mr *MockStakingKeeperMockRecorder) Validator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Validator", reflect.TypeOf((*MockStakingKeeper)(nil).Validator), arg0, arg1)
}

// ValidatorByConsAddr mocks base method.
func (m *MockStakingKeeper) ValidatorByConsAddr(arg0 types.Context, arg1 types.ConsAddress) types2.ValidatorI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(types2.ValidatorI)
	return ret0
}

// ValidatorByConsAddr indicates an expected call of ValidatorByConsAddr.
func (mr *MockStakingKeeperMockRecorder) ValidatorByConsAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorByConsAddr", reflect.TypeOf((*MockStakingKeeper)(nil).ValidatorByConsAddr), arg0, arg1)
}

// MockStakingHooks is a mock of StakingHooks interface.
type MockStakingHooks struct {
	ctrl     *gomock.Controller
	recorder *MockStakingHooksMockRecorder
}

// MockStakingHooksMockRecorder is the mock recorder for MockStakingHooks.
type MockStakingHooksMockRecorder struct {
	mock *MockStakingHooks
}

// NewMockStakingHooks creates a new mock instance.
func NewMockStakingHooks(ctrl *gomock.Controller) *MockStakingHooks {
	mock := &MockStakingHooks{ctrl: ctrl}
	mock.recorder = &MockStakingHooksMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingHooks) EXPECT() *MockStakingHooksMockRecorder {
	return m.recorder
}

// AfterValidatorBonded mocks base method.
func (m *MockStakingHooks) AfterValidatorBonded(ctx types.Context, consAddr types.ConsAddress, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorBonded", ctx, consAddr, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorBonded indicates an expected call of AfterValidatorBonded.
func (mr *MockStakingHooksMockRecorder) AfterValidatorBonded(ctx, consAddr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorBonded", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorBonded), ctx, consAddr, valAddr)
}

// AfterValidatorCreated mocks base method.
func (m *MockStakingHooks) AfterValidatorCreated(ctx types.Context, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorCreated", ctx, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorCreated indicates an expected call of AfterValidatorCreated.
func (mr *MockStakingHooksMockRecorder) AfterValidatorCreated(ctx, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorCreated", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorCreated), ctx, valAddr)
}

// AfterValidatorRemoved mocks base method.
func (m *MockStakingHooks) AfterValidatorRemoved(ctx types.Context, consAddr types.ConsAddress, valAddr types.ValAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AfterValidatorRemoved", ctx, consAddr, valAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// AfterValidatorRemoved indicates an expected call of AfterValidatorRemoved.
func (mr *MockStakingHooksMockRecorder) AfterValidatorRemoved(ctx, consAddr, valAddr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AfterValidatorRemoved", reflect.TypeOf((*MockStakingHooks)(nil).AfterValidatorRemoved), ctx, consAddr, valAddr)
}
