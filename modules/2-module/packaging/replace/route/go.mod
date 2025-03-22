module github.com/rafaelcamelo31/route

go 1.22.4

// Replace can be used to share packages between modules
replace github.com/rafaelcamelo31/graduate-go-course/2-module/packaging/math => ../math

require github.com/rafaelcamelo31/graduate-go-course/2-module/packaging/math v0.0.0-00010101000000-000000000000
