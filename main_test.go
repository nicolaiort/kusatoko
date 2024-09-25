package main

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Main", func() {
	BeforeEach(func() {
		initialize()
	})

	Context("handleRoot", func() {
		When("`/` is requested", func() {
			var rr *httptest.ResponseRecorder
			BeforeEach(func() {
				req, err := http.NewRequest("GET", "/", nil)
				Expect(err).To(BeNil())

				rr = httptest.NewRecorder()
				handler := http.HandlerFunc(handleRoot)

				handler.ServeHTTP(rr, req)
			})
			It("should return 'Hello, World!'", func() {
				Expect(rr.Body.String()).To(Equal("Hello, World!"))
			})
			It("should return status code 200", func() {
				Expect(rr.Code).To(Equal(http.StatusOK))
			})
		})
	})
	Context("handleHealthz", func() {
		When("`/healthz` is requested", func() {
			var rr *httptest.ResponseRecorder
			BeforeEach(func() {
				req, err := http.NewRequest("GET", "/healthz", nil)
				Expect(err).To(BeNil())

				rr = httptest.NewRecorder()
				handler := http.HandlerFunc(handleHealthz)

				handler.ServeHTTP(rr, req)
			})
			It("should return 'ok'", func() {
				Expect(rr.Body.String()).To(Equal("ok"))
			})
			It("should return status code 200", func() {
				Expect(rr.Code).To(Equal(http.StatusOK))
			})
		})
	})

	Context("handleWhatsMyIP", func() {
		When("`/whatsmyip` is requested", func() {
			var rr *httptest.ResponseRecorder
			BeforeEach(func() {
				req, err := http.NewRequest("GET", "/whatsmyip", nil)
				Expect(err).To(BeNil())

				rr = httptest.NewRecorder()
				handler := http.HandlerFunc(handleWhatsMyIP)

				handler.ServeHTTP(rr, req)
			})
			It("should return the IP address", func() {
				Expect(rr.Body.String()).To(Equal(""))
			})
			It("should return status code 200", func() {
				Expect(rr.Code).To(Equal(http.StatusOK))
			})
		})
	})

	Context("handleHeaders", func() {
		When("`/headers` is requested", func() {
			var rr *httptest.ResponseRecorder
			BeforeEach(func() {
				req, err := http.NewRequest("GET", "/headers", nil)
				req.Header.Set("X-Test-Header", "test")
				Expect(err).To(BeNil())

				rr = httptest.NewRecorder()
				handler := http.HandlerFunc(handleHeaders)

				handler.ServeHTTP(rr, req)
			})
			It("should return the headers", func() {
				Expect(rr.Body.String()).To(ContainSubstring("X-Test-Header: test"))
			})
			It("should return status code 200", func() {
				Expect(rr.Code).To(Equal(http.StatusOK))
			})
		})
	})

	Context("handleStatus", func() {
		When("`/status/200` is requested", func() {
			var rr *httptest.ResponseRecorder
			BeforeEach(func() {
				req, err := http.NewRequest("GET", "/status/200", nil)
				Expect(err).To(BeNil())

				rr = httptest.NewRecorder()
				handler := http.HandlerFunc(handleStatus)

				handler.ServeHTTP(rr, req)
			})
			It("should return status code 200", func() {
				Expect(rr.Code).To(Equal(200))
			})
		})
		When("`/status/302` is requested", func() {
			var rr *httptest.ResponseRecorder
			BeforeEach(func() {
				req, err := http.NewRequest("GET", "/status/302", nil)
				Expect(err).To(BeNil())

				rr = httptest.NewRecorder()
				handler := http.HandlerFunc(handleStatus)

				handler.ServeHTTP(rr, req)
			})
			It("should return status code 302", func() {
				Expect(rr.Code).To(Equal(302))
			})
		})
		When("`/status/404` is requested", func() {
			var rr *httptest.ResponseRecorder
			BeforeEach(func() {
				req, err := http.NewRequest("GET", "/status/404", nil)
				Expect(err).To(BeNil())

				rr = httptest.NewRecorder()
				handler := http.HandlerFunc(handleStatus)

				handler.ServeHTTP(rr, req)
			})
			It("should return status code 404", func() {
				Expect(rr.Code).To(Equal(404))
			})
		})
	})
})
