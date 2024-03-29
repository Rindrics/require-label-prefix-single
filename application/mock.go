package application

import (
	"github.com/stretchr/testify/mock"
)

type MockExitAction struct {
	mock.Mock
}

func (m *MockExitAction) Perform() error {
	args := m.Called()
	return args.Error(0)
}

type MockCommenter struct {
	mock.Mock
	PostCommentParams
}

func (m *MockCommenter) PostComment(p PostCommentParams) error {
	args := m.Called(p)
	return args.Error(0)
}

type MockLabeler struct {
	mock.Mock
	AddLabelsParams AddLabelsParams
}

func (m *MockLabeler) AddLabels(p AddLabelsParams) error {
	args := m.Called(p)
	return args.Error(0)
}

type MockCommand struct {
	mock.Mock
	OnSuccess MockExitAction
}

func (m *MockCommand) Execute() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockCommand) Perform() error {
	return m.Execute()
}

type MockLogger struct {
	mock.Mock
}

func (m *MockLogger) Info(msg string, args ...interface{}) {
	m.Called(msg, args)
}

func (m *MockLogger) Debug(msg string, args ...interface{}) {
	m.Called(msg, args)
}

func (m *MockLogger) Error(msg string, args ...interface{}) {
	m.Called(msg, args)
}

type MockGitHubClient struct {
	mock.Mock
}

func (m *MockGitHubClient) AddLabels(p AddLabelsParams) error {
	args := m.Called(p)
	return args.Error(0)
}

func (m *MockGitHubClient) PostComment(p PostCommentParams) error {
	args := m.Called(p)
	return args.Error(0)
}
