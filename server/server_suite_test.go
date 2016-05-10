package server

import (
  "os"
  "testing"
  "io/ioutil"

  log "github.com/Sirupsen/logrus"
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"

  "github.com/EMC-CMD/cf-persist-service-broker/mocks"
)

func TestServer(t *testing.T) {
  log.SetOutput(ioutil.Discard)

  RegisterFailHandler(Fail)
  RunSpecs(t, "Server Suite")
}

var _ = BeforeSuite(func() {
  err := godotenv.Load("../test.env")
  Expect(err).ToNot(HaveOccurred())

  go startServer()
})

func startServer() {
  os.Chdir("..")
  os.Setenv("PORT", "9900")
  s := Server{}
  devNull, err := os.Open(os.DevNull)
  if err != nil {
    log.Panic("Unable to open ", os.DevNull, err)
  }
  gin.SetMode(gin.ReleaseMode)
  gin.DefaultWriter = devNull
  gin.LoggerWithWriter(ioutil.Discard)
  s.SetClient(&mocks.MockClient{})
  s.Run("9900")
}

