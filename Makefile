cg:
	cd ./contacts1/components && templ generate

cgr:
	cd ./contacts1/components && templ generate && cd ../ && go run cmd/web/main.go

cr:
	cd ./contacts1 && go run cmd/web/main.go

cexp:
	cd ./contacts1 && go run cmd/exp/main.go

g:
	cd ./hypermedia.system/components && templ generate
	
gr:
	cd ./hypermedia.system/components && templ generate && cd ../ && go run cmd/main.go

r:
	cd ./hypermedia.system && go run cmd/main.go

hs:
	cd ./hypermedia.system && air

templ:
	cd ./templ && air