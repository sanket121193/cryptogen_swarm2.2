echo
echo "############# Join channel request for Org1 ##################"
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTkwODU4NjEsInVzZXJuYW1lIjoiSmltbXkiLCJvcmdOYW1lIjoiT3JnMSIsImlhdCI6MTU5OTA0OTg2MX0.w2k_zMLQlsc-ACWbrcu93kkDIXb6QrLPH38UsxHTK5c" \
  -H "content-type: application/json" \
  -d '{
                "peers": ["peer0.org1.example.com","peer1.org1.example.com","peer2.org1.example.com","peer3.org1.example.com"]
}'
echo
echo "############# Join channel request for Org1 completed. ################"

echo
echo "############# Join channel request for Org2 ##################"
curl -s -X POST \
  http://localhost:4000/channels/mychannel/peers \  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTkwODU4NzIsInVzZXJuYW1lIjoiQmFycnluIiwib3JnTmFtZSI6Ik9yZzIiLCJpYXQiOjE1OTkwNDk4NzJ9.BQxV7mWbeE-sj6W8QOsw7N3IowIfJvgj5p6ho5ZBTIY" \
  -H "content-type: application/json" \
  -d '{
                "peers": ["peer0.org2.example.com"]
}'
echo
echo "############# Join channel request for Org2 completed. ################"


