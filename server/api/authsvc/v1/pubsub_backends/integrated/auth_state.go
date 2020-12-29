package authstate

import (
	"errors"
	"sync"

	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	"github.com/harmony-development/legato/server/api/authsvc/v1/authsteps"
	"github.com/harmony-development/legato/server/logger"
)

type AuthSessionState struct {
	CurrentStep authsteps.Step
}

// AuthState ...
type AuthState struct {
	authChannels  map[authv1.AuthService_StreamStepsServer]chan struct{}
	authEvents    map[string]authv1.AuthService_StreamStepsServer
	sessionStates map[string]*AuthSessionState
	Logger        logger.ILogger
	sync.Mutex
}

func New(l logger.ILogger) *AuthState {
	return &AuthState{
		Logger:        l,
		authChannels:  make(map[authv1.AuthService_StreamStepsServer]chan struct{}),
		authEvents:    make(map[string]authv1.AuthService_StreamStepsServer),
		sessionStates: make(map[string]*AuthSessionState),
	}
}

// NewAuthSession ...
func (h *AuthState) NewAuthSession(authID string, step authsteps.Step) error {
	h.Lock()
	defer h.Unlock()

	if _, exists := h.sessionStates[authID]; exists {
		return errors.New("session already exists")
	}

	h.sessionStates[authID] = &AuthSessionState{
		CurrentStep: step,
	}

	return nil
}

func (h *AuthState) AuthSessionExists(authID string) bool {
	h.Lock()
	defer h.Unlock()

	_, exists := h.sessionStates[authID]

	return exists
}

func (h *AuthState) HasStream(authID string) bool {
	h.Lock()
	defer h.Unlock()

	_, exists := h.authEvents[authID]

	return exists
}

// Subscribe ...
func (h *AuthState) Subscribe(authID string, s authv1.AuthService_StreamStepsServer) (chan struct{}, error) {
	h.Lock()
	defer h.Unlock()

	if _, ok := h.sessionStates[authID]; !ok {
		return nil, errors.New("no session state")
	}

	go func() {
		<-s.Context().Done()
		h.DeleteAuthSession(authID)
	}()

	h.authChannels[s] = make(chan struct{})
	h.authEvents[authID] = s

	return h.authChannels[s], nil
}

// Unsubscribe ...
func (h *AuthState) Unsubscribe(authID string, s authv1.AuthService_StreamStepsServer) {
	h.Lock()
	defer h.Unlock()

	close(h.authChannels[s])
	delete(h.authChannels, s)
	h.authEvents[authID] = nil
}

// Broadcast ...
func (h *AuthState) Broadcast(authID string, e *authv1.AuthStep) {
	h.Lock()
	defer h.Unlock()

	serv, ok := h.authEvents[authID]
	if !ok {
		return
	}

	if err := serv.Send(e); err != nil {
		println(err)
	}
}

// SetStep ...
func (h *AuthState) SetStep(authID string, newStep authsteps.Step) {
	h.Lock()
	defer h.Unlock()

	_, ok := h.sessionStates[authID]
	if !ok {
		return
	}
	h.sessionStates[authID].CurrentStep = newStep
}

func (h *AuthState) GetStep(authID string) authsteps.Step {
	h.Lock()
	defer h.Unlock()

	return h.sessionStates[authID].CurrentStep
}

func (h *AuthState) DeleteAuthSession(authID string) {
	h.Lock()
	defer h.Unlock()

	s := h.authEvents[authID]

	if h.authChannels[s] != nil {
		close(h.authChannels[s])
		delete(h.authChannels, s)
	}
	h.authEvents[authID] = nil
	delete(h.sessionStates, authID)
}
