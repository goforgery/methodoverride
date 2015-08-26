package methodoverride

import (
	"github.com/goforgery/forgery2"
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestCreate(t *testing.T) {

	Describe("Create()", func() {

		var app *f.Application
		var req *f.Request
		var res *f.Response

		BeforeEach(func() {
			app, req, res, _ = f.CreateAppMock()
		})

		It("should return [GET]", func() {
			app.Use(Create())
			req.Method = "GET"
			app.Handle(req, res, 0)
			AssertEqual(req.Method, "GET")
		})

		It("should return [POST]", func() {
			app.Use(Create())
			req.Method = "GET"
			req.Header.Set("X-HTTP-Method-Override", "POST")
			app.Handle(req, res, 0)
			AssertEqual(req.Method, "POST")
		})

		It("should return [HEAD]", func() {
			app.Use(Create())
			req.Method = "HEAD"
			req.Header.Set("X-HTTP-Method-Override", "POST")
			app.Handle(req, res, 0)
			AssertEqual(req.Map["OriginalMethod"], "HEAD")
		})

		It("should return [DELETE]", func() {
			app.Use(Create())
			req.Method = "GET"
			req.Header.Set("x-http-method-override", "delete")
			app.Handle(req, res, 0)
			AssertEqual(req.Method, "DELETE")
		})

		It("should return [POST]", func() {
			app.Use(Create())
			req.Method = "PUT"
			req.Header.Set("x-http-method-override", "foo")
			app.Handle(req, res, 0)
			AssertEqual(req.Method, "PUT")
		})
	})

	Report(t)
}
