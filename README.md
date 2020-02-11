# get-token

🔑 Get an Azure Active Directory token using the OAuth 2.0 Device Authorization Grant

## Usage

Application Options:

- `/t, /tenantID`:             The directory ID
- `/c, /clientID`:             The application ID
- `/s, /scope`:                A space-separated list of scopes
- `/a, /authorizationServer`:  The authorization Server URL (default: https://login.microsoftonline.com)

## Example
`.\get-token.exe /t 11111111-2222-3333-4444-555555555555 /c 66666666-7777-8888-9999-000000000000 /s "openid profile email"`
