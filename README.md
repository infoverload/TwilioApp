# TwilioApp

TwilioApp is a simple Go program that sends a random SMS via the Twilio API.

The SMS is a random interesting German word along with its meaning. 

## Configuration

These config variables need to be defined in a JSON file:

- `AccountSid` - Twilio Account SID
- `AuthToken` - Twilio Auth Token
- `Receiver` - Phone number of receiver
- `Sender` - Phone number managed by Twilio

Check out [config.example.json](config/config.example.json)

## Usage

You can `go get` the twilio package by issuing:

```bash
$ go get github.com/subosito/twilio
```

To run the program:
```bash
$ go run main.go
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## To Do
- [ ] write tests