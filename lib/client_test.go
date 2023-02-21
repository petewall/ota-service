package lib_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"
	"net/http"

	"github.com/petewall/ota-service/lib"
)

var _ = Describe("Client", func() {
	var (
		client     *lib.Client
		server     *ghttp.Server
		data       []byte
		statusCode int
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = &lib.Client{
			Host: server.URL(),
		}
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("Update", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/update", "firmware=test&version=1.2.3"),
					ghttp.VerifyHeaderKV("x-esp8266-sta-mac", "aa:bb:cc:dd:ee:ff"),
					ghttp.RespondWithPtr(&statusCode, &data),
				),
			)
		})

		When("there is an update", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
				data = []byte("this is my firmware data")
			})

			It("sends the update request", func() {
				data, err := client.Update("aa:bb:cc:dd:ee:ff", "test", "1.2.3")
				Expect(err).ToNot(HaveOccurred())
				Expect(data).To(Equal([]byte("this is my firmware data")))
			})
		})

		When("there is no new updates", func() {
			BeforeEach(func() {
				statusCode = http.StatusNotModified
				data = []byte("")
			})

			It("returns 304", func() {
				data, err := client.Update("aa:bb:cc:dd:ee:ff", "test", "1.2.3")
				Expect(err).ToNot(HaveOccurred())
				Expect(data).To(BeNil())
			})
		})

		When("the hostname is invalid", func() {
			BeforeEach(func() {
				client = &lib.Client{
					Host: " invalid:hostname",
				}
			})
			It("returns an error", func() {
				_, err := client.Update("aa:bb:cc:dd:ee:ff", "test", "1.2.3")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(ContainSubstring("failed to parse host ( invalid:hostname): "))
			})
		})

		When("the service returns an unexpected return code", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})
			It("returns an error", func() {
				_, err := client.Update("aa:bb:cc:dd:ee:ff", "test", "1.2.3")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("update request failed: 418 I'm a teapot: "))
			})
		})
	})
})
