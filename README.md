# JWT-token-based-authentication-in-Golang
JWT token based authorization in Golang  
JWT stands for JSON Web Token  
It is used to do authentication and authorization based on token  
- I will do three things in this project
  - Login API: One is to create a login API so user can login. Once the user is logged in an authentication JWT token is created for it. For every request JWT token is passed from the client so the server can authenticate and authorize the request itself.
  - Home API: I will be creating a homepage API that will authenticate the token, authorize the token and then if the token is valid then the user is allowed to visit the homepage or else it will throw the unauthorized authorization exception.
  -  Refresh API: I will aslo create a API for refresh token, so suppose if i have created the token  that will be availabe for 5 minutes suppose, so after 5 minutes the token has to be refreshed and reused again.  
So I will create an API to refresh the token again and token can be used in the API's.
