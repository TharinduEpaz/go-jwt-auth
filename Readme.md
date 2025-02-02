## Understanding JSON Web Tokens (JWT)

JSON Web Tokens (JWT) are a compact, URL-safe means of representing claims to be transferred between two parties. The claims in a JWT are encoded as a JSON object that is used as the payload of a JSON Web Signature (JWS) structure or as the plaintext of a JSON Web Encryption (JWE) structure. JWTs are commonly used for authentication and information exchange.

### Structure of JWT

A JWT consists of three parts, each separated by a dot (`.`):

1. **Header**: This part typically contains two fields:
   - `alg`: The signing algorithm being used (e.g., HMAC SHA256, RSA).
   - `typ`: The type of token, which is usually "JWT".

   Example of a header in JSON format:
   ```json
   {
     "alg": "HS256",
     "typ": "JWT"
   }
   ```

2. **Payload**: This section contains the claims, which are statements about an entity (typically, the user) and additional data. Claims can include information such as user ID, expiration time, and issued time.

   Example of a payload:
   ```json
   {
     "sub": "user_id123",
     "iat": 1644744000,
     "exp": 1644768000
   }
   ```

3. **Signature**: To create the signature part, you take the encoded header and payload, concatenate them with a period (`.`), and then sign this string using the specified algorithm and a secret key.

   Example of generating a signature:
   ```python
   import hmac
   import hashlib

   secret = 'your-256-bit-secret'
   data = base64url_encode(header) + "." + base64url_encode(payload)
   signature = hmac.new(secret.encode(), data.encode(), hashlib.sha256).digest()
   ```

The final JWT looks like this: `xxxxx.yyyyy.zzzzz`, where `xxxxx` is the Base64Url-encoded header, `yyyyy` is the Base64Url-encoded payload, and `zzzzz` is the signature.

### Is JWT Encrypted?

JWTs can be signed or encrypted. Signing ensures that the sender of the JWT is who they claim to be and that the message wasn't changed along the way. This is achieved through cryptographic algorithms like HMAC or RSA. However, signing does not encrypt the payload; thus, anyone who has access to the token can read its contents.

Encryption can be applied to ensure that only authorized parties can read the payload. In this case, a JWE (JSON Web Encryption) structure would be used. 

### Digital Signature in JWT

JWTs are typically signed using either a secret key with the HMAC algorithm or a public/private key pair using algorithms like RSA or ECDSA. This signing process creates a digital signature that verifies two main aspects:

- Integrity: The signature ensures that the token has not been altered after it was issued. If any part of the JWT (header or payload) is modified, the signature will not match when verified, indicating tampering.
- Authenticity: The signature confirms that the token was created by a legitimate source. For instance, if a JWT is signed with a private key, only the holder of the corresponding public key can verify its authenticity, ensuring that it comes from a trusted issuer.

### Decoding JWTs on jwt.io

The website jwt.io provides a convenient tool for decoding JWTs. When you paste a JWT into jwt.io:

- The tool automatically splits it into its three components: header, payload, and signature.
- It decodes each part from Base64Url encoding back into JSON format.
- You can view both the raw token and its decoded contents.

This allows developers to inspect the claims contained in the token easily and verify its integrity by checking if the signature matches. 

In summary, while JWTs use signing for integrity verification and can optionally use encryption for confidentiality, they are fundamentally structured around three parts: header, payload, and signature.


