cd ./Server
# Build Server

go build -ldflags "-X main.release=true" 
go build -ldflags "-X main.release=false" -o DebugServer

# Build Client

cd ../Client
npm run build

