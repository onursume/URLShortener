 #!/bin/bash
 
 set -x
 set -m
 
 /entrypoint.sh couchbase-server &
 
 sleep 20
 
 # Setup index and memory quota
 curl -v POST http://127.0.0.1:8091/pools/default -d memoryQuota=2048 -d indexMemoryQuota=2048
 
 # Setup services
 curl -v http://127.0.0.1:8091/node/controller/setupServices -d services=kv%2Cn1ql%2Cindex
 
 # Setup credentials
 curl -v http://127.0.0.1:8091/settings/web -d port=8091 -d username=Administrator -d password=password
 
 # Setup Memory Optimized Indexes
 curl -i -u Administrator:password POST http://127.0.0.1:8091/settings/indexes -d 'storageMode=memory_optimized'
 
 # Load bucket
 curl -v -u Administrator:password POST http://127.0.0.1:8091/pools/default/buckets -d name=Url -d authType=sasl -d saslPassword=123qwe -d bucketType=couchbase -d flushEnabled=1 -d ramQuotaMB=100
 
 sleep 5
  
 fg 1