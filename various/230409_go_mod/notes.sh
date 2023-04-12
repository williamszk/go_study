
go mod init github.com/williamszk/my_first_module

go get rsc.io/quote

go list -m all
# $ go list -m all
# github.com/williamszk/my_first_module
# golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
# rsc.io/quote v1.5.2
# rsc.io/sampler v1.3.0

go get golang.org/x/text
go test


go get rsc.io/sampler
go test

go list -m -versions rsc.io/sampler
# $ go list -m -versions rsc.io/sampler
# rsc.io/sampler v1.0.0 v1.2.0 v1.2.1 v1.3.0 v1.3.1 v1.99.99

go get rsc.io/sampler@v1.3.1
# $ go get rsc.io/sampler@v1.3.1
# go: downloading rsc.io/sampler v1.3.1
# go get: downgraded rsc.io/sampler v1.99.99 => v1.3.1

go test
# $ go test
# PASS
# ok      github.com/williamszk/my_first_module   0.294s

go get rsc.io/quote/v3
# $ go get rsc.io/quote/v3
# go: downloading rsc.io/quote/v3 v3.1.0
# go get: added rsc.io/quote/v3 v3.1.0

go list -m all
# github.com/williamszk/my_first_module
# golang.org/x/mod v0.8.0
# golang.org/x/sys v0.5.0
# golang.org/x/text v0.9.0
# golang.org/x/tools v0.6.0
# rsc.io/quote v1.5.2
# rsc.io/quote/v3 v3.1.0
# rsc.io/sampler v1.3.1

# semanting import versioning
# https://research.swtch.com/vgo-import

go list -m -versions rsc.io/quote
# $ go list -m -versions rsc.io/quote
# rsc.io/quote v0.9.9-pre1 v1.0.0 v1.1.0 v1.2.0 v1.2.1 v1.3.0 v1.4.0 v1.5.0 v1.5.1 v1.5.2 v1.5.3-pre1

go list -m -versions rsc.io/quote
# $ go list -m -versions rsc.io/quote
# rsc.io/quote v0.9.9-pre1 v1.0.0 v1.1.0 v1.2.0 v1.2.1 v1.3.0 v1.4.0 v1.5.0 v1.5.1 v1.5.2 v1.5.3-pre1

# after removing the places where rsc.io/quote appear
# the go.mod will still keep using it
go list -m all
# $ go list -m all
# github.com/williamszk/my_first_module
# golang.org/x/mod v0.8.0
# golang.org/x/sys v0.5.0
# golang.org/x/text v0.9.0
# golang.org/x/tools v0.6.0
# rsc.io/quote v1.5.2       <--- the old version of rsc.io/quote
# rsc.io/quote/v3 v3.1.0
# rsc.io/sampler v1.3.1

# We use go mod tidy to clean up unused dependencies
go mod tidy


go list -m all
# $ go list -m all
# github.com/williamszk/my_first_module
# golang.org/x/mod v0.8.0
# golang.org/x/sys v0.5.0
# golang.org/x/text v0.9.0
# golang.org/x/tools v0.6.0
# rsc.io/quote/v3 v3.1.0
# rsc.io/sampler v1.3.1




