package v1

import (
	"testing"

	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	authstate "github.com/harmony-development/legato/server/api/authsvc/v1/pubsub_backends/integrated"
	"github.com/harmony-development/legato/server/responses"
	"github.com/harmony-development/legato/server/test"
	"github.com/labstack/echo/v4"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

func newAuthAPI() *V1 {
	return New(Dependencies{
		DB:          test.NewMockDB(),
		Logger:      test.MockLogger{},
		Sonyflake:   sonyflake.NewSonyflake(sonyflake.Settings{}),
		AuthManager: test.MockAuthManager{},
		AuthState:   authstate.New(test.MockLogger{}),
		Config:      test.DefaultConf(),
	})
}

func beginAuth(c echo.Context, api *V1) (string, error) {
	resp, err := api.BeginAuth(c, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return resp.AuthId, nil
}

func initialChoice(c echo.Context, api *V1, authID string) (*authv1.AuthStep, error) {
	return api.NextStep(c, &authv1.NextStepRequest{
		AuthId: authID,
	})
}

func TestInitialChoice(t *testing.T) {
	a := require.New(t)
	api := newAuthAPI()
	ctx := test.DummyContext(echo.New())
	authID, err := beginAuth(ctx, api)
	a.NoError(err)
	a.NotEmpty(authID)
	step, err := api.NextStep(ctx, &authv1.NextStepRequest{
		AuthId: authID,
	})
	a.NoError(err)
	a.False(step.CanGoBack)
	a.IsType(&authv1.AuthStep_Choice_{}, step.Step)
	a.Equal("initial-choice", step.GetChoice().Title)
	a.ElementsMatch(step.GetChoice().Options, []string{"login", "register", "other-options"})
}

func TestStepBack(t *testing.T) {
	a := require.New(t)
	api := newAuthAPI()
	ctx := test.DummyContext(echo.New())
	authID, _ := beginAuth(ctx, api)
	_, _ = api.NextStep(ctx, &authv1.NextStepRequest{
		AuthId: authID,
	})
	_, err := api.StepBack(ctx, &authv1.StepBackRequest{
		AuthId: authID,
	})
	a.NotNil(err)
	_, err = api.NextStep(ctx, &authv1.NextStepRequest{
		AuthId: authID,
		Step: &authv1.NextStepRequest_Choice_{
			Choice: &authv1.NextStepRequest_Choice{
				Choice: "login",
			},
		},
	})
	a.NoError(err)
	step, err := api.StepBack(ctx, &authv1.StepBackRequest{
		AuthId: authID,
	})
	a.NoError(err)
	a.IsType(&authv1.AuthStep_Choice_{}, step.Step)
	a.Equal("initial-choice", step.GetChoice().Title)
	a.Equal(step.CanGoBack, false)
}

func TestLogin(t *testing.T) {
	var testTable = []struct {
		email       string
		password    string
		expectError string
	}{
		{"amadeus@home.cern", "@&GyubhjA^GYUH1", ""},
		{"amadeus@home.cern", "", ""},
	}

	for _, testCase := range testTable {
		t.Run(testCase.email, func(t *testing.T) {
			a := require.New(t)
			api := newAuthAPI()
			ctx := test.DummyContext(echo.New())
			hashed, err := bcrypt.GenerateFromPassword([]byte(testCase.password), 0)
			a.NoError(err)
			api.DB.AddLocalUser(12345, testCase.email, "amadeus", hashed)
			authID, _ := beginAuth(ctx, api)
			_, _ = initialChoice(ctx, api, authID)
			_, _ = api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Choice_{
					Choice: &authv1.NextStepRequest_Choice{
						Choice: "login",
					},
				},
			})
			sessionStep, err := api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Form_{
					Form: &authv1.NextStepRequest_Form{
						Fields: []*authv1.NextStepRequest_FormFields{
							{
								Field: &authv1.NextStepRequest_FormFields_String_{
									String_: testCase.email,
								},
							},
							{
								Field: &authv1.NextStepRequest_FormFields_Bytes{
									Bytes: []byte(testCase.password),
								},
							},
						},
					},
				},
			})
			if testCase.expectError != "" {
				a.Error(err)
				a.Equal(testCase.expectError, err.Error())
			} else {
				a.NoError(err)
				a.NotNil(sessionStep)
				a.True(sessionStep.CanGoBack)
				a.IsType(&authv1.AuthStep_Session{}, sessionStep.Step)
				a.Equal(uint64(12345), sessionStep.GetSession().UserId)
				a.Greater(len(sessionStep.GetSession().SessionToken), 8)
				id, err := api.DB.SessionToUserID(sessionStep.GetSession().SessionToken)
				a.NoError(err)
				a.Equal(uint64(12345), id)
			}
		})
	}
}

