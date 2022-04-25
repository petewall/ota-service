package lib_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/ota-service/v2/lib"
	. "github.com/petewall/ota-service/v2/lib/libfakes"
)

type FailingReader struct {
	Message string
}

func (r *FailingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New(r.Message)
}

var _ = Describe("Device Service", func() {
	var (
		httpClient    *FakeHTTPClient
		deviceService DeviceService
	)
	BeforeEach(func() {
		httpClient = &FakeHTTPClient{}
		deviceService = &DeviceServiceImpl{
			Host:       "example.petewall.net",
			Port:       9876,
			HTTPClient: httpClient,
		}

		device := &Device{
			MAC:              "aa:bb:cc:dd:ee:ff",
			CurrentFirmware:  "bootstrap",
			CurrentVersion:   "1.2.3",
			AssignedFirmware: "switch",
			AssignedVersion:  "2.0.0",
		}

		encoded, err := json.Marshal(device)
		Expect(err).ToNot(HaveOccurred())

		response := &http.Response{
			StatusCode: http.StatusOK,
			Body:       ioutil.NopCloser(bytes.NewReader(encoded)),
		}
		httpClient.GetReturns(response, nil)
	})

	Describe("GetDevice", func() {
		When("getting a device", func() {
			It("returns the device", func() {
				device, err := deviceService.GetDevice("aa:bb:cc:dd:ee:ff")
				Expect(err).ToNot(HaveOccurred())
				Expect(device.MAC).To(Equal("aa:bb:cc:dd:ee:ff"))
				Expect(device.CurrentFirmware).To(Equal("bootstrap"))
				Expect(device.CurrentVersion).To(Equal("1.2.3"))
				Expect(device.AssignedFirmware).To(Equal("switch"))
				Expect(device.AssignedVersion).To(Equal("2.0.0"))
			})
		})

		When("getting the device fails", func() {
			BeforeEach(func() {
				httpClient.GetReturns(nil, errors.New("get device failed"))
			})
			It("returns an error", func() {
				_, err := deviceService.GetDevice("aa:bb:cc:dd:ee:ff")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to get device: get device failed"))
			})
		})

		When("reading the response body fails", func() {
			BeforeEach(func() {
				response := &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(&FailingReader{Message: "read failed"}),
				}
				httpClient.GetReturns(response, nil)
			})
			It("returns an error", func() {
				_, err := deviceService.GetDevice("aa:bb:cc:dd:ee:ff")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to read device response: read failed"))
			})
		})

		When("parsing] the response body fails", func() {
			BeforeEach(func() {
				response := &http.Response{
					StatusCode: http.StatusOK,
					Body:       ioutil.NopCloser(strings.NewReader("this is --- not json")),
				}
				httpClient.GetReturns(response, nil)
			})
			It("returns an error", func() {
				_, err := deviceService.GetDevice("aa:bb:cc:dd:ee:ff")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("failed to parse device response: invalid character 'h' in literal true (expecting 'r')"))
			})
		})
	})
})
