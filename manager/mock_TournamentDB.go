// Code generated by mockery v1.0.0. DO NOT EDIT.
package manager

import mock "github.com/stretchr/testify/mock"
import tournament "github.com/Ragnar-BY/gamingwebsite_testtask/tournament"

// MockTournamentDB is an autogenerated mock type for the TournamentDB type
type MockTournamentDB struct {
	mock.Mock
}

// CreateTournament provides a mock function with given fields: deposit
func (_m *MockTournamentDB) CreateTournament(deposit float32) (int, error) {
	ret := _m.Called(deposit)

	var r0 int
	if rf, ok := ret.Get(0).(func(float32) int); ok {
		r0 = rf(deposit)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(float32) error); ok {
		r1 = rf(deposit)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTournament provides a mock function with given fields: id
func (_m *MockTournamentDB) DeleteTournament(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// TournamentByID provides a mock function with given fields: id
func (_m *MockTournamentDB) TournamentByID(id int) (tournament.Tournament, error) {
	ret := _m.Called(id)

	var r0 tournament.Tournament
	if rf, ok := ret.Get(0).(func(int) tournament.Tournament); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(tournament.Tournament)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTournament provides a mock function with given fields: id, t
func (_m *MockTournamentDB) UpdateTournament(id int, t tournament.Tournament) error {
	ret := _m.Called(id, t)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, tournament.Tournament) error); ok {
		r0 = rf(id, t)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}