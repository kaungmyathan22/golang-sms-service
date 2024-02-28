# SMS otp verification service.


## API Documentation
---
### Send OTP

Send a POST request to the /otp endpoint with the following body to send an OTP to a user's phone number

POST /otp

Request Body

```json
{
  "phoneNumber": "<phone-number-with-country-code>"
}
```

```bash
curl -H "Content-Type: application/json" -X POST -d '{"phoneNumber": "+917420840576"}' http://localhost:8000/otp
```

_Be sure to include the country code in the phone number_

Response

```json
{
  "status": 202,
  "message": "success",
  "data": "OTP sent successfully"
}
```
---
### Verify OTP

Verify a user's OTP by sending a POST request to the /verify endpoint with the following body that contains the phone number and the OTP code received by the user

POST /verifyOTP

Request Body

```json
{
  "user": {
    "phoneNumber": "<phone-number-with-country-code>"
  },
  "code": "<code here>"
}
```

```bash
curl -H "Content-Type: application/json" -X POST -d '{"user": {"phoneNumber": "+917420840576"}, "code":"795279"}' http://localhost:8000/verifyOTP
```

Response

```json
{
  "status": 202,
  "message": "success",
  "data": "OTP verified successfully"
}
```
---