// Code generated by mockery v1.0.0. DO NOT EDIT.
package manager

import context "context"
import mock "github.com/stretchr/testify/mock"
import tournament "github.com/Ragnar-BY/gamingwebsite_testtask/tournament"

// MockTournamentDB is an autogenerated mock type for the TournamentDB type
type MockTournamentDB struct {
	mock.Mock
}

// CreateTournament provides a mock function with given fields: ctx, deposit
func (_m *MockTournamentDB) CreateTournament(ctx context.Context, deposit float32) (int, error) {
	ret := _m.Called(ctx, deposit)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, float32) int); ok {
		r0 = rf(ctx, deposit)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, float32) error); ok {
		r1 = rf(ctx, deposit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTournament provides a mock function with given fields: ctx, id
func (_m *MockTournamentDB) DeleteTournament(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TournamentByID provides a mock function with given fields: ctx, id
func (_m *MockTournamentDB) TournamentByID(ctx context.Context, id int) (tournament.Tournament, error) {
	ret := _m.Called(ctx, id)

	var r0 tournament.Tournament
	if rf, ok := ret.Get(0).(func(context.Context, int) tournament.Tournament); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(tournament.Tournament)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTournament provides a mock function with given fields: ctx, id, t
func (_m *MockTournamentDB) UpdateTournament(ctx context.Context, id int, t tournament.Tournament) error {
	ret := _m.Called(ctx, id, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, tournament.Tournament) error); ok {
		r0 = rf(ctx, id, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