func TestRegister(t *testing.T) {
	var testTable = []struct {
		email       string
		username    string
		password    string
		expectError string
		name        string
	}{
		{
			name:        "Normal Registration",
			email:       "amadeus@home.cern",
			username:    "Amadeus",
			password:    "@&GyubhjA^GYUH1",
			expectError: "",
		},
		{
			name:        "Registration with a bad email",
			email:       "this is not an email",
			username:    "Amadeus",
			password:    "@&GyubhjA^GYUH1",
			expectError: responses.BadEmail,
		},
		{
			name:        "Registration with a bad password",
			email:       "amadeus@home.cern",
			username:    "Amadeus",
			password:    "this is not a password that will work since its just lowercase",
			expectError: responses.BadPassword,
		},
		{
			name:        "Registration with a short username",
			email:       "amadeus@home.cern",
			username:    "a",
			password:    "@&GyubhjA^GYUH1",
			expectError: responses.BadUsername,
		},
		{
			name:        "Registration with a long username",
			email:       "amadeus@home.cern",
			username:    "Hello my name is very long. Nobody should have a name this long. Long names are annoying. Just a few more words to make sure...",
			password:    "@&GyubhjA^GYUH1",
			expectError: responses.BadUsername,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			a := require.New(t)
			api := newAuthAPI()
			ctx := test.DummyContext(echo.New())
			authID, _ := beginAuth(ctx, api)
			initialChoice(ctx, api, authID)
			api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Choice_{
					Choice: &authv1.NextStepRequest_Choice{
						Choice: "register",
					},
				},
			})
			sessionStep, err := api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Form_{
					Form: &authv1.NextStepRequest_Form{
						Fields: []*authv1.NextStepRequest_FormFields{
							{
								Field: &authv1.NextStepRequest_FormFields_String_{
									String_: testCase.email,
								},
							},
							{
								Field: &authv1.NextStepRequest_FormFields_String_{
									String_: testCase.username,
								},
							},
							{
								Field: &authv1.NextStepRequest_FormFields_Bytes{
									Bytes: []byte(testCase.password),
								},
							},
						},
					},
				},
			})
			if testCase.expectError != "" {
				a.EqualError(err, testCase.expectError)
			} else {
				a.NoError(err)
				a.NotNil(sessionStep)
				a.True(sessionStep.CanGoBack)
				a.IsType(&authv1.AuthStep_Session{}, sessionStep.Step)
				user, err := api.DB.GetUserByID(sessionStep.GetSession().UserId)
				a.NoError(err)
				a.Equal(testCase.username, user.Username)
				userID, err := api.DB.SessionToUserID(sessionStep.GetSession().SessionToken)
				a.NoError(err)
				a.Equal(sessionStep.GetSession().UserId, userID)
				a.Greater(len(sessionStep.GetSession().SessionToken), 8)
			}
		})
	}
}

func BenchmarkBeginAuth(b *testing.B) {
	ctx := test.DummyContext(echo.New())
	api := newAuthAPI()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, _ = beginAuth(ctx, api)
		}
	})
}

func BenchmarkLogin(b *testing.B) {
	ctx := test.DummyContext(echo.New())
	api := newAuthAPI()
	hashed, err := bcrypt.GenerateFromPassword([]byte("@&GyubhjA^GYUH1"), 0)
	if err != nil {
		panic(err)
	}
	_ = api.DB.AddLocalUser(12345, "amadeus@home.cern", "amadeus", hashed)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			authID, err := beginAuth(ctx, api)
			if err != nil {
				panic(err)
			}
			_, err = api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Choice_{
					Choice: &authv1.NextStepRequest_Choice{
						Choice: "login",
					},
				},
			})
			if err != nil {
				panic(err)
			}
			_, err = api.NextStep(ctx, &authv1.NextStepRequest{
				AuthId: authID,
				Step: &authv1.NextStepRequest_Form_{
					Form: &authv1.NextStepRequest_Form{
						Fields: []*authv1.NextStepRequest_FormFields{
							{
								Field: &authv1.NextStepRequest_FormFields_String_{
									String_: "amadeus@home.cern",
								},
							},
							{
								Field: &authv1.NextStepRequest_FormFields_Bytes{
									Bytes: []byte("@&GyubhjA^GYUH1"),
								},
							},
						},
					},
				},
			})
			if err != nil {
				panic(err)
			}
		}
	})
}
