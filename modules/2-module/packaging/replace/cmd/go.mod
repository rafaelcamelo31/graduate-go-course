module github.com/rafaelcamelo31/packaging

go 1.22.4

replace github.com/rafaelcamelo31/graduate-go-course/2-module/packaging/math => ../math

replace github.com/rafaelcamelo31/graduate-go-course/2-module/packaging/route => ../route

require (
	github.com/rafaelcamelo31/graduate-go-course/2-module/packaging/math v0.0.0-00010101000000-000000000000
	github.com/rafaelcamelo31/graduate-go-course/2-module/packaging/route v0.0.0-00010101000000-000000000000
)
