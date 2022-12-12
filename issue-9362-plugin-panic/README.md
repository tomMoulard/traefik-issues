ERRO[2022-09-27T10:03:34Z] plugins-local/src/github.com/tomMoulard/ratelimit/ratelimit.go:132:22: panic  plugin=plugin-ratelimit module=github.com/tomMoulard/ratelimit
ERRO[2022-09-27T10:03:34Z] Recovered from panic in HTTP handler [172.18.0.1:35950 - /]: interface conversion: interface {} is interp.valueInterface, not *struct { Xmu sync.Mutex; Xlimit float64; Xburst int; Xtokens float64; Xlast time.Time; XlastEvent time.Time }  middlewareName=traefik-internal-recovery middlewareType=Recovery
ERRO[2022-09-27T10:03:34Z] Stack: goroutine 129 [running]:
github.com/traefik/traefik/v2/pkg/middlewares/recovery.recoverFunc({0x3faf0f8, 0xc0007f6460}, 0xc00021e800)
	github.com/traefik/traefik/v2/pkg/middlewares/recovery/recovery.go:46 +0x225
panic({0x2dfd9e0, 0xc000835ec0})
	runtime/panic.go:1038 +0x215
github.com/traefik/yaegi/interp.runCfg.func1()
	github.com/traefik/yaegi@v0.12.0/interp/run.go:192 +0x145
panic({0x2dfd9e0, 0xc000835ec0})
	runtime/panic.go:1038 +0x215
github.com/traefik/yaegi/interp.typeAssert.func3(0xc001332210)
	github.com/traefik/yaegi@v0.12.0/interp/run.go:440 +0x5c6
github.com/traefik/yaegi/interp.runCfg(0xc000d49320, 0xc001332210, 0x39, 0x352fc20)
	github.com/traefik/yaegi@v0.12.0/interp/run.go:200 +0x2ac
github.com/traefik/yaegi/interp.genFunctionWrapper.func2.1({0xc000591080, 0x2, 0x4})
	github.com/traefik/yaegi@v0.12.0/interp/run.go:1022 +0x4a5
github.com/traefik/yaegi/stdlib._net_http_Handler.ServeHTTP(...)
	github.com/traefik/yaegi@v0.12.0/stdlib/go1_17_net_http.go:288
github.com/traefik/traefik/v2/pkg/middlewares/accesslog.(*FieldHandler).ServeHTTP(0xc00114dd40, {0x3faf0f8, 0xc0007f6460}, 0x2fa2b20)
	github.com/traefik/traefik/v2/pkg/middlewares/accesslog/field_middleware.go:29 +0x122
github.com/gorilla/mux.(*Router).ServeHTTP(0xc0002e5c20, {0x3faf0f8, 0xc0007f6460}, 0xc00021e800)
	github.com/gorilla/mux@v1.8.0/mux.go:141 +0x24c
github.com/traefik/traefik/v2/pkg/middlewares/recovery.(*recovery).ServeHTTP(0x100000000519a0d, {0x3faf0f8, 0xc0007f6460}, 0x40fe54)
	github.com/traefik/traefik/v2/pkg/middlewares/recovery/recovery.go:32 +0x82
github.com/traefik/traefik/v2/pkg/middlewares/accesslog.(*FieldHandler).ServeHTTP(0xc00114de40, {0x3faf0f8, 0xc0007f6460}, 0xc00007c800)
	github.com/traefik/traefik/v2/pkg/middlewares/accesslog/field_middleware.go:29 +0x122
github.com/traefik/traefik/v2/pkg/middlewares.(*HTTPHandlerSwitcher).ServeHTTP(0x4101a7, {0x3faf0f8, 0xc0007f6460}, 0x3f25301)
	github.com/traefik/traefik/v2/pkg/middlewares/handler_switcher.go:23 +0x62
github.com/traefik/traefik/v2/pkg/middlewares/requestdecorator.(*RequestDecorator).ServeHTTP(0xc000206ee0, {0x3faf0f8, 0xc0007f6460}, 0xc00021e700, 0xc00025e8f0)
	github.com/traefik/traefik/v2/pkg/middlewares/requestdecorator/request_decorator.go:47 +0x57b
github.com/traefik/traefik/v2/pkg/middlewares/requestdecorator.WrapHandler.func1.1({0x3faf0f8, 0xc0007f6460}, 0xc000590e10)
	github.com/traefik/traefik/v2/pkg/middlewares/requestdecorator/request_decorator.go:84 +0x68
net/http.HandlerFunc.ServeHTTP(0xc000224af0, {0x3faf0f8, 0xc0007f6460}, 0x9)
	net/http/server.go:2047 +0x2f
github.com/traefik/traefik/v2/pkg/middlewares/forwardedheaders.(*XForwarded).ServeHTTP(0xc000224af0, {0x3faf0f8, 0xc0007f6460}, 0xc00021e700)
	github.com/traefik/traefik/v2/pkg/middlewares/forwardedheaders/forwarded_header.go:189 +0xca
golang.org/x/net/http2/h2c.h2cHandler.ServeHTTP({{0x3f6e540, 0xc000224af0}, 0xc0002a3ac0}, {0x3faf0f8, 0xc0007f6460}, 0xc00021e700)
	golang.org/x/net@v0.0.0-20220425223048-2871e0cb64e4/http2/h2c/h2c.go:104 +0x3e4
net/http.serverHandler.ServeHTTP({0xc000590c60}, {0x3faf0f8, 0xc0007f6460}, 0xc00021e700)
	net/http/server.go:2879 +0x43b
net/http.(*conn).serve(0xc0004225a0, {0x3fd6020, 0xc0002c6210})
	net/http/server.go:1930 +0xb08
created by net/http.(*Server).Serve
	net/http/server.go:3034 +0x4e8  middlewareName=traefik-internal-recovery middlewareType=Recovery
