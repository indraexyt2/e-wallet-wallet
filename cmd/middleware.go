package cmd

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"e-wallet-wallet/constants"
	"e-wallet-wallet/helpers"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func (d *Dependency) MiddlewareValidateToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		log.Error("Authorization header is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}

	tokenData, err := d.External.ValidateToken(c.Request.Context(), auth)
	if err != nil {
		log.Error("failed to validate token", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}

	c.Set("token", tokenData)
	c.Next()
}

func (d *Dependency) MiddlewareSignatureValidation(c *gin.Context) {
	var (
		log = helpers.Logger
	)

	clientID := c.Request.Header.Get("Client-ID")
	if clientID == "" {
		log.Error("Signature is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}

	secretKey := constants.MappingClient[clientID]
	if secretKey == "" {
		log.Error("client id not found")
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}

	timestamp := c.Request.Header.Get("Timestamp")
	if timestamp == "" {
		log.Error("Timestamp is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}

	requestTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil || time.Since(requestTime) > 5*time.Minute {
		log.Error("invalid timestamp: ", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}

	signature := c.Request.Header.Get("Signature")
	if signature == "" {
		log.Error("Signature is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}

	strPayload := ""

	if c.Request.Method != http.MethodGet {
		byteData, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Error("failed to read request body", err)
			helpers.SendResponseHTTP(
				c,
				http.StatusInternalServerError,
				false,
				"unauthorized",
				nil,
			)
			c.Abort()
			return
		}

		endpoint := c.Request.URL.Path
		strPayload = string(byteData)
		re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
		strPayload = re.ReplaceAllString(strPayload, "")
		strPayload = strings.ToLower(strPayload) + timestamp + endpoint

		copyBody := io.NopCloser(bytes.NewBuffer(byteData))
		c.Request.Body = copyBody
	}

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(strPayload))
	generateSignature := hex.EncodeToString(h.Sum(nil))

	if signature != generateSignature {
		log.Error(fmt.Printf("invalid signature, signature: %s, generateSignature: %s", signature, generateSignature))
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"unauthorized",
			nil,
		)
		c.Abort()
		return
	}
	c.Set("Client-ID", clientID)
	c.Next()
}
