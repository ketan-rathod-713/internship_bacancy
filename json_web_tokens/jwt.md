# Json Web Tokens

is a open standard that defines securly transmitting information between parties as JSON object.

This information can be verified and trusted because it is digitally signed.

When to use jwt ?
- authorization
- information exchange.


Json web tokens structure

Header.Payload.Signature

Header : type of token and which signing algorithm used. Then, this JSON is Base64Url encoded to form the first part of the JWT.

Payload : Contains claims, statement about entity. additional data. 3 types of claims. Public private and registered claims.
The payload is then Base64Url encoded to form the second part of the JSON Web Token.

Signature :

It will make signature from 3 parts.

HMACSHA256(
  base64UrlEncode(header) + "." +
  base64UrlEncode(payload),
  secret)

Then putting it all together we will get jwt token.


# Important Methods in golang jwt package

- NewWithClaims(signingMethod, claims Claims, options...) *Token
  - creates new token with method and claims

-token.SignedString(byte secreat)  
  SignedString creates and returns a complete, signed JWT. The token is signed using the SigningMethod specified in the token

