package infrastructure_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/MSHR-Dec/go_backend/internal/infrastructure"
)

var _ = Describe("Redis", func() {
	var host string

	BeforeEach(func() {
		host = "127.0.0.1:6379"
	})

	Describe("Test initializing Redis", func() {
		Context("Connect Redis", func() {
			It("has connected successfully", func() {
				Expect(InitRedis(host)).NotTo(BeNil())
			})
		})
	})
})
