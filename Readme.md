# Documentation

JSON Web Tokens (JWT) are a compact, URL-safe means of representing claims to be transferred between two parties. The claims in a JWT are encoded as a JSON object that is used as the payload of a JSON Web Signature (JWS) structure or as the plaintext of a JSON Web Encryption (JWE) structure. JWTs are commonly used for authentication and information exchange.

This a a complete project with full user authentication and planned to implement more robust and RBAC systems in future.

This project is configured to deployed automatically into AWS EC2.

## Table of Contents

- [Documentation](#documentation)
  - [Table of Contents](#table-of-contents)
  - [Structure of JWT](#structure-of-jwt)
    - [Header](#header)
    - [Payload](#payload)
    - [Signature](#signature)
  - [Is JWT Encrypted?](#is-jwt-encrypted)
  - [Digital Signature in JWT](#digital-signature-in-jwt)
  - [Decoding JWTs on jwt.io](#decoding-jwts-on-jwtio)
- [Docker File, Github Actions and Build information](#docker-file-github-actions-and-build-information)

## Structure of JWT

A JWT consists of three parts, each separated by a dot (`.`):

### Header

This part typically contains two fields:
- `alg`: The signing algorithm being used (e.g., HMAC SHA256, RSA)
- `typ`: The type of token, which is usually "JWT"

Example of a header in JSON format:

```json
{
  "alg": "HS256",
  "typ": "JWT"
}
```

### Payload

This section contains the claims, which are statements about an entity (typically, the user) and additional data. Claims can include information such as user ID, expiration time, and issued time.

Example of a payload:

```json
{
  "sub": "user_id123",
  "iat": 1644744000,
  "exp": 1644768000
}
```

### Signature

To create the signature part, you take the encoded header and payload, concatenate them with a period (`.`), and then sign this string using the specified algorithm and a secret key.

Example of generating a signature:

```python
import hmac
import hashlib

secret = 'your-256-bit-secret'
data = base64url_encode(header) + "." + base64url_encode(payload)
signature = hmac.new(secret.encode(), data.encode(), hashlib.sha256).digest()
```

The final JWT looks like this: `xxxxx.yyyyy.zzzzz`, where:
- `xxxxx` is the Base64Url-encoded header
- `yyyyy` is the Base64Url-encoded payload
- `zzzzz` is the signature

## Is JWT Encrypted?

JWTs can be signed or encrypted. Signing ensures that the sender of the JWT is who they claim to be and that the message wasn't changed along the way. This is achieved through cryptographic algorithms like HMAC or RSA. However, signing does not encrypt the payload; thus, anyone who has access to the token can read its contents.

Encryption can be applied to ensure that only authorized parties can read the payload. In this case, a JWE (JSON Web Encryption) structure would be used.

## Digital Signature in JWT

JWTs are typically signed using either a secret key with the HMAC algorithm or a public/private key pair using algorithms like RSA or ECDSA. This signing process creates a digital signature that verifies two main aspects:

- **Integrity**: The signature ensures that the token has not been altered after it was issued. If any part of the JWT (header or payload) is modified, the signature will not match when verified, indicating tampering.
- **Authenticity**: The signature confirms that the token was created by a legitimate source. For instance, if a JWT is signed with a private key, only the holder of the corresponding public key can verify its authenticity, ensuring that it comes from a trusted issuer.

## Decoding JWTs on jwt.io

The website jwt.io provides a convenient tool for decoding JWTs. When you paste a JWT into jwt.io:

- The tool automatically splits it into its three components: header, payload, and signature
- It decodes each part from Base64Url encoding back into JSON format
- You can view both the raw token and its decoded contents

This allows developers to inspect the claims contained in the token easily and verify its integrity by checking if the signature matches.

---

# Docker File, Github Actions and Build information

- This project has a automated CICD pipeline that automatically deploy the application into AWS ec2 instance. 
- AWS instance has a github runner installed. 
- Self-hosted runners allow you to customize the hardware and software environment according to your specific needs. This could include using specialized hardware or software not available on GitHub-hosted runners.
You can deploy self-hosted runners on physical machines, virtual machines, or even in cloud environments. They require you to manage updates and maintenance.These runners can be configured at different organizational levels—repository-level, organization-level, or enterprise-level—allowing flexibility in how they are utilized across projects
  - [Docs - Github self hosted runners ](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners/about-self-hosted-runners)
- The runner is run on the EC2 instance as a service (using systemctl)
  - [Docs - How to runner app as a service](https://docs.github.com/en/actions/hosting-your-own-runners/managing-self-hosted-runners/configuring-the-self-hosted-runner-application-as-a-service)
  
  