Write-Host "Generating Go gRPC code..."
$protoFiles = Get-ChildItem -Path "./proto" -Recurse -Filter "*.proto" | ForEach-Object { $_.FullName -replace "\\", "/" -replace ".*proto/", "proto/" }
foreach ($file in $protoFiles) {
    protoc --go_opt=module=github.com/esc-chula/intania-vote --go_out=./ --go-grpc_opt=module=github.com/esc-chula/intania-vote --go-grpc_out=./ --proto_path=proto $file
}

Write-Host "Generating TypeScript gRPC code..."
foreach ($file in $protoFiles) {
    protoc --plugin=protoc-gen-ts_proto=.\node_modules\.bin\protoc-gen-ts_proto.cmd --ts_proto_out=.\libs\grpc-ts\src --ts_proto_opt=outputServices=grpc-js,env=node,esModuleInterop=true --proto_path=proto $file
}

Write-Host "Protocol Buffers generation completed!"
