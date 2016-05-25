package commands_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("CM", func() {
	It("handle no params call by outputing help text", func() {
		session := runCommand()

		Eventually(session).Should(Exit(1))
		Eventually(session.Err).Should(Say("Please specify one command of"))
	})
})