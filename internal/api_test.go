package internal_test

import (
	"fmt"
	. "github.com/petewall/ota-service/internal"
	. "github.com/petewall/ota-service/internal/internalfakes"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
)

var _ = Describe("API", func() {
	var (
		api     *API
		log     *gbytes.Buffer
		res     *httptest.ResponseRecorder
		updater *FakeUpdater
	)
	BeforeEach(func() {
		updater = &FakeUpdater{}
		log = gbytes.NewBuffer()
		api = &API{
			Updater:   updater,
			LogOutput: log,
		}
		res = httptest.NewRecorder()
	})

	Describe("/update", func() {
		When("no mac is sent", func() {
			It("returns 400", func() {
				req, err := http.NewRequest("GET", "/update", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(res.Body.String()).To(Equal("X-ESP8266-STA-MAC is not set"))

				Expect(updater.UpdateCallCount()).To(Equal(0))
			})
		})

		When("no current firmware is sent", func() {
			It("returns 400", func() {
				req, err := http.NewRequest("GET", "/update", nil)
				Expect(err).ToNot(HaveOccurred())
				req.Header.Set("X-ESP8266-STA-MAC", "aa:bb:cc:dd:ee:ff")

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(res.Body.String()).To(Equal("current firmware type was not sent"))

				Expect(updater.UpdateCallCount()).To(Equal(0))
			})
		})

		When("no current firmware version is sent", func() {
			It("returns 400", func() {
				req, err := http.NewRequest("GET", "/update?firmware=bootstrap", nil)
				Expect(err).ToNot(HaveOccurred())
				req.Header.Set("X-ESP8266-STA-MAC", "aa:bb:cc:dd:ee:ff")

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(res.Body.String()).To(Equal("current firmware version was not sent"))

				Expect(updater.UpdateCallCount()).To(Equal(0))
			})
		})

		When("the updater returns an error", func() {
			BeforeEach(func() {
				updater.UpdateReturns(nil, fmt.Errorf("update failed"))
			})

			It("returns 500", func() {
				req, err := http.NewRequest("GET", "/update?firmware=bootstrap&version=1.2.3", nil)
				Expect(err).ToNot(HaveOccurred())
				req.Header.Set("X-ESP8266-STA-MAC", "aa:bb:cc:dd:ee:ff")

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to get update: update failed"))

				Expect(updater.UpdateCallCount()).To(Equal(1))
				mac, firmwareType, firmwareVersion := updater.UpdateArgsForCall(0)
				Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("1.2.3"))
			})
		})

		When("the updater returns no firmware", func() {
			BeforeEach(func() {
				updater.UpdateReturns(nil, nil)
			})

			It("returns 304", func() {
				req, err := http.NewRequest("GET", "/update?firmware=bootstrap&version=1.2.3", nil)
				Expect(err).ToNot(HaveOccurred())
				req.Header.Set("X-ESP8266-STA-MAC", "aa:bb:cc:dd:ee:ff")

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusNotModified))

				Expect(updater.UpdateCallCount()).To(Equal(1))
				mac, firmwareType, firmwareVersion := updater.UpdateArgsForCall(0)
				Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("1.2.3"))
			})
		})

		When("the updater returns a firmware", func() {
			BeforeEach(func() {
				updater.UpdateReturns([]byte("this is the firmware data"), nil)
			})

			It("returns the firmware", func() {
				req, err := http.NewRequest("GET", "/update?firmware=bootstrap&version=1.2.3", nil)
				Expect(err).ToNot(HaveOccurred())
				req.Header.Set("X-ESP8266-STA-MAC", "aa:bb:cc:dd:ee:ff")

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK))
				Expect(res.Header().Get("Content-Type")).To(Equal("application/octet-stream"))
				Expect(res.Body.String()).To(Equal("this is the firmware data"))

				Expect(updater.UpdateCallCount()).To(Equal(1))
				mac, firmwareType, firmwareVersion := updater.UpdateArgsForCall(0)
				Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("1.2.3"))
			})
		})
	})
})
