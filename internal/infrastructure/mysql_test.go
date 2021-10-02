package infrastructure_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MSHR-Dec/go_backend/internal/infrastructure"
)

var _ = Describe("Mysql", func() {
	var (
		user     string
		password string
		host     string
		db       string
	)

	BeforeEach(func() {
		user = "patune"
		password = "patune"
		host = "127.0.0.1:53306"
		db = "patune"
	})

	Describe("Test initializing MySQL", func() {
		Context("Connect MySQL", func() {
			It("has connected successfully", func() {
				Expect(InitMySQL(user, password, host, db)).NotTo(BeNil())
			})
		})
	})
})
