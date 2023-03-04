# go-chatgpt

## Quick start
```
go-chatgpt --authorization <replace your chatgpt api key> "hello"
```
Your chatgpt api key should not have a "Bearer " prefix

## Config
> The first four are the same as the official chatgpt parameters, see: https://platform.openai.com/docs/guides/chat
 - `authorization` chatgpt api key (without 'Bearer ' as prefix)
 - `model` same as the official chatgpt parameter, default: 'gpt-3.5-turbo'
 - `temperature` same as the official chatgpt parameters
 - `max_tokens` same as the official chatgpt parameters
 - `proxy` proxy, supports Socks5 or HTTP
 
There are two ways to configure parameters
### via command parameters
Prefix with `--` in the startup command, like this
```
go-chatgpt -- authorization <your chatgpt api key> -- model gpt-3.5-turb "Hi~chatgpt~"
```
### via config.json
Run `cp config.json.template config.json`

Fill in the appropriate values in the `config.json` file

Then run command, like this
```
go-chatgpt "Hi~chatgpt~"
```

The priority is 'command parameters' > 'config.json'
## Session mode
Use in a complete conversational context until you actively exit (or exceed token limit)

If you don’t add what you want to ask to the start command, you enter session mode, such as
```
go-chatgpt --authorization <replace your chatgpt api key>
```
Then start talking to Chatgpt~

Enter `exit` or end the terminal to exit the session
