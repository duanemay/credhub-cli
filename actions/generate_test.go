package actions_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-cf/cm-cli/actions"

	"net/http"

	"errors"

	"github.com/pivotal-cf/cm-cli/client"
	"github.com/pivotal-cf/cm-cli/config"
	cmcli_errors "github.com/pivotal-cf/cm-cli/errors"
	"github.com/pivotal-cf/cm-cli/models"
	"github.com/pivotal-cf/cm-cli/repositories/repositoriesfakes"
)

var _ = Describe("Generate", func() {

	var (
		subject          Generate
		secretRepository repositoriesfakes.FakeSecretRepository
	)

	BeforeEach(func() {
		config := config.Config{ApiURL: "pivotal.io"}

		subject = NewGenerate(&secretRepository, config)
	})

	Describe("GenerateSecret", func() {
		It("generates a secret", func() {
			request := client.NewGenerateSecretRequest("pivotal.io", "my-name")
			expectedBody := models.SecretBody{Value: "my-value"}
			expectedSecret := models.NewSecret("my-name", expectedBody)
			secretRepository.SendRequestStub = func(req *http.Request) (models.SecretBody, error) {
				Expect(req).To(Equal(request))
				return expectedBody, nil
			}

			secret, err := subject.GenerateSecret("my-name")

			Expect(err).ToNot(HaveOccurred())
			Expect(secret).To(Equal(expectedSecret))
		})

		Describe("Errors", func() {
			It("returns a invalid target error when no api is set", func() {
				subject = NewGenerate(&secretRepository, config.Config{})

				_, error := subject.GenerateSecret("my-secret")

				Expect(error).To(MatchError(cmcli_errors.NewNoTargetUrlError()))
			})

			It("returns an error if the request fails", func() {
				request := client.NewGenerateSecretRequest("pivotal.io", "my-secret")
				expectedError := errors.New("My Special Error")
				secretRepository.SendRequestStub = func(req *http.Request) (models.SecretBody, error) {
					Expect(req).To(Equal(request))
					return models.SecretBody{}, expectedError
				}

				_, err := subject.GenerateSecret("my-secret")

				Expect(err).To(Equal(expectedError))
			})
		})
	})
})